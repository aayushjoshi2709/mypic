APP_NAME=mypic

install:
	go mod tidy
	go mod download

run:
	go run main.go

dev:
	chmod +x run.shdocker logs ede4a605e3c1lsof -i :6379
	CompileDaemon --exclude-dir=docs --command="./run.sh"

build:
	go build -o $(APP_NAME) main.go

swag:
	swag init

clean:
	rm -f $(APP_NAME)