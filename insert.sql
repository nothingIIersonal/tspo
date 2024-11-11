-- Роли
INSERT INTO
    roles (role)
VALUES
    ('admin'),
    ('user'),
    ('manager');

-- Характеристики товаров
INSERT INTO
    features (feature, isDeleted)
VALUES
    ('Водонепроницаемый', false),
    ('Противоударный', false),
    ('Беспроводной', false),
    ('Складной', false);

-- Пользователи
INSERT INTO
    users (
        name,
        phone,
        address,
        email,
        passwordHash,
        isDeleted
    )
VALUES
    (
        'Иван Петров',
        '+79001234567',
        'ул. Ленина 1, кв. 10',
        'ivan@mail.ru',
        'hash1',
        false
    ),
    (
        'Мария Сидорова',
        '+79009876543',
        'ул. Пушкина 5, кв. 20',
        'maria@mail.ru',
        'hash2',
        false
    ),
    (
        'Алексей Иванов',
        '+79007894561',
        'ул. Гагарина 3, кв. 15',
        'alex@mail.ru',
        'hash3',
        false
    ),
    (
        'Игорь Фомин',
        '+79037425928',
        'ул. Иваныкина 2, кв. 40',
        'igor@mail.ru',
        'hash4',
        false
    );

-- Сотрудники
INSERT INTO
    employees (userId, salary, position, KPI)
VALUES
    (1, 70000.00, 'Менеджер', 85),
    (2, 45000.00, 'Консультант', 90),
    (4, 60000.00, 'Администратор', 95);

-- Связь пользователей и ролей
INSERT INTO
    usersRoles (userId, roleId)
VALUES
    (1, 1),
    (2, 2),
    (3, 2),
    (4, 3);

-- Поставщики
INSERT INTO
    vendors (phone, orgName, INN, OGRN, address, isDeleted)
VALUES
    (
        '+74951234567',
        'ООО Техника',
        '1234567890',
        '1234567890123',
        'ул. Промышленная 10',
        false
    ),
    (
        '+74959876543',
        'ООО Электроника',
        '0987654321',
        '3210987654321',
        'ул. Заводская 5',
        false
    );

-- Товары
INSERT INTO
    goods (name, description, price, count, isDeleted)
VALUES
    (
        'Смартфон',
        'Современный смартфон с большим экраном',
        15000.00,
        10,
        false
    ),
    (
        'Ноутбук',
        'Мощный ноутбук для работы',
        45000.00,
        5,
        false
    ),
    (
        'Планшет',
        'Компактный планшет',
        20000.00,
        8,
        false
    );

-- Связь товаров и поставщиков
INSERT INTO
    goodsVendors (goodId, vendorId)
VALUES
    (1, 1),
    (2, 1),
    (3, 2);

-- Связь товаров и характеристик
INSERT INTO
    goodsFeatures (goodId, featureId)
VALUES
    (1, 1),
    (1, 2),
    (2, 2);

-- Избранное
INSERT INTO
    favorites (userId, goodId)
VALUES
    (1, 1),
    (2, 2);

-- Корзины
INSERT INTO
    baskets (userId, goodId)
VALUES
    (1, 1),
    (1, 2),
    (1, 3),
    (2, 2),
    (3, 3);

-- Заказы
INSERT INTO
    orders (
        deliveryType,
        deliveryTime,
        orderTime,
        totalPrice,
        canceled
    )
VALUES
    (
        'Курьер',
        '2023-12-25 12:00:00',
        '2023-12-24 15:30:00',
        15000.00,
        false
    ),
    (
        'Самовывоз',
        '2023-12-26 14:00:00',
        '2023-12-24 16:45:00',
        45000.00,
        false
    );

-- Связь пользователей и заказов
INSERT INTO
    usersOrders (userId, orderId)
VALUES
    (1, 1),
    (2, 2);

-- Товары в заказах
INSERT INTO
    ordersGoods (orderId, goodId, count)
VALUES
    (1, 1, 1),
    (2, 2, 1);