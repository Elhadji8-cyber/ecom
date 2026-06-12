package service

import (
	"server/internal/product/dto"
	"server/internal/product/models"
	"server/internal/product/repository"

	"github.com/google/uuid"
)

type ProductService interface {
	CreateProduct(req dto.ProductRequest) (*dto.ProductResponse, error)
	GetAllProducts() ([]dto.ProductResponse, error)
	GetProductByID(id uuid.UUID) (*dto.ProductResponse, error)
	UpdateProduct(id uuid.UUID, req dto.ProductRequest) (*dto.ProductResponse, error)
	DeleteProduct(id uuid.UUID) error
	CreateCategory(name, description string) (*models.Category, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) CreateProduct(req dto.ProductRequest) (*dto.ProductResponse, error) {
	product := &models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CategoryID:  req.CategoryID,
		ImageURL:    req.ImageURL,
	}

	if err := s.repo.Create(product); err != nil {
		return nil, err
	}

	return s.GetProductByID(product.ID)
}

func (s *productService) GetAllProducts() ([]dto.ProductResponse, error) {
	products, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var res []dto.ProductResponse
	for _, p := range products {
		res = append(res, dto.ProductResponse{
			ID:           p.ID,
			Name:         p.Name,
			Description:  p.Description,
			Price:        p.Price,
			Stock:        p.Stock,
			CategoryID:   p.CategoryID,
			CategoryName: p.Category.Name,
			ImageURL:     p.ImageURL,
		})
	}
	return res, nil
}

func (s *productService) GetProductByID(id uuid.UUID) (*dto.ProductResponse, error) {
	p, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		ID:           p.ID,
		Name:         p.Name,
		Description:  p.Description,
		Price:        p.Price,
		Stock:        p.Stock,
		CategoryID:   p.CategoryID,
		CategoryName: p.Category.Name,
		ImageURL:     p.ImageURL,
	}, nil
}

func (s *productService) UpdateProduct(id uuid.UUID, req dto.ProductRequest) (*dto.ProductResponse, error) {
	p, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	p.Name = req.Name
	p.Description = req.Description
	p.Price = req.Price
	p.Stock = req.Stock
	p.CategoryID = req.CategoryID
	p.ImageURL = req.ImageURL

	if err := s.repo.Update(p); err != nil {
		return nil, err
	}

	return s.GetProductByID(id)
}

func (s *productService) DeleteProduct(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *productService) CreateCategory(name, description string) (*models.Category, error) {
	category := &models.Category{
		Name:        name,
		Description: description,
	}
	err := s.repo.CreateCategory(category)
	return category, err
}
