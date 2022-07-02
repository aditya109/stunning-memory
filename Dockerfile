FROM openjdk:18-alpine AS builder
RUN apk --no-cache add maven
WORKDIR /java/app
COPY . /java/app/
RUN mvn clean package
LABEL Name=stunningmemory Version=0.0.1
EXPOSE 3000
ENTRYPOINT ["java","-jar","/java/app/target/stunning-memory-0.0.1-SNAPSHOT.jar"]




