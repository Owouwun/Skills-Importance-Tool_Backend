package skilltree

import (
	"vacanciesParser/internal/core/logic/skilltree"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SkillNode struct {
	ID       primitive.ObjectID  `bson:"_id,omitempty"`
	Name     string              `bson:"name"`
	ParentID *primitive.ObjectID `bson:"parent_id,omitempty"`
}

type SkillPath struct {
	ID       primitive.ObjectID  `bson:"_id,omitempty"`
	Name     string              `bson:"name"`
	ParentID *primitive.ObjectID `bson:"parent_id,omitempty"`
	NodePath []string            `bson:"parents,omitempty"`
}

type SkillTree struct {
	RootID   primitive.ObjectID `bson:"_id,omitempty"`
	RootName string             `bson:"name"`
	Nodes    []SkillNode        `bson:"nodes"` // Except root
}

func (skillTree *SkillTree) ToLogic() *skilltree.Node {
	IDToNode := make(map[primitive.ObjectID]*SkillNode)
	children := make(map[primitive.ObjectID]map[primitive.ObjectID]struct{})

	children[skillTree.RootID] = make(map[primitive.ObjectID]struct{})
	for _, node := range skillTree.Nodes {
		IDToNode[node.ID] = &node
		children[node.ID] = make(map[primitive.ObjectID]struct{})
	}

	for _, node := range skillTree.Nodes {
		children[*node.ParentID][node.ID] = struct{}{}
	}

	var buildChildrenNodes func(id primitive.ObjectID) []skilltree.Node
	buildChildrenNodes = func(id primitive.ObjectID) []skilltree.Node {
		if len(children[id]) == 0 {
			return nil
		}

		curChildren := make([]skilltree.Node, 0, len(children[id]))

		for childID := range children[id] {
			curChildren = append(curChildren, skilltree.Node{
				Name:     IDToNode[childID].Name,
				Children: buildChildrenNodes(childID),
			})
		}

		return curChildren
	}

	result := &skilltree.Node{
		Name:     skillTree.RootName,
		Children: buildChildrenNodes(skillTree.RootID),
	}

	return result
}
