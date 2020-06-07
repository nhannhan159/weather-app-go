FROM golang:1.13

# env
ENV APP_NAME=weather-app
ENV BASE_DIR=/
ENV PORT=8080
EXPOSE 8080

# util scripts
WORKDIR /scripts
COPY script/wait-for-it/wait-for-it.sh ./
RUN chmod +x ./wait-for-it.sh

# main build
VOLUME /logs/weather-app
WORKDIR /go/src/github.com/nhannhan159/weather-app-go
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install

CMD ["/go/bin/weather-app-go"]
