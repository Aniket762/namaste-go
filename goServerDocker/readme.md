### Pre-requisites

```
go get "github.com/gin-gonic/gin"
```

### Flow of the Server 💪

Aim was to build a basic server using golang and gin framework running inside Docker environment.

server -> /hello ->HelloWorldHandler func

### To Run the server Locally

go run main.go

### To Run the server inside Docker

docker compose up  

### Stop the server and remove containers inside Docker

docker compose down

###  Choose Dockerfile

You can choose which docker file need to be used for creating the container, choose multistage dockerfile for low size.
specify it in bellow part of docker-compose.yaml file  

      dockerfile: Dockerfile.multistage # Specifies the Dockerfile to use (multi-stage / normal Dockerfile).



