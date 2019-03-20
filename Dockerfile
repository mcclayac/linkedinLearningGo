FROM golang:1.8

RUN mkdir /go/src/booksvc
RUN chdir /go/src/booksvc
WORKDIR /go/src/booksvc
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...


CMD ["booksvc"]
