package courses

import (
	"context"
	"errors"
	"teachable/pkg/teachable"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestCreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	teachableServiceMock := NewMockTeachableService(ctrl)

	service := NewService(teachableServiceMock)

	t.Run("Should get courses info: names", func(t *testing.T) {
		teachableServiceMock.EXPECT().GetPublishedCourses(ctx).
			Return(teachable.CoursesResponse{Courses: []teachable.Course{{
				Name:    "mock Name",
				Heading: "mock Heading",
			}}}, nil)

		info, err := service.GetCoursesInfo(ctx, "names")

		expectedInfo := []string{"mock Name"}

		assert.ErrorIs(t, err, nil)
		assert.Equal(t, expectedInfo[0], info[0])
	})

	t.Run("Should get courses info: headings", func(t *testing.T) {
		teachableServiceMock.EXPECT().GetPublishedCourses(ctx).
			Return(teachable.CoursesResponse{Courses: []teachable.Course{{
				Name:    "mock Name",
				Heading: "mock Heading",
			}}}, nil)

		info, err := service.GetCoursesInfo(ctx, "headings")

		expectedInfo := []string{"mock Heading"}

		assert.ErrorIs(t, err, nil)
		assert.Equal(t, expectedInfo[0], info[0])
	})

	t.Run("Should return error", func(t *testing.T) {
		expectedErr := errors.New("get error")

		teachableServiceMock.EXPECT().GetPublishedCourses(ctx).
			Return(teachable.CoursesResponse{}, expectedErr)

		_, err := service.GetCoursesInfo(ctx, "headings")

		assert.ErrorIs(t, err, expectedErr)
	})
}
