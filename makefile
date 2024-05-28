run-dev:
	@GO_ENV="development" go run ./cmd/main.go

run-prod:
	@GO_ENV="production" go run ./cmd/main.go
