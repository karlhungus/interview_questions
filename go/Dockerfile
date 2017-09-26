FROM golang:1.9

RUN apt-get update && apt-get install -y \
    --no-install-recommends && \
    rm -rf /var/lib/apt/lists/*

RUN mkdir /app
WORKDIR /app

ADD . /app

WORKDIR /app/
CMD ["go", "test", "./..."]
