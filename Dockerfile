FROM golang:1.15
WORKDIR /go/src/teste
COPY . .
#RUN GOOS=linux go build -ldflags="-s -w"
RUN go run 
CMD ["./teste"]