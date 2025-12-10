package repositories

import (
	"fmt"

	"github.com/pedrogbraz/microservice-go/internal/entities"
)

type inMemoryCategoryRepository struct {
	db     []*entities.Category
	lastID int
}

func NewInMemoryCategoryRepository() *inMemoryCategoryRepository {
	return &inMemoryCategoryRepository{
		db: make([]*entities.Category, 0),
	}
}

func (r *inMemoryCategoryRepository) Save(category *entities.Category) error {
	if category.ID == 0 {
		r.lastID++
		category.ID = r.lastID
	}

	r.db = append(r.db, category)
	return nil
}

func (r *inMemoryCategoryRepository) List() ([]*entities.Category, error) {
	return r.db, nil
}

func (r *inMemoryCategoryRepository) Delete(id int) error {
	for index, category := range r.db {
		if category.ID == id {
			r.db = append(r.db[:index], r.db[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("%w: %d", ErrCategoryNotFound, id)
}
