-- Esquema mínimo para el ejemplo de Lazy Load (dominio tienda)

CREATE TABLE customers (
    id    BIGSERIAL PRIMARY KEY,
    name  TEXT NOT NULL,
    email TEXT NOT NULL
);

CREATE TABLE products (
    id          BIGSERIAL PRIMARY KEY,
    name        TEXT   NOT NULL,
    price_cents BIGINT NOT NULL
);

CREATE TABLE orders (
    id          BIGSERIAL PRIMARY KEY,
    customer_id BIGINT NOT NULL REFERENCES customers (id),
    status      TEXT   NOT NULL,
    total_cents BIGINT NOT NULL,
    version     BIGINT NOT NULL DEFAULT 1
);

CREATE TABLE order_lines (
    order_id   BIGINT NOT NULL REFERENCES orders (id),
    product_id BIGINT NOT NULL REFERENCES products (id),
    quantity   INT    NOT NULL
);

-- Datos semilla: el pedido 12 con sus líneas
INSERT INTO customers (id, name, email) VALUES
    (3, 'Carla Robles', 'carla@example.com');

INSERT INTO products (id, name, price_cents) VALUES
    (1, 'Teclado mecánico', 89900),
    (2, 'Mouse inalámbrico', 39900),
    (5, 'Base para laptop', 25100);

INSERT INTO orders (id, customer_id, status, total_cents, version) VALUES
    (12, 3, 'pending', 129900, 1);

INSERT INTO order_lines (order_id, product_id, quantity) VALUES
    (12, 1, 1),
    (12, 2, 1),
    (12, 5, 2);

-- Ajustamos las secuencias para que los nuevos ids no choquen con la semilla
SELECT setval('customers_id_seq', 100, true);
SELECT setval('products_id_seq', 100, true);
SELECT setval('orders_id_seq', 100, true);
