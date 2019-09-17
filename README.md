# jtechchallenge

## Requirements

* mac OSX / Linux
* Docker
* Docker Compose

## Microservices

* MySQL
* Consul (orchestration)
* Micro
* Petstore API Service (Go/go-micro)
* Pet Service (Go/go-micro)
* Store Service (Go/go-micro)
* User Service (Go/go-micro)

## Usage

### Starting

You should only need to run ....

`docker-compose up -d`

I'm going to provide a basic overview of the workflow I use to test against Swagger/OpenAPI specs.

* Import the `swagger.json` doc into Postman
* Edit the collection and set the baseURL variable to `http://<docker physical addr>:8080/petstore` 
* Go to the `/user/login` endpoint and generate an auth token from the following details
  * username: `admin`
  * password: `admin`
* Add the following to the headers of the other requests:
  * `Authorization: Bearer <token>`
  
## Trouble-shooting

### Problems with the `quic-go` package when building, trying to verify.

There is a problem with `v0.20.0` at the moment with the repos checksum. This should be fixed in their `v0.21.0` release.

In the mean time, open `go.sum` and remove any lines pertaining to `github.com/lucas-clemente/quic-go v0.12.0/go.mod h1:UXJJPE4RfFef/xPO5wQm0tITK8gNfqwTxjbE7s3Vb8s=` then run `docker-compose up --build` again. 
  