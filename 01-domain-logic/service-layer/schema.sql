CREATE TABLE customers (
  id    BIGSERIAL PRIMARY KEY,
  name  TEXT,
  email TEXT
);

CREATE TABLE products (
  id          BIGSERIAL PRIMARY KEY,
  name        TEXT,
  price_cents BIGINT
);

CREATE TABLE orders (
  id          BIGSERIAL PRIMARY KEY,
  customer_id BIGINT REFERENCES customers(id),
  status      TEXT,
  total_cents BIGINT,
  version     BIGINT
);

CREATE TABLE order_lines (
  order_id   BIGINT REFERENCES orders(id),
  product_id BIGINT REFERENCES products(id),
  quantity   INT
);

-- Datos semilla para el ejemplo.
INSERT INTO customers (id, name, email) VALUES
  (42, 'Ada López', 'ada@example.com');

INSERT INTO products (id, name, price_cents) VALUES
  (1, 'Playera Go',  4500),
  (2, 'Taza gopher', 2500),
  (3, 'Sudadera',   12000);

SELECT setval('customers_id_seq', 42);
SELECT setval('products_id_seq', 3);
