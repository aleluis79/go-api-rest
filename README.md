# Example of rest api in golang with swagger

## Build project

### Downloads dependencies

go tidy

### Generate swagger documentation:

swag init

#http://localhost:8080/swagger/index.html# Para arrancar la aplicación correr el siguiente commando:

go run .

### Try in the browser:

http://localhost:8080/swagger/index.html

## Docker

### Create image

docker build -t apigo .

### Run image

docker run --rm -d --name apigo -p 8080:8080 apigo
