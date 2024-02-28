package courses

import (
	"context"
	"fmt"
	"teachable/pkg/teachable"
)

func (s *Service) GetStudentsInACourse(ctx context.Context, courseID int) (response []teachable.User, err error) {

	students, err := s.teachableService.GetEnrolledStudents(ctx, courseID)
	if err != nil {
		fmt.Println("Error fetching students", err)
		return response, err
	}

	for _, student := range students.Enrollments {
		studentInfo, err := s.teachableService.GetStudent(ctx, student.UserId)
		if err != nil {
			fmt.Println("Error fetching student info", err)
			continue
		}
		response = append(response, studentInfo)
	}

	return response, err
}
