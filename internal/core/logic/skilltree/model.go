package skilltree

type NodePath struct {
	RootName string      `json:"root_name"`
	Path     []*NodePath `json:"path"`
}

func (st *NodePath) GetNode() *NodePath {
	if len(st.Path) == 0 {
		return st
	}

	return st.Path[len(st.Path)-1]
}

type Node struct {
	Name        string `json:"name"`
	Descendants []Node `json:"descendants"`
}
