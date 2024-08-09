build:
	go build cmd/main.go

run-dev: run-docker
	nodemon --watch "./**/*.go" --exec "go" run cmd/main.go

run:
	go run cmd/main.go
	
run-docker:
	docker compose -f docker-compose.yaml up -d


