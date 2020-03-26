FROM golang:latest 
RUN mkdir /go-rest-api
ADD . /go-rest-api/ 
WORKDIR /go-rest-api
RUN go build -o main . 
EXPOSE 8081
CMD ["/go-rest-api/main"]