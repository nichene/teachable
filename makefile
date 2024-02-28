get-published-courses-names:
	go run ./cmd/*.go get-published-courses-info names

get-published-courses-headings:
	go run ./cmd/*.go get-published-courses-info headings

get-students:
	go run ./cmd/*.go get-students 2002430

mocks:
	go generate ./...

tests:
	go test ./...