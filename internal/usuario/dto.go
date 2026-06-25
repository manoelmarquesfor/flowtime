package usuario

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

func (dto UsuarioCreate) UsuarioCreateRepository(newPassword string) UsuarioCreateRepository {
	dto.Perfil = strings.ToUpper(dto.Perfil)
	dto.Nome = strings.TrimSpace(dto.Nome)
	dto.Email = strings.TrimSpace(dto.Email)

	return UsuarioCreateRepository{
		ID:        uuid.New(),
		Nome:      dto.Nome,
		Email:     dto.Email,
		Password:  newPassword,
		Perfil:    dto.Perfil,
		Ativo:     true,
		DtCreated: time.Now(),
	}
}

func (dto UsuarioRepository) ToUsuario() Usuario {
	return Usuario(dto)
}

func (dto UsuarioCreateRepository) ToUsuario() Usuario {
	return Usuario{
		ID:     dto.ID,
		Nome:   dto.Nome,
		Email:  dto.Email,
		Perfil: dto.Perfil,
		Ativo:  dto.Ativo,
	}
}
