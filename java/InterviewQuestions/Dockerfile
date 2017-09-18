FROM openjdk 

RUN apt-get update && apt-get install -y maven --no-remove && rm -rf /var/lib/apt/lists*

WORKDIR /data
ADD . /data

CMD ["mvn", "clean", "test"]
