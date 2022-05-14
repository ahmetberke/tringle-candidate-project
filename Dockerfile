# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR /app

CMD ["ping", "google.com", "-c", "1"]

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /tringle-candidate-project

EXPOSE 8080

ENTRYPOINT [ "/tringle-candidate-project" ]