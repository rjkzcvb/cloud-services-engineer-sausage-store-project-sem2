# Stage 1: Build dependencies
FROM node:14-alpine AS deps

WORKDIR /app
COPY package*.json ./

# Устанавливаем зависимости отдельно
RUN npm install --no-optional --no-audit --no-fund

# Stage 2: Build
FROM node:14-alpine AS builder

ENV NODE_OPTIONS="--max-old-space-size=4096"

WORKDIR /app

# Копируем зависимости из предыдущего этапа
COPY --from=deps /app/node_modules ./node_modules
COPY . .

# Собираем
RUN npm run build -- --prod || npm run build

# Stage 3: Run
FROM nginx:alpine

COPY --from=builder /app/dist /usr/share/nginx/html

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
