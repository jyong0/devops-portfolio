run:
	go run ./app/cmd/server/main.go

docker-build:
	docker build -t devops-app:latest .

docker-run:
	docker run -p 8080:8080 --env-file .env devops-app:latest
