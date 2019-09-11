FROM golang:1.13-alpine as build-env

RUN mkdir /user
WORKDIR /user
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix
cgo -o /go/bin/user

FROM scratch
COPY --from=build-env /go/bin/user /go/bin/user
ENTRYPOINT ["/go/bin/user"]