FROM golang:latest

WORKDIR /app
ADD ./src /app


RUN go get -u github.com/gin-gonic/gin
RUN go get github.com/pilu/fresh
RUN go get github.com/stretchr/testify/assert
RUN go get github.com/gin-contrib/cors

CMD fresh

EXPOSE 80