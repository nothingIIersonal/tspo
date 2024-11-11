CREATE TABLE
    IF NOT EXISTS roles (
        roleId SERIAL PRIMARY KEY,
        role CHARACTER VARYING(10) NOT NULL
    );

CREATE TABLE
    IF NOT EXISTS features (
        featureId SERIAL PRIMARY KEY,
        feature CHARACTER VARYING(32) NOT NULL,
        isDeleted BOOLEAN NOT NULL
    );

CREATE TABLE
    IF NOT EXISTS users (
        userId SERIAL PRIMARY KEY,
        name CHARACTER VARYING(32) NOT NULL,
        phone CHARACTER VARYING(32) NOT NULL,
        address CHARACTER VARYING(128) NOT NULL,
        email CHARACTER VARYING(32) NOT NULL,
        passwordHash CHARACTER VARYING(128) NOT NULL,
        isDeleted BOOLEAN NOT NULL
    );

CREATE TABLE
    IF NOT EXISTS employees (
        userId BIGINT NOT NULL REFERENCES users (userId),
        salary FLOAT NOT NULL,
        position CHARACTER VARYING(32) NOT NULL,
        KPI SMALLINT NOT NULL,
        PRIMARY KEY (userId)
    );

CREATE TABLE
    IF NOT EXISTS usersRoles (
        userId BIGINT NOT NULL REFERENCES users (userId),
        roleId BIGINT NOT NULL REFERENCES roles (roleId),
        PRIMARY KEY (userId, roleId)
    );

CREATE TABLE
    IF NOT EXISTS vendors (
        vendorId SERIAL PRIMARY KEY,
        phone CHARACTER VARYING(32) NOT NULL,
        orgName CHARACTER VARYING(32) NOT NULL,
        INN CHARACTER VARYING(10) NOT NULL,
        OGRN CHARACTER VARYING(13) NOT NULL,
        address CHARACTER VARYING(128) NOT NULL,
        isDeleted BOOLEAN NOT NULL
    );

CREATE TABLE
    IF NOT EXISTS goods (
        goodId SERIAL PRIMARY KEY,
        name CHARACTER VARYING(32) NOT NULL,
        description CHARACTER VARYING(128) NOT NULL,
        price FLOAT NOT NULL,
        count INT NOT NULL,
        isDeleted BOOLEAN NOT NULL
    );

CREATE TABLE
    IF NOT EXISTS goodsVendors (
        goodId BIGINT NOT NULL REFERENCES goods (goodId),
        vendorId BIGINT NOT NULL REFERENCES vendors (vendorId),
        PRIMARY KEY (goodId, vendorId)
    );

CREATE TABLE
    IF NOT EXISTS goodsFeatures (
        goodId BIGINT NOT NULL REFERENCES goods (goodId),
        featureId BIGINT NOT NULL REFERENCES features (featureId),
        PRIMARY KEY (goodId, featureId)
    );

CREATE TABLE
    IF NOT EXISTS favorites (
        userId BIGINT NOT NULL REFERENCES users (userId),
        goodId BIGINT NOT NULL REFERENCES goods (goodId),
        PRIMARY KEY (userId, goodId)
    );

CREATE TABLE
    IF NOT EXISTS baskets (
        basketId SERIAL,
        userId BIGINT NOT NULL REFERENCES users (userId),
        goodId BIGINT NOT NULL REFERENCES goods (goodId),
        PRIMARY KEY (basketId, userId)
    );

CREATE TABLE
    IF NOT EXISTS orders (
        orderId SERIAL PRIMARY KEY,
        deliveryType CHARACTER VARYING(32) NOT NULL,
        deliveryTime TIMESTAMP NOT NULL,
        orderTime TIMESTAMP NOT NULL,
        totalPrice FLOAT NOT NULL,
        canceled BOOLEAN NOT NULL
    );

CREATE TABLE
    IF NOT EXISTS usersOrders (
        userId BIGINT NOT NULL REFERENCES users (userId),
        orderId BIGINT NOT NULL REFERENCES orders (orderId),
        PRIMARY KEY (userId, orderId)
    );

CREATE TABLE
    IF NOT EXISTS ordersGoods (
        orderId BIGINT NOT NULL REFERENCES orders (orderId),
        goodId BIGINT NOT NULL REFERENCES goods (goodId),
        count INT NOT NULL,
        PRIMARY KEY (orderId, goodId)
    );