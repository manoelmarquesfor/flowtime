package usuario

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

func (dto UsuarioCreate) UsuarioCreateRepository(newPassword string) UsuarioCreateRepository {
	dto.Regra = strings.ToUpper(dto.Regra)
	dto.Nome = strings.TrimSpace(dto.Nome)
	dto.Email = strings.TrimSpace(dto.Email)

	return UsuarioCreateRepository{
		ID:        uuid.New(),
		Nome:      dto.Nome,
		Email:     dto.Email,
		Password:  newPassword,
		Regra:     dto.Regra,
		Ativo:     true,
		DtCreated: time.Now(),
	}
}

func (dto UsuarioRepository) ToUsuario() Usuario {
	return Usuario{
		ID:    dto.ID,
		Nome:  dto.Nome,
		Email: dto.Email,
		Regra: dto.Regra,
		Ativo: dto.Ativo,
	}
}

func (dto UsuarioCreateRepository) ToUsuario() Usuario {
	return Usuario{
		ID:    dto.ID,
		Nome:  dto.Nome,
		Email: dto.Email,
		Regra: dto.Regra,
		Ativo: dto.Ativo,
	}
}
