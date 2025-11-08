package skilltree

import (
	"vacanciesParser/internal/core/logic/skilltree"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SkillPath struct {
	ID       primitive.ObjectID  `bson:"_id,omitempty"`
	Name     string              `bson:"name"`
	ParentID *primitive.ObjectID `bson:"parent_id,omitempty"`
	NodePath []string            `bson:"parents,omitempty"`
}

type SkillNode struct {
	ID       primitive.ObjectID  `bson:"_id"`
	Name     string              `bson:"name"`
	ParentID *primitive.ObjectID `bson:"parent_id,omitempty"`

	Descendants []SkillNode `bson:"descendants"`
}

func (n *SkillNode) ToLogic() *skilltree.Node {
	descendants := make([]skilltree.Node, len(n.Descendants))
	for i, d := range n.Descendants {
		descendants[i] = *d.ToLogic()
	}

	return &skilltree.Node{
		Name:        n.Name,
		Descendants: descendants,
	}
}
