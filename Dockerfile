FROM golang:latest
# RUN echo "docker start"

# ENV GOPATH /go
WORKDIR /app
ADD . /app

# CMD which go
# CMD echo $GOPATH

RUN go get -u github.com/gin-gonic/gin
RUN go get github.com/pilu/fresh
# RUN go install github.com/pilu/fresh


# CMD export PATH=$PATH:$GOPATH

# RUN echo "docker end"
# CMD ls -la /usr/local/go/pkg
# CMD ls -la ../bin/
CMD fresh
# CMD ["go", "run", "main.go"]