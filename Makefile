vendor:
	go mod tidy
	go mod vendor

tests:
	go test ./...

mocks:
	mockery --all --exclude vendor --keeptree