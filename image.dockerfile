FROM golang:1.16-alpine

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./server ./server
COPY ./model ./model
COPY ./server/main.go ./
COPY ./view ./view
EXPOSE 9090
RUN go build -o /crud

CMD ["/crud"]