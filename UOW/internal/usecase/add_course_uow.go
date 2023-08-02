package usecase

import (
	"teste/uow/internal/entity"
	"teste/uow/internal/repository"
	"teste/uow/pkg/uow"

	"golang.org/x/net/context"
)

type InputUseCaseUow struct {
	CategoryName     string
	CourseName       string
	CourseCategoryID int
}

type AddCourseUseCaseUow struct {
	Uow uow.UowInterface
}

func NewAddCourseCategoryUow(uow uow.UowInterface) *AddCourseUseCaseUow {
	return &AddCourseUseCaseUow{
		Uow: uow,
	}
}

func (a *AddCourseUseCaseUow) ExecuteUow(ctx context.Context, input InputUseCase) error {
	return a.Uow.Do(ctx, (func(uow *uow.Uow) error {
		category := entity.Category{
			Name: input.CategoryName,
		}

		repoCategory := a.getCategoryRepository(ctx)
		err := repoCategory.Insert(ctx, category)
		if err != nil {
			return err
		}

		course := entity.Course{
			Name:       input.CourseName,
			CategoryID: input.CourseCategoryID,
		}

		repositoryCourse := a.getCourseRepository(ctx)
		err = repositoryCourse.Insert(ctx, course)
		if err != nil {
			return err
		}
		return nil
	}))

}

func (a *AddCourseUseCaseUow) getCategoryRepository(ctx context.Context) repository.CategoryRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "categoryRepository")
	if err != nil {
		panic(err)
	}
	return repo.(repository.CategoryRepositoryInterface)
}

func (a *AddCourseUseCaseUow) getCourseRepository(ctx context.Context) repository.CourseRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "courseRepository")
	if err != nil {
		panic(err)
	}
	return repo.(repository.CourseRepositoryInterface)
}
