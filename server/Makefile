APP_NAME=mypic

run:
	go run main.go

dev:
	swag init
	go run main.go

build:
	go build -o $(APP_NAME) main.go

swag:
	swag init

clean:
	rm -f $(APP_NAME)