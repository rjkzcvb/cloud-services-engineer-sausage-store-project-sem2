# Stage 1: Build
FROM node:18-alpine AS builder

WORKDIR /app

COPY package*.json ./
RUN npm ci

COPY . .
RUN npm run build --prod

# Stage 2: Serve with nginx
FROM nginx:alpine

COPY --from=builder /app/dist/sausage-store /usr/share/nginx/html

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
    } \
    location /report/ { \
        proxy_pass http://sausage-backend-report:8081/; \
        proxy_set_header Host $host; \
    } \
}' > /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
