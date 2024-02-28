package teachable

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	neturl "net/url"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

func (s *Service) GetPublishedCourses(ctx context.Context) (response CoursesResponse, err error) {
	retryClient := retryablehttp.NewClient()
	retryClient.Logger = nil
	retryClient.RetryMax = 2
	client := retryClient.StandardClient()
	client.Timeout = time.Duration(15) * time.Second

	url := fmt.Sprintf("%s/v1/courses", s.teachableURL)

	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apiKey", s.teachableAPIKey)

	req.URL.RawQuery = neturl.Values{"is_published": {"true"}}.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("External Teachable request - Could not make get courses names request", err)
		return response, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("External Teachable request - Get courses names request error", resp.StatusCode)
		return response, err
	}

	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		fmt.Println("External Teachable request - Could not read response body from courses names request", err)
		return response, err
	}

	err = json.Unmarshal(body.Bytes(), &response)
	if err != nil {
		fmt.Println(ctx, "External - Could not unmarshal body from courses names request", err)
	}

	return response, err
}

func (s *Service) GetEnrolledStudents(ctx context.Context, courseID int) (response EnrollmentsResponse, err error) {
	retryClient := retryablehttp.NewClient()
	retryClient.Logger = nil
	retryClient.RetryMax = 2
	client := retryClient.StandardClient()
	client.Timeout = time.Duration(15) * time.Second

	url := fmt.Sprintf("%s/v1/courses/%d/enrollments", s.teachableURL, courseID)

	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apiKey", s.teachableAPIKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("External Teachable request - Could not make get enrolled students request", err)
		return response, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("External Teachable request - Get enrolled students request error", resp.StatusCode)
		return response, err
	}

	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		fmt.Println("External Teachable request - Could not read response body from enrolled students request", err)
		return response, err
	}

	err = json.Unmarshal(body.Bytes(), &response)
	if err != nil {
		fmt.Println(ctx, "External - Could not unmarshal body from enrolled students request", err)
	}

	return response, err
}

func (s *Service) GetStudent(ctx context.Context, userID int) (response User, err error) {
	retryClient := retryablehttp.NewClient()
	retryClient.Logger = nil
	retryClient.RetryMax = 2
	client := retryClient.StandardClient()
	client.Timeout = time.Duration(15) * time.Second

	url := fmt.Sprintf("%s/v1/users/%d", s.teachableURL, userID)

	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apiKey", s.teachableAPIKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("External Teachable request - Could not make get student request", err)
		return response, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("External Teachable request - Get student request error", resp.StatusCode)
		return response, err
	}

	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		fmt.Println("External Teachable request - Could not read response body from student request", err)
		return response, err
	}

	err = json.Unmarshal(body.Bytes(), &response)
	if err != nil {
		fmt.Println(ctx, "External - Could not unmarshal body from student request", err)
	}

	return response, err
}
