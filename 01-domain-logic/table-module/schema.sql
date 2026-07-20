CREATE TABLE customers (
  id    BIGSERIAL PRIMARY KEY,
  name  TEXT,
  email TEXT
);

CREATE TABLE orders (
  id          BIGSERIAL PRIMARY KEY,
  customer_id BIGINT REFERENCES customers(id),
  status      TEXT,
  total_cents BIGINT,
  version     BIGINT,
  created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Datos semilla para el ejemplo.
INSERT INTO customers (id, name, email) VALUES
  (42, 'Ada López', 'ada@example.com');

INSERT INTO orders (customer_id, status, total_cents, version, created_at) VALUES
  (42, 'confirmado', 21000, 1, now() - interval '60 days'),
  (42, 'pendiente',   4500, 1, now() - interval '45 days'), -- vieja: la marcará MarkOverdue(30)
  (42, 'pendiente',  12000, 1, now() - interval '5 days'),
  (42, 'cancelado',   9900, 1, now() - interval '10 days'); -- no cuenta para el total

SELECT setval('customers_id_seq', 42);
