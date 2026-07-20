-- Estrategia A: Single Table Inheritance
-- Una tabla, columna discriminadora, campos de subtipo nullables.
CREATE TABLE products (
    id           BIGSERIAL PRIMARY KEY,
    type         TEXT NOT NULL,      -- 'physical' | 'digital'
    name         TEXT NOT NULL,
    price_cents  BIGINT NOT NULL,
    weight_grams INT,                -- solo physical (NULL en digital)
    stock        INT,                -- solo physical
    download_url TEXT,               -- solo digital (NULL en physical)
    file_bytes   BIGINT              -- solo digital
);

-- Estrategia B: Class Table Inheritance
-- Tabla por clase (base incluida), unidas por el mismo id.
CREATE TABLE products_base (
    id          BIGSERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    price_cents BIGINT NOT NULL
);
CREATE TABLE physical_products (
    id           BIGINT PRIMARY KEY REFERENCES products_base(id),
    weight_grams INT NOT NULL,
    stock        INT NOT NULL
);
CREATE TABLE digital_products (
    id           BIGINT PRIMARY KEY REFERENCES products_base(id),
    download_url TEXT NOT NULL,
    file_bytes   BIGINT NOT NULL
);

-- Estrategia C: Concrete Table Inheritance
-- Tabla por clase concreta; las columnas comunes se repiten.
CREATE TABLE physical_products_c (
    id           BIGSERIAL PRIMARY KEY,
    name         TEXT NOT NULL,      -- repetida
    price_cents  BIGINT NOT NULL,    -- repetida
    weight_grams INT NOT NULL,
    stock        INT NOT NULL
);
CREATE TABLE digital_products_c (
    id           BIGSERIAL PRIMARY KEY,
    name         TEXT NOT NULL,      -- repetida
    price_cents  BIGINT NOT NULL,    -- repetida
    download_url TEXT NOT NULL,
    file_bytes   BIGINT NOT NULL
);
