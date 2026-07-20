-- Clientes de la tienda (referenciada por orders)
CREATE TABLE customers (
  id   BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL
);

-- Tabla de pedidos de la tienda
CREATE TABLE orders (
  id          BIGSERIAL PRIMARY KEY,
  customer_id BIGINT REFERENCES customers(id),
  status      TEXT,
  total_cents BIGINT,
  version     BIGINT
);

-- Datos semilla: el cliente 42 que usa el ejemplo
INSERT INTO customers (id, name) VALUES (42, 'Ada Lovelace');
SELECT setval('customers_id_seq', 42);
