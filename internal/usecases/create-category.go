package usecases

import (
	"github.com/pedrogbraz/microservice-go/internal/entities"
	"github.com/pedrogbraz/microservice-go/internal/repositories"
)

type createCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewCreateCategoryUseCase(repository repositories.ICategoryRepository) *createCategoryUseCase {
	return &createCategoryUseCase{repository}
}

func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	if err := u.repository.Save(category); err != nil {
		return err
	}

	return nil
}
