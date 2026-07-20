CREATE TABLE customers (
    id   BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE orders (
    id          BIGSERIAL PRIMARY KEY,
    customer_id BIGINT NOT NULL REFERENCES customers(id),
    status      TEXT NOT NULL DEFAULT 'pending',
    total_cents BIGINT NOT NULL DEFAULT 0,
    version     BIGINT NOT NULL DEFAULT 1
);

CREATE TABLE locks (
    resource    TEXT PRIMARY KEY,      -- p. ej. 'order:12'
    owner       TEXT NOT NULL,         -- la sesión/usuario dueño del lock
    acquired_at TIMESTAMPTZ NOT NULL   -- para detectar locks abandonados
);

-- Datos semilla: el pedido por el que compiten los dos empleados
INSERT INTO customers (id, name) VALUES (7, 'Carla Gómez');
INSERT INTO orders (id, customer_id, status, total_cents, version)
VALUES (12, 7, 'pending', 45900, 3);

SELECT setval('customers_id_seq', 100);
SELECT setval('orders_id_seq', 100);
