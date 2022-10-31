FROM golang:1.16-alpine

WORKDIR /view
COPY view/*.tmpl /view/

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./server ./
COPY . ./
EXPOSE 9090
RUN go build -o /crud

CMD ["/crud"]