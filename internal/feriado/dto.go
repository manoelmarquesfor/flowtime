package feriado

import "time"

func (f RepositoryModel) ToResponse() Response {
	return Response{
		Data:      f.Data.Format(time.DateOnly),
		Descricao: f.Descricao,
	}
}
