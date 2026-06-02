FROM nginx:alpine

# Удаляем дефолтную страницу
RUN rm -rf /usr/share/nginx/html/*

# Создаем простую HTML страницу с API вызовами
RUN echo '<!DOCTYPE html>
<html>
<head>
    <title>Sausage Store</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 0; padding: 20px; background: #f5f5f5; }
        .container { max-width: 1200px; margin: 0 auto; }
        h1 { color: #d32f2f; }
        .product { background: white; border-radius: 8px; padding: 15px; margin: 10px; display: inline-block; width: 200px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
        .product h3 { margin: 0 0 10px 0; }
        .product p { color: #666; }
        button { background: #d32f2f; color: white; border: none; padding: 8px 16px; border-radius: 4px; cursor: pointer; }
        button:hover { background: #b71c1c; }
        .cart { background: white; padding: 15px; border-radius: 8px; margin-top: 20px; }
        .status { padding: 10px; margin: 10px 0; border-radius: 4px; }
        .status.healthy { background: #4caf50; color: white; }
        .status.error { background: #f44336; color: white; }
    </style>
</head>
<body>
    <div class="container">
        <h1>🌭 Sausage Store</h1>
        <div id="status" class="status"></div>
        <div id="products">Loading...</div>
        <div class="cart">
            <h2>Shopping Cart</h2>
            <div id="cart">Empty</div>
        </div>
    </div>

    <script>
        async function checkHealth() {
            try {
                const res = await fetch("/api/health");
                if (res.ok) {
                    document.getElementById("status").innerHTML = "✅ Backend Connected";
                    document.getElementById("status").className = "status healthy";
                } else {
                    throw new Error();
                }
            } catch(e) {
                document.getElementById("status").innerHTML = "❌ Backend Not Available";
                document.getElementById("status").className = "status error";
            }
        }

        async function loadProducts() {
            try {
                const res = await fetch("/api/products");
                if (!res.ok) throw new Error();
                const products = await res.json();
                const html = products.map(p => `
                    <div class="product">
                        <h3>${p.name}</h3>
                        <p>${p.price} ₽</p>
                        <button onclick="addToCart(${p.id})">Add to Cart</button>
                    </div>
                `).join("");
                document.getElementById("products").innerHTML = html;
            } catch(e) {
                document.getElementById("products").innerHTML = "<p>Cannot load products. Make sure backend is running.</p>";
            }
        }

        async function addToCart(id) {
            try {
                const res = await fetch("/api/cart", {
                    method: "POST",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify({ product_id: id, quantity: 1 })
                });
                if (res.ok) {
                    alert("Added to cart!");
                    loadCart();
                }
            } catch(e) {
                alert("Error adding to cart");
            }
        }

        async function loadCart() {
            try {
                const res = await fetch("/api/cart");
                if (!res.ok) throw new Error();
                const cart = await res.json();
                if (cart.length === 0) {
                    document.getElementById("cart").innerHTML = "Empty";
                } else {
                    document.getElementById("cart").innerHTML = cart.map(i => 
                        `<div>${i.product_name} x ${i.quantity} = ${i.total} ₽</div>`
                    ).join("");
                }
            } catch(e) {
                document.getElementById("cart").innerHTML = "Cannot load cart";
            }
        }

        checkHealth();
        loadProducts();
        loadCart();
    </script>
</body>
</html>' > /usr/share/nginx/html/index.html

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
