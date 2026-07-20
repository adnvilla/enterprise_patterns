CREATE TABLE customers (
    id    BIGSERIAL PRIMARY KEY,
    name  TEXT,
    email TEXT
);

-- Embedded Value: los campos de Address se aplanan en customers.
ALTER TABLE customers
    ADD COLUMN address_street TEXT,
    ADD COLUMN address_city   TEXT,
    ADD COLUMN address_zip    TEXT;

-- Semilla: Ana (id 7) con su dirección aplanada en columnas.
INSERT INTO customers (id, name, email, address_street, address_city, address_zip)
VALUES (7, 'Ana', 'ana@example.com', 'Av. Reforma 123', 'CDMX', '06600');
SELECT setval('customers_id_seq', 7);
