CREATE TABLE products (
    id          BIGSERIAL PRIMARY KEY,
    name        TEXT,
    price_cents BIGINT
);

-- Muchos-a-muchos: la relación vive en una tabla intermedia.
CREATE TABLE tags (
    id   BIGSERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE product_tags (            -- Association Table Mapping
    product_id BIGINT REFERENCES products(id),
    tag_id     BIGINT REFERENCES tags(id),
    PRIMARY KEY (product_id, tag_id)   -- sin id propio: no es una entidad
);

-- Semilla: el producto 42 y los tags que usa el ejemplo.
INSERT INTO products (id, name, price_cents)
VALUES (42, 'Teclado mecánico', 159900);
SELECT setval('products_id_seq', 42);

INSERT INTO tags (id, name) VALUES
    (1, 'oferta'),
    (2, 'hogar'),
    (3, 'gadgets');
SELECT setval('tags_id_seq', 3);
