CREATE TABLE
    IF NOT EXISTS roles (
        role_id SERIAL PRIMARY KEY NOT NULL,
        role CHARACTER VARYING(10) NOT NULL
    );

CREATE TABLE
    IF NOT EXISTS features (
        feature_id SERIAL PRIMARY KEY NOT NULL,
        feature CHARACTER VARYING(32) NOT NULL,
        is_deleted BOOLEAN NOT NULL
    );

CREATE TABLE
    IF NOT EXISTS users (
        user_id SERIAL PRIMARY KEY NOT NULL,
        name CHARACTER VARYING(32) NOT NULL,
        phone CHARACTER VARYING(32) NOT NULL,
        address CHARACTER VARYING(128) NOT NULL,
        email CHARACTER VARYING(32) NOT NULL,
        password_hash CHARACTER VARYING(128) NOT NULL,
        is_deleted BOOLEAN NOT NULL
    );

CREATE TABLE
    IF NOT EXISTS employees (
        user_id BIGINT NOT NULL REFERENCES users (user_id),
        salary FLOAT NOT NULL,
        position CHARACTER VARYING(32) NOT NULL,
        KPI SMALLINT NOT NULL,
        PRIMARY KEY (user_id)
    );

CREATE TABLE
    IF NOT EXISTS users_roles (
        user_id BIGINT NOT NULL REFERENCES users (user_id),
        role_id BIGINT NOT NULL REFERENCES roles (role_id),
        PRIMARY KEY (user_id, role_id)
    );

CREATE TABLE
    IF NOT EXISTS vendors (
        vendor_id SERIAL PRIMARY KEY NOT NULL,
        phone CHARACTER VARYING(32) NOT NULL,
        org_name CHARACTER VARYING(32) NOT NULL,
        INN CHARACTER VARYING(10) NOT NULL,
        OGRN CHARACTER VARYING(13) NOT NULL,
        address CHARACTER VARYING(128) NOT NULL,
        is_deleted BOOLEAN NOT NULL
    );

CREATE TABLE
    IF NOT EXISTS goods (
        good_id SERIAL PRIMARY KEY NOT NULL,
        name CHARACTER VARYING(32) NOT NULL,
        description CHARACTER VARYING(128) NOT NULL,
        price FLOAT NOT NULL,
        count INT NOT NULL,
        is_deleted BOOLEAN NOT NULL
    );

CREATE TABLE
    IF NOT EXISTS goods_vendors (
        good_id BIGINT NOT NULL REFERENCES goods (good_id),
        vendor_id BIGINT NOT NULL REFERENCES vendors (vendor_id),
        PRIMARY KEY (good_id, vendor_id)
    );

CREATE TABLE
    IF NOT EXISTS goods_features (
        good_id BIGINT NOT NULL REFERENCES goods (good_id),
        feature_id BIGINT NOT NULL REFERENCES features (feature_id),
        PRIMARY KEY (good_id, feature_id)
    );

CREATE TABLE
    IF NOT EXISTS favorites (
        user_id BIGINT NOT NULL REFERENCES users (user_id),
        good_id BIGINT NOT NULL REFERENCES goods (good_id),
        PRIMARY KEY (user_id, good_id)
    );

CREATE TABLE
    IF NOT EXISTS baskets (
        user_id BIGINT NOT NULL REFERENCES users (user_id),
        good_id BIGINT NOT NULL REFERENCES goods (good_id),
        count INT NOT NULL,
        PRIMARY KEY (user_id, good_id)
    );

CREATE TABLE
    IF NOT EXISTS orders (
        order_id SERIAL PRIMARY KEY NOT NULL,
        delivery_type CHARACTER VARYING(32) NOT NULL,
        delivery_time TIMESTAMP NOT NULL,
        order_time TIMESTAMP NOT NULL,
        total_price FLOAT NOT NULL,
        canceled BOOLEAN NOT NULL
    );

CREATE TABLE
    IF NOT EXISTS users_orders (
        user_id BIGINT NOT NULL REFERENCES users (user_id),
        order_id BIGINT NOT NULL REFERENCES orders (order_id),
        PRIMARY KEY (user_id, order_id)
    );

CREATE TABLE
    IF NOT EXISTS orders_goods (
        order_id BIGINT NOT NULL REFERENCES orders (order_id),
        good_id BIGINT NOT NULL REFERENCES goods (good_id),
        count INT NOT NULL,
        PRIMARY KEY (order_id, good_id)
    );