package repository

import (
	"context"
	"database/sql"
	"teste/uow/internal/db"
	"teste/uow/internal/entity"
)

type CourseRepositoryInterface interface {
	Insert(ctx context.Context, course entity.Course) error
}

type CourseRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewCourseRepository(dtb *sql.DB) *CourseRepository {
	return &CourseRepository{
		DB:      dtb,
		Queries: db.New(dtb),
	}
}

func (r *CourseRepository) Insert(ctx context.Context, course entity.Course) error {
	return r.Queries.CreateCategory(ctx, db.CreateCategoryParams{
		Name: course.Name,
		ID:   int32(course.CategoryID),
	})
}
