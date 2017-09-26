FROM golang:1.9

RUN apt-get update && apt-get install -y \
    --no-install-recommends && \
    rm -rf /var/lib/apt/lists/*


ADD . /go
WORKDIR /go

RUN ["go", "get", "./..."]
RUN ["go", "build", "./src/github.com/karlhungus/..."]
CMD ["go", "test", "./src/github.com/karlhungus/..."]
