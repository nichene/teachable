//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -package=courses -source=contracts.go -destination=mock_contracts.go .
package courses

import (
	"context"
	"teachable/pkg/teachable"
)

type Service struct {
	teachableService TeachableService
}

func NewService(teachableService TeachableService) *Service {
	return &Service{
		teachableService: teachableService,
	}
}

type TeachableService interface {
	GetPublishedCourses(ctx context.Context) (response teachable.CoursesResponse, err error)
	GetEnrolledStudents(ctx context.Context, courseID int) (response teachable.EnrollmentsResponse, err error)
	GetStudent(ctx context.Context, userID int) (response teachable.User, err error)
}
