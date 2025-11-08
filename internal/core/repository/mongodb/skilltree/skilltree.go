package skilltree

import (
	"context"
	"fmt"
	"log"
	"vacanciesParser/internal/core/logic/skilltree"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func getSkillTreeCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("skill_importance").Collection("skill_tree")
}

func NewSkillTreeRepository(client *mongo.Client) *Repository {
	return &Repository{
		collection: getSkillTreeCollection(client),
	}
}

// TODO: Fix pipiline: all elements are just descendants of root
func (r *Repository) GetSkillTree(ctx context.Context) (*skilltree.Node, error) {
	findRoot := bson.D{
		{Key: "$match", Value: bson.D{
			{Key: "parent_id", Value: bson.D{
				{Key: "$exists", Value: false},
			}},
		}},
	}

	findDescendants := bson.D{
		{Key: "$graphLookup", Value: bson.D{
			{Key: "from", Value: "skill_tree"},
			{Key: "startWith", Value: "$_id"},
			{Key: "connectFromField", Value: "_id"},
			{Key: "connectToField", Value: "parent_id"},
			{Key: "as", Value: "descendants"},
		}},
	}

	pipeline := mongo.Pipeline{findRoot, findDescendants}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("ошибка агрегации: %w", err)
	}
	defer cursor.Close(ctx)

	var resultDoc *SkillNode
	resultsNum := 0
	for cursor.Next(ctx) {
		resultsNum++
		if resultsNum > 1 {
			return nil, fmt.Errorf("найдено более одного корневого элемента")
		}

		if err := cursor.Decode(&resultDoc); err != nil {
			return nil, fmt.Errorf("ошибка декодирования документа: %w", err)
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("ошибка курсора после итерации: %w", err)
	}

	return resultDoc.ToLogic(), nil
}

func isRoot(node *skilltree.NodePath) bool {
	return len(node.Path) == 0
}

func isSubroot(node *skilltree.NodePath) bool {
	return len(node.Path) == 1
}

func (r *Repository) findPotentialParents(ctx context.Context, node *skilltree.NodePath) (map[*SkillNode]struct{}, error) {
	parentIdx := len(node.Path) - 1

	cursor, err := r.collection.Find(ctx, bson.M{
		"name": node.Path[parentIdx],
	})
	if err != nil {
		return nil, fmt.Errorf("не удалось найти навык под названием %s: %v", node.Path[parentIdx], err)
	}
	defer func() {
		if err := cursor.Close(ctx); err != nil {
			log.Printf("ошибка при закрытии курсора: %v", err)
		}
	}()

	potentialParents := make(map[*SkillNode]struct{})
	for cursor.Next(ctx) {
		var doc *SkillNode
		if err := cursor.Decode(&doc); err != nil {
			log.Printf("ошибка при декодировании документа: %v", err)
			continue
		}

		potentialParents[doc] = struct{}{}
	}

	return potentialParents, nil
}

func (r *Repository) getByID(ctx context.Context, ID primitive.ObjectID) (*SkillPath, error) {
	var result *SkillPath
	err := r.collection.FindOne(ctx, bson.M{"_id": ID}).Decode(&result)

	return result, err
}

func (r *Repository) isParent(ctx context.Context, potentialParent *SkillNode, node *skilltree.NodePath) (bool, error) {
	grandparentIdx := len(node.Path) - 2
	potentialAncestor := potentialParent

	for checkIdx := grandparentIdx; checkIdx >= 0; checkIdx-- {
		potentialAncestor, err := r.getByID(ctx, *potentialAncestor.ParentID)
		if err == mongo.ErrNoDocuments {
			log.Printf("объект с этим ID не найден: %v", potentialAncestor.ParentID)
			return false, nil
		} else if err != nil {
			log.Printf("ошибка при поиске предка %s: %v\n", potentialAncestor.ParentID.Hex(), err)
			return false, err
		}

		if potentialAncestor.Name != node.Path[checkIdx] {
			return false, nil
		}
	}

	return true, nil
}

func (r *Repository) insertChildren(ctx context.Context, parentID primitive.ObjectID, childrenName string) error {
	if _, err := r.collection.InsertOne(ctx, SkillPath{
		Name:     childrenName,
		ParentID: &parentID,
	}); err != nil {
		return fmt.Errorf("не удалось выполнить вставку в MongoDB (узел: %s, родитель: %s): %v", childrenName, parentID.Hex(), err)
	}

	return nil
}

// TODO: Add name duplication check for the same parent
func (r *Repository) CreateNode(ctx context.Context, node *skilltree.NodePath) error {
	if isRoot(node) {
		return fmt.Errorf("нельзя создавать новые корневые узлы")
	}

	if isSubroot(node) {
		root := &SkillPath{}
		err := r.collection.FindOne(ctx, bson.M{"parent_id": nil}).Decode(&root)
		if err != nil {
			return fmt.Errorf("не удалось выполнить поиск корневого узла в MongoDB: %v", err)
		}

		return r.insertChildren(ctx, root.ID, node.Name)
	}

	potentialParents, err := r.findPotentialParents(ctx, node)
	if err != nil {
		return fmt.Errorf("не удалось выполнить поиск потенциальных родителей в MongoDB: %v", err)
	}

	for potentialParent := range potentialParents {
		isParent, err := r.isParent(ctx, potentialParent, node)
		if err != nil {
			return fmt.Errorf("не удалось выполнить проверку потенциального родителя в MongoDB: %v", err)
		}

		if isParent {
			return r.insertChildren(ctx, potentialParent.ID, node.Name)
		}
	}

	return fmt.Errorf("некорректный путь к создаваемому элементу дерева навыков")
}
