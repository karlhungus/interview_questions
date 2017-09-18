FROM ruby:2.4.1

RUN apt-get update && apt-get install -y \
    --no-install-recommends && \
    rm -rf /var/lib/apt/lists/*

# throw errors if Gemfile has been modified since Gemfile.lock
RUN bundle config --global frozen 1

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

ADD Gemfile /usr/src/app/
ADD Gemfile.lock /usr/src/app/
RUN bundle install --system

ADD . /usr/src/app

WORKDIR /usr/src/app/
CMD ["ruby", "palindrome_client.rb", "-x", "-v"]
