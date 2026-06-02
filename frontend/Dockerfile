# Stage 1: Build
FROM node:18-alpine AS builder

WORKDIR /app

# Копируем package.json
COPY package*.json ./
RUN npm ci || npm install

# Копируем исходники
COPY . .

# Собираем
RUN npm run build -- --prod || npm run build

# Stage 2: Serve
FROM nginx:alpine

COPY --from=builder /app/dist /usr/share/nginx/html

# Конфиг nginx
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
    } \
    location /report/ { \
        proxy_pass http://sausage-backend-report:8081/; \
        proxy_set_header Host $host; \
    } \
}' > /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
