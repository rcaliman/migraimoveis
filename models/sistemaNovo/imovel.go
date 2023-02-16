package sistemaNovo

import "gorm.io/gorm"

type Imovel struct {
	gorm.Model
	Tipo         string  `json:"tipo"`
	Numero       string  `json:"numero"`
	Local        string  `json:"local"`
	Cliente_id   int     `json:"cliente"`
	ValorAluguel float64 `json:"valor_aluguel"`
	Observacao   string  `json:"observacao"`
	DiaBase      int     `json:"dia_Base"`
}
