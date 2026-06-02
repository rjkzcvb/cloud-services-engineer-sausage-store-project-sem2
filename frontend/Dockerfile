# Stage 1: Build
FROM node:16-alpine AS builder

WORKDIR /app

# Копируем package.json
COPY package*.json ./

# Устанавливаем зависимости
RUN npm install

# Копируем исходники
COPY . .

# Добавляем переменную окружения для OpenSSL legacy provider
ENV NODE_OPTIONS=--openssl-legacy-provider

# Собираем приложение
RUN npm run build -- --prod

# Stage 2: Serve
FROM nginx:alpine

# Копируем собранный фронтенд
COPY --from=builder /app/dist /usr/share/nginx/html

# Создаем конфиг nginx
RUN echo 'server { \
    listen 80; \
    server_name localhost; \
    root /usr/share/nginx/html; \
    index index.html; \
    location / { \
        try_files $uri $uri/ /index.html; \
    } \
    location /api/ { \
        proxy_pass http://sausage-backend:8080/; \
        proxy_set_header Host $host; \
        proxy_set_header X-Real-IP $remote_addr; \
    } \
    location /report/ { \
        proxy_pass http://sausage-backend-report:8081/; \
        proxy_set_header Host $host; \
        proxy_set_header X-Real-IP $remote_addr; \
    } \
}' > /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
