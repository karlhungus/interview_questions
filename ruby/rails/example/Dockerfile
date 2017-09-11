FROM ruby:2.4.1

# Install Rails dependencies
RUN apt-get update && apt-get install -y nodejs --no-install-recommends && rm -rf /var/lib/apt/lists/*
RUN apt-get update && apt-get install -y \
      # mysql-client \ 
      # postgresql-client \
    sqlite3 \
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

# Initialize log
RUN cat /dev/null > /usr/src/app/log/production.log
RUN chmod -R a+w /usr/src/app/log

EXPOSE 3000


WORKDIR /usr/src/app/
CMD ["bin/rails", "server", "-b", "0.0.0.0"]
