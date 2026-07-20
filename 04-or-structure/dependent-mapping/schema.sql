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

-- Dependent Mapping: tabla sin identidad propia para el dominio,
-- solo se accede vía orders.
CREATE TABLE order_lines (
    order_id   BIGINT REFERENCES orders(id),
    product_id BIGINT REFERENCES products(id),
    quantity   INT
);

-- Semilla: el pedido 12 que actualiza el ejemplo, con una línea vieja
-- que el Update reemplaza.
INSERT INTO customers (id, name, email)
VALUES (7, 'Ana', 'ana@example.com');
SELECT setval('customers_id_seq', 7);

INSERT INTO products (id, name, price_cents) VALUES
    (42, 'Teclado mecánico', 159900),
    (43, 'Mouse inalámbrico', 25900);
SELECT setval('products_id_seq', 43);

INSERT INTO orders (id, customer_id, status, total_cents, version)
VALUES (12, 7, 'nuevo', 25900, 1);
SELECT setval('orders_id_seq', 12);

INSERT INTO order_lines (order_id, product_id, quantity)
VALUES (12, 43, 1);
