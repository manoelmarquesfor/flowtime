package feriado

import "time"

type Response struct {
	Data      string `json:"data"`
	Descricao string `json:"descricao"`
}

type CreateRequest struct {
	Data      string `json:"data"      validate:"required"`
	Descricao string `json:"descricao" validate:"required,max=255"`
}

type RepositoryModel struct {
	Data      time.Time `db:"data"`
	Descricao string    `db:"descricao"`
}

type DeleteResponse struct {
	Detail string `json:"detail"`
}

type CreateResponse struct {
	Detail string `json:"detail"`
}
