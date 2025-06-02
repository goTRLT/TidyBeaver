FROM golang:latest

RUN mkdir /src/
WORKDIR /src
COPY . /src
RUN go get -d -v ./...
RUN go install -v ./...

RUN go build -o app .
CMD ["./app" ]