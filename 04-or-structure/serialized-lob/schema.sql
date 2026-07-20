CREATE TABLE products (
    id          BIGSERIAL PRIMARY KEY,
    name        TEXT,
    price_cents BIGINT
);

-- Serialized LOB: un grafo pequeño como JSONB dentro de products.
ALTER TABLE products
    ADD COLUMN specs JSONB;

-- Semilla: el producto 42 cuyas especificaciones guarda el ejemplo.
INSERT INTO products (id, name, price_cents)
VALUES (42, 'Teclado mecánico', 159900);
SELECT setval('products_id_seq', 42);
