FROM golang:1.13-alpine as build-env

RUN mkdir /petstore
WORKDIR /petstore
COPY go.mod .
COPY go.sum .

RUN apk add --no-cache git mercurial

ENV GOSUMDB off
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/petstore

FROM scratch
COPY --from=build-env /go/bin/petstore /go/bin/petstore
ENTRYPOINT ["/go/bin/petstore"]