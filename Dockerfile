FROM golang:1.13

ARG APP_NAME=weather-app
ENV APP_NAME=${APP_NAME}
ENV BASE_DIR=/
ENV PORT=8080
EXPOSE 8080

VOLUME /logs/${APP_NAME}
WORKDIR /go/src/github.com/nhannhan159/weather-app-go
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ${APP_NAME} .

ENTRYPOINT ["/go/bin/weather-app"]
