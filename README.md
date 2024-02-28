## teachable

command line service developed for assessment.

## Pre configurations

- To load env vars copy the `.env.example` to a `.env` file and modify accordingly.

## To run commands

- use `make get-published-courses-names` to retrieve public courses names

- use `make get-published-courses-headings` to retrieve public courses headings

- use `go run ./cmd/*.go get-students {course_id} ` to retrive a list of the names and emails of students actively enrolled in the course. Alternatively use `make get-students` for an example.

## Testing

Run the project tests by running `make tests`.

## Generate mocks

- To generate test mocks install mockgen `go install github.com/golang/mock/mockgen@v1.6.0`. Docs at: [mockgen](https://github.com/golang/mock)

- run `make mocks`
