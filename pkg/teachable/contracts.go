package teachable

type Service struct {
	teachableURL    string
	teachableAPIKey string
}

func NewService(teachableURL string, teachableAPIKey string) *Service {
	return &Service{
		teachableURL:    teachableURL,
		teachableAPIKey: teachableAPIKey,
	}
}

type CoursesResponse struct {
	Courses []Course `json:"courses"`
}

type Course struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Heading     string `json:"heading"`
	IsPublished bool   `json:"is_published"`
}

type EnrollmentsResponse struct {
	Enrollments []Enrollment `json:"enrollments"`
}

type Enrollment struct {
	UserId int `json:"user_id"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
