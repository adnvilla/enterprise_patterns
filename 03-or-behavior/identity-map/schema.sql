-- Esquema mínimo para el ejemplo de Identity Map (dominio tienda)

CREATE TABLE customers (
    id    BIGSERIAL PRIMARY KEY,
    name  TEXT NOT NULL,
    email TEXT NOT NULL
);

CREATE TABLE orders (
    id          BIGSERIAL PRIMARY KEY,
    customer_id BIGINT NOT NULL REFERENCES customers (id),
    status      TEXT   NOT NULL,
    total_cents BIGINT NOT NULL,
    version     BIGINT NOT NULL DEFAULT 1
);

-- Datos semilla: el pedido 12 que carga el ejemplo
INSERT INTO customers (id, name, email) VALUES
    (3, 'Carla Robles', 'carla@example.com');

INSERT INTO orders (id, customer_id, status, total_cents, version) VALUES
    (12, 3, 'pending', 129900, 1);

-- Ajustamos las secuencias para que los nuevos ids no choquen con la semilla
SELECT setval('customers_id_seq', 100, true);
SELECT setval('orders_id_seq', 100, true);
