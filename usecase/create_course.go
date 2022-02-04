package usecase

import (
	"github.com/google/uuid"
	"github.com/robsantossilva/intensivao-golang-6/entity"
)

type CreateCourse struct {
	Repository entity.CourseRepository
}

func (c CreateCourse) Execute(input CreateCourseInputDto) (CreateCourseOutputDto, error) {

	course := entity.Course{}
	course.ID = uuid.New().String()
	course.Name = input.Name
	course.Description = input.Description
	course.Status = input.Status

	err := c.Repository.Insert(course)
	if err != nil {
		return CreateCourseOutputDto{}, err
	}

	output := CreateCourseOutputDto{
		ID:          course.ID,
		Name:        course.Name,
		Description: course.Description,
		Status:      course.Status,
	}

	return output, nil
}
