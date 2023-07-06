package services

import (
	"context"
	"teste/GRPC/internal/database"
	"teste/GRPC/internal/pb"
)

type CategoryServices struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryServices(categoryDB database.Category) *CategoryServices {
	return &CategoryServices{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryServices) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
	return &pb.CategoryResponse{Category: categoryResponse}, nil
}
