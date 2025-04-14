## Geofence

A simple CRUD Geofence micro service, written in Go using the Gin web framework.

#### Goals
* Learn the basics of Go
* Build a robust and generic CRUD service for Geofences

### Database

The project uses MongoDB and the [Go MongoDB Driver](https://github.com/mongodb/mongo-go-driver).

### Configuration

The project uses [godotenv](https://www.github.com/joho/godotenv) to work with environment variables.

Any missing required env vars will cause an error and the program to exit.


| Variable | Default | Description                                |
| -------- | ------- | ------------------------------------------ |
| PORT     | 8080    | The port the server should listen on.      |
| DB_HOST  | N/A     | The host address of the database.          |
| DB_NAME  | N/A     | The name of the database.                  |
| DB_USER  | N/A     | The username to use during authentication. |
| DB_PASS  | N/A     | The password to use during authentication. |


### Docker

Docker is used to lift and manage a Mongo DB for development.

To start the Docker service.
```sh
docker compose up
```

### Compile
Compile the files.

```sh
go build .
```

### Run
Run the build.

```sh
go run main.go
```