package skilltree

import (
	"context"
	"fmt"
	"log"
	"vacanciesParser/internal/core/logic/skilltree"
	"vacanciesParser/internal/core/repository/mongodb"

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

func (r *Repository) findPotentialParents(ctx context.Context, node *skilltree.NodePath) (*mongo.Cursor, error) {
	parentIdx := len(node.Path) - 2

	return r.collection.Find(ctx, SkillPath{
		RootName: node.Path[parentIdx].RootName,
	})
}

func (r *Repository) getByID(ctx context.Context, ID primitive.ObjectID) (*SkillPath, error) {
	var result *SkillPath
	err := r.collection.FindOne(ctx, SkillPath{ID: ID}).Decode(&result)

	return result, err
}

func (r *Repository) isParentID(ctx context.Context, potentialParentID primitive.ObjectID, node *skilltree.NodePath) (bool, error) {
	grandparentIdx := len(node.Path) - 3

	for checkIdx := grandparentIdx; checkIdx >= 0; checkIdx-- {
		if potentialParentID.IsZero() {
			return false, nil
		}

		skill, err := r.getByID(ctx, potentialParentID)
		if err == mongo.ErrNoDocuments {
			return false, nil
		} else if err != nil {
			log.Printf("ошибка при поиске предка %s: %v\n", potentialParentID.Hex(), err)
			return false, err
		}

		if skill.RootName != node.Path[checkIdx].RootName {
			return false, nil
		}

		potentialParentID = *skill.ParentID
	}

	return true, nil
}

func (r *Repository) insertChildren(ctx context.Context, parentID primitive.ObjectID, childrenName string) error {
	if _, err := r.collection.InsertOne(ctx, SkillPath{
		RootName: childrenName,
		ParentID: &parentID,
	}); err != nil {
		return fmt.Errorf("не удалось выполнить вставку в MongoDB (узел: %s, родитель: %s): %v", childrenName, parentID.Hex(), err)
	}

	return nil
}

func (r *Repository) CreateNode(ctx context.Context, node *skilltree.NodePath) error {
	if isRoot(node) {
		return fmt.Errorf("нельзя создавать новые корневые узлы")
	}

	if isSubroot(node) {
		root := &SkillPath{}
		err := r.collection.FindOne(ctx, SkillPath{RootName: "root"}).Decode(&root)
		if err != nil {
			return fmt.Errorf("не удалось выполнить поиск корневого узла в MongoDB: %v", err)
		}

		return r.insertChildren(ctx, root.ID, node.GetNode().RootName)
	}

	cursor, err := r.findPotentialParents(ctx, node)
	if err != nil {
		return fmt.Errorf("не удалось выполнить поиск потенциальных родителей в MongoDB: %v", err)
	}
	defer func() {
		if err := cursor.Close(ctx); err != nil {
			log.Printf("ошибка при закрытии курсора: %v", err)
		}
	}()

	potentialParentIDs := mongodb.GetIDs(ctx, cursor)

	for parentID := range potentialParentIDs {
		isParentID, err := r.isParentID(ctx, parentID, node)
		if err != nil {
			return fmt.Errorf("не удалось выполнить проверку потенциального родителя в MongoDB: %v", err)
		}

		if isParentID {
			return r.insertChildren(ctx, parentID, node.GetNode().RootName)
		}
	}

	return fmt.Errorf("некорректный путь к создаваемому элементу дерева навыков")
}
