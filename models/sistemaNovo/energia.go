package sistemaNovo

import (
	"gorm.io/gorm"
	"time"
)

type Energia struct {
	gorm.Model
	Data       time.Time
	Relogio1   int     `json:"relogio_1"`
	Relogio2   int     `json:"relogio_2"`
	Relogio3   int     `json:"relogio_3"`
	ValorKwh   float64 `json:"valor_kwh"`
	ValorConta float64 `json:"valor_conta"`
}
