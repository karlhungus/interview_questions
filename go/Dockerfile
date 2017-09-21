FROM go:1.9

RUN apt-get update && apt-get install -y \
    --no-install-recommends && \
    rm -rf /var/lib/apt/lists/*

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

ADD . /usr/src/app

WORKDIR /usr/src/app/
CMD ["ruby", "palindrome_client.rb", "-x", "-v"]
