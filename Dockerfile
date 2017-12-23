FROM golang:latest

WORKDIR /app
ADD . /app


RUN go get -u github.com/gin-gonic/gin
RUN go get github.com/pilu/fresh

CMD fresh

EXPOSE 80