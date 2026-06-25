-- +goose Up
CREATE TABLE usuario (
    id UUID PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    sessao_id UUID,
    regra VARCHAR(50) NOT NULL,
    ativo BOOLEAN NOT NULL DEFAULT TRUE,
    dt_criacao TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_usuario_email ON usuario(email);
CREATE INDEX idx_usuario_sessao_id ON usuario(sessao_id);
INSERT INTO usuario (id, nome, email, password, regra, ativo)
VALUES (
        '439b7929-1906-44bb-b8af-158a7dd5d6de',
        'Administrador',
        'admin@admin.com',
        '$2a$10$8MVfJuY3ndbIs5wOmT5dFePplpObMh/ebg.IsMTB9XxEKFqnZZwLu',
        'ADMIN',
        TRUE
    );
-- +goose Down
DROP TABLE usuario;