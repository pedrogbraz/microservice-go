package entities

import (
	"fmt"
	"time"
)

type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreateAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCategory(name string) (*Category, error) {
	category := &Category{
		Name:      name,
		CreateAt:  time.Now(),
		UpdatedAt: time.Now(),
	}
	// Business rules
	if err := category.isValid(); err != nil {
		return nil, err
	}

	return category, nil
}

func (c *Category) isValid() error {
	if len(c.Name) < 5 {
		return fmt.Errorf("name must be greater than 5. got %d", len(c.Name))
	}
	return nil
}
