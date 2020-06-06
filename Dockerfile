FROM golang:1.13

ENV APP_NAME=weather-app
ENV BASE_DIR=/
ENV PORT=8080
EXPOSE 8080

VOLUME /logs/weather-app
WORKDIR /go/src/github.com/nhannhan159/weather-app-go
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install
RUN ls
RUN ls /go
RUN ls /go/bin

ENTRYPOINT ["/go/bin/weather-app-go"]
