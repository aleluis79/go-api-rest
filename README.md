# Example of rest api in golang with swagger

## Build project

### Downloads dependencies

go mod tidy

### Generate swagger documentation:

* Install swag

  go install github.com/swaggo/swag/cmd/swag@latest

* Add to path:

  export PATH=$PATH:/home/$USER/go/bin

* Create docs

  swag init

### Para arrancar la aplicación correr el siguiente commando:

go run .

### Try in the browser:

http://localhost:8080/swagger/index.html

## Docker

### Create image

docker build -t apigo .

### Run image

docker run --rm -d --name apigo -p 8080:8080 apigo
