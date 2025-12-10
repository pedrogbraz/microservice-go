package usecases

import (
	"fmt"

	"github.com/pedrogbraz/microservice-go/internal/repositories"
)

type deleteCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewDeleteCategoryUseCase(repository repositories.ICategoryRepository) *deleteCategoryUseCase {
	return &deleteCategoryUseCase{repository}
}

func (u *deleteCategoryUseCase) Execute(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid category id: %d", id)
	}

	if err := u.repository.Delete(id); err != nil {
		return err
	}

	return nil
}
