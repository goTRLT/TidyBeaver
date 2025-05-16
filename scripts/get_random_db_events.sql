CREATE OR REPLACE FUNCTION get_random_db_events(n integer)
RETURNS TABLE(
    level text,
    "column" text,
    "constraint" text,
    datatype text,
    table_name text,
    schema text,
    errcode text,
    detail text
) AS $$
DECLARE
    levels text[] := ARRAY['DEBUG', 'LOG', 'INFO', 'NOTICE', 'WARNING', 'EXCEPTION'];
    columns text[] := ARRAY['user_id', 'email', 'created_at', 'order_id', 'amount', 'status'];
    constraints text[] := ARRAY['users_pkey', 'orders_fkey', 'unique_email', 'check_amount_positive', 'not_null_status'];
    datatypes text[] := ARRAY['integer', 'text', 'timestamp', 'numeric', 'boolean'];
    tables text[] := ARRAY['users', 'orders', 'products', 'payments', 'sessions'];
    schemas text[] := ARRAY['public', 'sales', 'auth', 'inventory'];
    errcodes text[] := ARRAY['23505', '23503', '23502', '22003', '22007', '42601'];
    details text[] := ARRAY[
        'Key (email)=(test@example.com) already exists.',
        'Null value in column "status" violates not-null constraint.',
        'Insert or update on table "orders" violates foreign key constraint.',
        'Value too large for type numeric.',
        'Invalid input syntax for type timestamp.',
        'Syntax error at or near "FROM".'
    ];
    i integer;
BEGIN
    FOR i IN 1..n LOOP
        level := levels[1 + floor(random() * array_length(levels, 1))::int];
        "column" := columns[1 + floor(random() * array_length(columns, 1))::int];
        "constraint" := constraints[1 + floor(random() * array_length(constraints, 1))::int];
        datatype := datatypes[1 + floor(random() * array_length(datatypes, 1))::int];
        table_name := tables[1 + floor(random() * array_length(tables, 1))::int];
        schema := schemas[1 + floor(random() * array_length(schemas, 1))::int];
        errcode := errcodes[1 + floor(random() * array_length(errcodes, 1))::int];
        detail := details[1 + floor(random() * array_length(details, 1))::int];
        RETURN NEXT;
    END LOOP;
END;
$$ LANGUAGE plpgsql;