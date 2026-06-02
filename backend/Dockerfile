# Stage 1: Build
FROM maven:3.8.4-openjdk-17-slim AS builder

WORKDIR /app

COPY pom.xml .
RUN mvn dependency:go-offline -B

COPY src ./src
RUN mvn clean package -DskipTests

# Stage 2: Run
FROM openjdk:17-slim

WORKDIR /app

COPY --from=builder /app/target/*.jar app.jar

# Добавляем wait-for-it для ожидания БД
ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

EXPOSE 8080

# Ждём PostgreSQL и запускаем приложение
ENTRYPOINT ["/wait-for-it.sh", "postgresql:5432", "--", "java", "-jar", "app.jar"]
