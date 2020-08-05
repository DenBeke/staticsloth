FROM golang:latest AS build

WORKDIR /random_work_dir

# first download dependencies
COPY go.mod /random_work_dir
COPY go.sum /random_work_dir
RUN go mod download

# then copy source code
COPY / /random_work_dir


RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /staticsloth ./cmd/staticsloth


FROM alpine:latest

WORKDIR /

COPY --from=build /staticsloth /staticsloth/

WORKDIR /staticsloth

RUN chmod +x ./staticsloth

EXPOSE 1234

CMD ["./staticsloth"]
