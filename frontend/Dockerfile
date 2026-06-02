# Stage 1: Build
FROM node:16-alpine AS builder

ENV NODE_OPTIONS="--max-old-space-size=4096 --openssl-legacy-provider"

WORKDIR /app

# Копируем package.json
COPY package*.json ./

RUN npm install --no-optional --no-audit --no-fund

COPY . .

# Проверяем что у нас в package.json
RUN cat package.json | grep -E '"build"' || echo "No build script"

# Собираем по-разному в зависимости от проекта
RUN npm run build -- --prod 2>/dev/null || \
    npm run build 2>/dev/null || \
    npx ng build --prod 2>/dev/null || \
    echo "No build command found"

# Stage 2: Serve
FROM nginx:alpine

# Ищем где лежит dist
COPY --from=builder /app/dist /usr/share/nginx/html || \
     COPY --from=builder /app/build /usr/share/nginx/html || \
     COPY --from=builder /app/public /usr/share/nginx/html

RUN echo "Frontend files:" && ls -la /usr/share/nginx/html/

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
