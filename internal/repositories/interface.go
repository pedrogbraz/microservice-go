package repositories

import "github.com/pedrogbraz/microservice-go/internal/entities"

type ICategoryRepository interface {
	// Define os métodos que o repositório deve implementar
	Save(category *entities.Category) error
	List() ([]*entities.Category, error)
}
