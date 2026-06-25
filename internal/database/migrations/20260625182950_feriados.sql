-- +goose Up
CREATE TABLE IF NOT EXISTS feriados (
    data DATE PRIMARY KEY,
    descricao VARCHAR(50) NOT NULL
);
CREATE INDEX idx_feriados_data ON feriados(data);
WITH RECURSIVE years(y) AS (
    SELECT 2026
    UNION ALL
    SELECT y + 1
    FROM years
    WHERE y < 2045
),
holidays(day, descricao) AS (
    SELECT '01-01',
        'Ano Novo'
    UNION ALL
    SELECT '04-21',
        'Tiradentes'
    UNION ALL
    SELECT '05-01',
        'Dia do Trabalho'
    UNION ALL
    SELECT '09-07',
        'Independência do Brasil'
    UNION ALL
    SELECT '10-12',
        'Nossa Senhora Aparecida'
    UNION ALL
    SELECT '11-02',
        'Finados'
    UNION ALL
    SELECT '11-15',
        'Proclamação da República'
    UNION ALL
    SELECT '11-20',
        'Dia Nacional de Zumbi e da Consciência Negra'
    UNION ALL
    SELECT '12-25',
        'Natal'
)
INSERT
    OR IGNORE INTO feriados(data, descricao)
SELECT printf('%04d-%s', y, day),
    descricao
FROM years
    CROSS JOIN holidays;
-- +goose Down
DROP TABLE IF EXISTS feriados;