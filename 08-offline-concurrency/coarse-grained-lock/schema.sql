CREATE TABLE customers (
    id   BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE products (
    id          BIGSERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    price_cents BIGINT NOT NULL
);

-- La versión vive SOLO en la raíz del agregado
CREATE TABLE orders (
    id          BIGSERIAL PRIMARY KEY,
    customer_id BIGINT NOT NULL REFERENCES customers(id),
    status      TEXT NOT NULL DEFAULT 'pending',
    total_cents BIGINT NOT NULL DEFAULT 0,
    version     BIGINT NOT NULL DEFAULT 1
);

-- Las líneas no tienen versión propia: las cubre la de la raíz
CREATE TABLE order_lines (
    order_id   BIGINT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    product_id BIGINT NOT NULL REFERENCES products(id),
    quantity   INT NOT NULL,
    PRIMARY KEY (order_id, product_id)
);

-- Datos semilla: el agregado que ambos usuarios van a editar
INSERT INTO customers (id, name) VALUES (7, 'Carla Gómez');
INSERT INTO products (id, name, price_cents) VALUES
    (1, 'Teclado mecánico', 18950),
    (5, 'Mouse inalámbrico', 8000);
INSERT INTO orders (id, customer_id, status, total_cents, version)
VALUES (12, 7, 'pending', 45900, 3);
INSERT INTO order_lines (order_id, product_id, quantity) VALUES
    (12, 1, 2),
    (12, 5, 1);

SELECT setval('customers_id_seq', 100);
SELECT setval('products_id_seq', 100);
SELECT setval('orders_id_seq', 100);
