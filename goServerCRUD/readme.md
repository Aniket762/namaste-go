### Pre-requisites

```
go get "github.com/gorilla/mux"
```

#### Flow of the server

| Routes    | Functions   | Endpoints  | Methods |
| --------- | ----------- | ---------- | ------- |
| GET ALL   | getMovies   | /movies    | GET     |
| GET BY ID | getMovie    | /movies/id | GET     |
| CREATE    | createMovie | /movies    | POST    |
| UPDATE    | updateMovie | /movies/id | PUT     |
| DELETE    | deleteMovie | /movies/id | DELETE  |


#### Run the server

go run main.go

#### Test the server

go test
