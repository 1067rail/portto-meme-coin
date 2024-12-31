FROM golang:1.23.4 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o ./main

FROM golang:1.23.4
WORKDIR /app

COPY --from=build /app/main ./

ENV SERVER_PORT=8080
EXPOSE 8080

CMD ["./main"]
