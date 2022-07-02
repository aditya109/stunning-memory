FROM openjdk:18-alpine
RUN apk --no-cache add maven
WORKDIR /java/app
COPY . .
RUN mvn clean package
ENTRYPOINT ["java","-jar","/target/stunning-memory 0.0.1-SNAPSHOT.jar"]
