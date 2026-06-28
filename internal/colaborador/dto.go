package colaborador

import (
	"strings"

	"github.com/google/uuid"
)

func (c CreateRequest) toRepositoryCreate() ColaboradorRepository {
	c.Nome = strings.TrimSpace(c.Nome)
	c.Matricula = strings.TrimSpace(c.Matricula)
	c.Setor = strings.TrimSpace(c.Setor)
	c.Tag = strings.TrimSpace(c.Tag)

	return ColaboradorRepository{
		ID:        uuid.New(),
		Nome:      c.Nome,
		Matricula: c.Matricula,
		Tag:       c.Tag,
		Setor:     c.Setor,
		Ativo:     true,
	}
}

func (c ColaboradorRepository) toColaborador() Colaborador {
	return Colaborador(c)
}
