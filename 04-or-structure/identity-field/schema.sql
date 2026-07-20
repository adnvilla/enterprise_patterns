CREATE TABLE customers (
    id    BIGSERIAL PRIMARY KEY,
    name  TEXT,
    email TEXT
);

-- Clave sustituta: la genera la base, no el negocio.
-- Unicidad por tabla: el id 7 puede existir también en products.
CREATE TABLE orders (
    id          BIGSERIAL PRIMARY KEY, -- Identity Field
    customer_id BIGINT REFERENCES customers(id),
    status      TEXT,
    total_cents BIGINT,
    version     BIGINT
);

-- Semilla: la clienta 7 que referencia el ejemplo.
INSERT INTO customers (id, name, email)
VALUES (7, 'Ana', 'ana@example.com');
SELECT setval('customers_id_seq', 7);
