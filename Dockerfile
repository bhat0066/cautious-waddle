FROM golang:latest 
RUN mkdir /app && go get -u github.com/gorilla/mux
ADD . /app/ 
WORKDIR /app 
EXPOSE 8080
RUN go build -o main . 
CMD ["/app/main"]