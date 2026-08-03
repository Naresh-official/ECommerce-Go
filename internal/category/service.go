package category

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/naresh-official/ecommerce_go/internal/database/sqlc"
)

var ErrCategoryNotFound = errors.New("Category not found")

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllCategories(ctx context.Context) ([]CategoryResponse, error) {
	categories, err := s.repo.GetAllCategories(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]CategoryResponse, 0, len(categories))
	for _, category := range categories {
		responses = append(responses, mapCategory(category))
	}

	return responses, nil
}

func (s *Service) GetCategoryByID(ctx context.Context, categoryID string) (*CategoryResponse, error) {
	category, err := s.repo.GetCategoryByID(ctx, uuid.MustParse(categoryID))
	if err != nil {
		if errors.Is(err, NoRowsError) {
			return nil, ErrCategoryNotFound
		}
		return nil, err
	}

	response := mapCategory(*category)
	return &response, nil
}

func (s *Service) CreateCategory(ctx context.Context, input CreateCategoryRequest) (*CategoryResponse, error) {
	category, err := s.repo.CreateCategory(ctx, input)
	if err != nil {
		return nil, err
	}

	response := mapCategory(*category)
	return &response, nil
}

func (s *Service) UpdateCategory(ctx context.Context, categoryID string, input CreateCategoryRequest) (*CategoryResponse, error) {
	category, err := s.repo.UpdateCategory(ctx, uuid.MustParse(categoryID), input)
	if err != nil {
		if errors.Is(err, NoRowsError) {
			return nil, ErrCategoryNotFound
		}
		return nil, err
	}

	response := mapCategory(*category)
	return &response, nil
}

func (s *Service) DeleteCategory(ctx context.Context, categoryID string) error {
	if err := s.repo.DeleteCategory(ctx, uuid.MustParse(categoryID)); err != nil {
		if errors.Is(err, NoRowsError) {
			return ErrCategoryNotFound
		}
		return err
	}

	return nil
}

func mapCategory(category sqlc.Category) CategoryResponse {
	response := CategoryResponse{
		ID:   category.ID.String(),
		Name: category.Name,
	}

	if category.Description.Valid {
		response.Description = category.Description.String
	}

	return response
}
