package courses

import (
	"context"
	"fmt"
)

func (s *Service) GetCoursesInfo(ctx context.Context, param string) (response []string, err error) {
	courses, err := s.teachableService.GetPublishedCourses(ctx)
	if err != nil {
		fmt.Println("Error fetching your courses names :(")
		return response, err
	}

	if param == "names" {
		for _, course := range courses.Courses {
			response = append(response, course.Name)
		}
	}
	if param == "headings" {
		for _, course := range courses.Courses {
			response = append(response, course.Heading)
		}
	}

	return response, err
}
