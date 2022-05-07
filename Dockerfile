FROM golang:1.15
WORKDIR /go/src/teste
COPY . .
RUN GOOS=linux go build -ldflags="-s -w"
#RUN GOOS=linux go run 
CMD ["./teste"]