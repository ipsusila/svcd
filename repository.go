package svcd

import (
	"github.com/asdine/storm/v3"
)

// Repository for storing node info to BoltDB
type Repository struct {
	db *storm.DB
}

// NewRepository with given database name
func NewRepository(name string) (*Repository, error) {
	db, err := storm.Open(name)
	if err != nil {
		return nil, err
	}

	return &Repository{db: db}, nil
}

// Close repository (not thread safe)
func (r *Repository) Close() error {
	var err error
	if r.db != nil {
		err = r.db.Close()
		r.db = nil
	}
	return err
}

// Save node to repository
func (r *Repository) Save(n *Node) error {
	if err := r.db.Save(n); err != nil {
		return err
	}
	return nil
}
