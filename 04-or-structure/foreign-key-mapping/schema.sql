CREATE TABLE customers (
    id    BIGSERIAL PRIMARY KEY,
    name  TEXT,
    email TEXT
);

-- Uno-a-muchos: el hijo apunta al padre.
CREATE TABLE orders (
    id          BIGSERIAL PRIMARY KEY,
    customer_id BIGINT REFERENCES customers(id), -- Foreign Key Mapping
    status      TEXT,
    total_cents BIGINT,
    version     BIGINT
);

-- Semilla: Ana (id 7) con dos pedidos existentes.
INSERT INTO customers (id, name, email)
VALUES (7, 'Ana', 'ana@example.com');
SELECT setval('customers_id_seq', 7);

INSERT INTO orders (customer_id, status, total_cents, version) VALUES
    (7, 'pagado',  25900, 1),
    (7, 'enviado', 51800, 1);
