CREATE TABLE products (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    title text NOT NULL,
    description text NOT NULL,
    price bigint NOT NULL,
    status text NOT NULL,
    city text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW()
);

