FROM golang:1.13-alpine as build-env

RUN mkdir /pet
WORKDIR /pet
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix
cgo -o /go/bin/pet

FROM scratch
COPY --from=build-env /go/bin/pet /go/bin/pet
ENTRYPOINT ["/go/bin/pet"]