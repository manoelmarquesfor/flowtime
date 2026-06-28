-- +goose Up
CREATE TABLE colaboradores (
    id UUID PRIMARY KEY,
    matricula VARCHAR(15) NOT NULL UNIQUE,
    tag_id VARCHAR(15) NOT NULL UNIQUE,
    nome VARCHAR(255) NOT NULL,
    setor VARCHAR(30) NOT NULL,
    ativo BOOLEAN NOT NULL DEFAULT TRUE,
    dt_criacao TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_matricula ON colaboradores(matricula);
CREATE INDEX idx_tag_id ON colaboradores(tag_id);
-- +goose Down
DROP TABLE colaboradores;