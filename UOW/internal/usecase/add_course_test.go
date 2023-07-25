package usecase

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddCourse(t *testing.T) {
	dbt, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	assert.NoError(t, err)

	dbt.Exec("DROP TABLE IF EXISTS courses")
	dbt.Exec("DROP TABLE IF EXISTS categories")

	dbt.Exec("CREATE TABLE `categories` (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255) NOT NULL, PRIMARY KEY (id))")
	dbt.Exec("CREATE TABLE `courses` (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255) NOT NULL, category_id INT NOT NULL, PRIMARY KEY (id), FOREIGN KEY (category_id) REFERENCES categories(id))")

	input := InputUseCase{
		CategoryName:     "Category 1",
		CourseName:       "Course 1",
		CourseCategoryID: 1,
	}

	ctx := context.Background()
	useCase := NewAddCourseCategory(NewCourseRepository(dbt), NewCategoryRepository(dbt))
	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}
