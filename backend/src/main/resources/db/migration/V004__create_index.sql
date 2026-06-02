CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_orders_date_created_status ON orders(date_created, status);
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_order_product_composite ON order_product(order_id, product_id);
