-- Создание таблицы users
CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       email VARCHAR(255) NOT NULL UNIQUE,
                       password VARCHAR(255) NOT NULL,
                       created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                       updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
                       deleted_at TIMESTAMP DEFAULT NULL
);

-- Создание таблицы tasks с внешним ключом
CREATE TABLE tasks (
                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       name VARCHAR(255) NOT NULL,
                       is_done BOOLEAN DEFAULT FALSE,
                       user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                       created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                       updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
                       deleted_at TIMESTAMP DEFAULT NULL
);

