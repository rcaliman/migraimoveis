package sistemaNovo

import (
	"gorm.io/gorm"
	"time"
)

type Cliente struct {
	gorm.Model
	Nome           string    `json:"nome""`
	DataNascimento time.Time `json:"data_nascimento"`
	Ci             string    `json:"ci"`
	Cpf            string    `json:"cpf"`
	Telefone1      string    `json:"telefone_1"`
	Telefone2      string    `json:"telefone_2"`
}
