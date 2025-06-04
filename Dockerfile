FROM golang:latest

RUN mkdir /app
WORKDIR /app
COPY . /app
RUN go get -d -v ./...
RUN go install -v ./...

RUN go build -o app ./cmd
CMD ["./app" ]