package svcd

// NodeAddress information
type NodeAddress struct {
	Network string `json:"network"`
	Address string `json:"address"`
}

// NodeInterface description
type NodeInterface struct {
	Name      string         `json:"name"`
	HWAddr    string         `json:"hwAddr"`
	Addresses []*NodeAddress `json:"addresses"`
}

// Node of the service
type Node struct {
	ID          int              `json:"id" storm:"id,increment"`
	Name        string           `json:"name" storm:"unique"`
	Description string           `json:"description"`
	Interfaces  []*NodeInterface `json:"interfaces"`
}
