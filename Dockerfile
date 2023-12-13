FROM golang:1.21 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build ./cmd/workouts

FROM golang:1.21 AS app

WORKDIR /app

COPY --from=build /app/workouts .

EXPOSE 3000

CMD ["./workouts"]
