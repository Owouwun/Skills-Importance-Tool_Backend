package skilltree

type NodePath struct {
	Name string   `json:"name"`
	Path []string `json:"path"`
}

type Node struct {
	Name     string `json:"name"`
	Children []Node `json:"children"`
}
