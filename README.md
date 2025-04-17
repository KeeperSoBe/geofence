<a name="readme-top"></a>

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/keepersobe/notes-client">
    <svg width="80" height="80" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 576 512">
        <path fill="#79d4fd" d="M408 120c0 54.6-73.1 151.9-105.2 192c-7.7 9.6-22 9.6-29.6 0C241.1 271.9 168 174.6 168 120C168 53.7 221.7 0 288 0s120 53.7 120 120zm8 80.4c3.5-6.9 6.7-13.8 9.6-20.6c.5-1.2 1-2.5 1.5-3.7l116-46.4C558.9 123.4 576 135 576 152l0 270.8c0 9.8-6 18.6-15.1 22.3L416 503l0-302.6zM137.6 138.3c2.4 14.1 7.2 28.3 12.8 41.5c2.9 6.8 6.1 13.7 9.6 20.6l0 251.4L32.9 502.7C17.1 509 0 497.4 0 480.4L0 209.6c0-9.8 6-18.6 15.1-22.3l122.6-49zM327.8 332c13.9-17.4 35.7-45.7 56.2-77l0 249.3L192 449.4 192 255c20.5 31.3 42.3 59.6 56.2 77c20.5 25.6 59.1 25.6 79.6 0zM288 152a40 40 0 1 0 0-80 40 40 0 1 0 0 80z"/>
    </svg>
  </a>

<h3 align="center">Geofence</h3>

  <p align="center">
    A simple CRUD Geofence microservice, written in Go using the Gin web framework.
    <br />
    <br />
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li>
      <a href="#usage">Usage</a>
      <ul>
        <li><a href="#development">Development</a></li>
      </ul>
    </li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

A learning exercise in the basics of Go lang and the Gin web framework.

#### Goals

- Learn the basics of Go
- Build a robust and generic CRUD service for Geofences

### Built With

- [![Go][Go]][Go-url]
- [![Gin][Gin]][Gin-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->

## Getting Started

This is an example of how you may give instructions on setting up your project locally.

To get a local copy up and running follow these simple example steps.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- PREREQUISITES -->

### Prerequisites

- Install [Go lang](https://go.dev/doc/install)

<!-- INSTALLATION -->

### Installation

- Clone the repository

```sh
$ git clone https://github.com/keepersobe/geofence.git
```

- Compile the files.

```sh
go build .
```

- Run the build.

```sh
go run .
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Development

### Database

The project uses MongoDB and the [Go MongoDB Driver](https://github.com/mongodb/mongo-go-driver).

### Docker

Docker is used to lift and manage a Mongo DB for development.

To start the Docker service.

```sh
docker compose up
```

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

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE -->

## Usage

### Entities

Database entities all share the same following properties.

| Property  | Type      | Description                                                                                  |
| --------- | --------- | -------------------------------------------------------------------------------------------- |
| id        | `string`  | The UUID of the entity, referenced by the program and unrelated to the underlying mongo \_id |
| type      | `"point"` | The type of the geofence.                                                                    |
| createdAt | `string`  | The datetime stamp the entity was created, ISO 8601 string format.                           |
| updatedAt | `string`  | The datetime stamp the entity was last updated, ISO 8601 string format.                      |

#### Point

The simplest geofence and the building block of other more complex models.

| Property    | type      | Description                  |
| ----------- | --------- | ---------------------------- |
| radius      | `uint`    | The size of the point        |
| coordinates | `[2]uint` | The coordinates of the point |

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Endpoints

### List

Gets all entities.

#### Example Request

GET: `http://localhost:8080`

#### Example Response

```json
// Status: 200
[
  {
    "id": "abcd-1234-efgh-5678",
    "type": "point",
    "radius": 1,
    "coordinates": [1234, 5678],
    "createdAt": "2025-02-27T03:36:32.616Z",
    "updatedAt": "2025-02-27T03:36:32.616Z"
  }
]
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Get

Gets an entity by its ID.

#### Example Request

GET: `http://localhost:8080/abcd-1234-efgh-5678`

#### Example Response

```json
// Status: 200
{
  "id": "abcd-1234-efgh-5678",
  "type": "point",
  "radius": 1,
  "coordinates": [1234, 5678],
  "createdAt": "2025-02-27T03:36:32.616Z",
  "updatedAt": "2025-02-27T03:36:32.616Z"
}
```

##### Error responses

- [BadRequest](#badrequest)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Create

Creates a new entity.

#### Example Request

POST: `http://localhost:8080`

```json
{
  "type": "point",
  "radius": 1,
  "coordinates": [1234, 5678]
}
```

#### Example Response

```json
// Status: 201
{
  "id": "abcd-1234-efgh-5678",
  "type": "point",
  "radius": 1,
  "coordinates": [1234, 5678],
  "createdAt": "2025-02-27T03:36:32.616Z",
  "updatedAt": "2025-02-27T03:36:32.616Z"
}
```

##### Error responses

- [BadRequest](#badrequest)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Update

Updates an existing entity.

Returns the updated properties, including updatedAt.

#### Example Request

PATCH: `http://localhost:8080/abcd-1234-efgh-5678`

```json
{
  "type": "point",
  "radius": 1,
  "coordinates": [1234, 5678],
  "updatedAt": "2025-02-27T03:36:32.616Z"
}
```

#### Example Responses

```json
// Status: 200
{
  "id": "abcd-1234-efgh-5678",
  "type": "point",
  "radius": 1,
  "coordinates": [1234, 5678],
  "createdAt": "2025-02-27T03:36:32.616Z",
  "updatedAt": "2025-02-27T03:36:32.616Z"
}
```

##### Error responses

- [NotFound](#notfound)
- [BadRequest](#badrequest)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Delete

Deletes an existing entity.

#### Example Request

DELETE: `http://localhost:8080/abcd-1234-efgh-5678`

#### Example Responses

```json
// Status: 200
{
  "deletedAt": "2025-02-27T03:36:32.616Z"
}
```

##### Error responses

- [NotFound](#notfound)
- [BadRequest](#badrequest)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Error Responses

The API will respond to all errors in a generic fashion.

All error responses will follow the same structure.

| Property   | Description                                                                 |
| ---------- | --------------------------------------------------------------------------- |
| statusCode | The status code of the error, this will be the same as the http status code |
| message    | The generic error message                                                   |

<p align="right">(<a href="#readme-top">back to top</a>)</p>

#### BadRequest

```json
// Status: 400
{
  "statusCode": 400,
  "message": "Bad request error"
}
```

#### NotFound

```json
// Status: 404
{
  "statusCode": 404,
  "message": "Not found error"
}
```

#### InternalServer

\*Internal server errors will also be logged and cause the program to exit.

```json
// Status: 500
{
  "statusCode": 500,
  "message": "Internal server error"
}
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->

[Go]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://go.dev/
[Gin]: https://img.shields.io/badge/Gin-008ECF.svg?style=for-the-badge&logo=Gin&logoColor=white
[Gin-url]: https://gin-gonic.com/
