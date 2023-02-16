package main

import (
	"migraalugueis/databases/sistemaNovoDB"
	"migraalugueis/databases/sistemaVelhoDB"
	"migraalugueis/models/sistemaNovo"
	"migraalugueis/models/sistemaVelho"
	"strings"
	"time"
)

func parseDate(t time.Time) time.Time {
	ano := t.Year()
	dia := t.Day()
	data := time.Date(ano, t.Month(), dia, 0, 0, 0, 0, time.Local)
	return data
}

func stripDocs(doc string) string {
	doc = strings.ReplaceAll(doc, "/", "")
	doc = strings.ReplaceAll(doc, "-", "")
	doc = strings.ReplaceAll(doc, ".", "")
	return doc
}

func main() {
	sistemaVelhoDB.ConectaDBAlugueis()
	sistemaNovoDB.ConectaDBImoveis()

	var sistemaVelhoClientes []sistemaVelho.Al_Cliente
	sistemaVelhoDB.DB.Find(&sistemaVelhoClientes)

	var sistemaNovoCliente sistemaNovo.Cliente
	for _, cliente := range sistemaVelhoClientes {
		sistemaNovoCliente = sistemaNovo.Cliente{
			Nome:           cliente.Nome,
			DataNascimento: parseDate(cliente.DataNasc),
			Ci:             cliente.Ci,
			Cpf:            stripDocs(cliente.Cpf),
			Telefone1:      cliente.Tel1,
			Telefone2:      cliente.Tel2,
		}
		sistemaNovoDB.DB.Create(&sistemaNovoCliente)

	}

	var sistemaVelhoImoveis []sistemaVelho.Al_Imoveis
	sistemaVelhoDB.DB.Find(&sistemaVelhoImoveis)

	for _, imovel := range sistemaVelhoImoveis {
		var cliente sistemaNovo.Cliente
		sistemaNovoDB.DB.Where("nome like ?", imovel.Cliente).First(&cliente)
		sistemaNovoImovel := sistemaNovo.Imovel{
			Tipo:         strings.ToLower(imovel.Tipo),
			Numero:       imovel.Numero,
			Local:        imovel.Local,
			Cliente_id:   int(cliente.ID),
			ValorAluguel: imovel.ValorAluguel,
			Observacao:   imovel.Observacao,
			DiaBase:      imovel.DiaBase,
		}
		sistemaNovoDB.DB.Create(&sistemaNovoImovel)
	}

	var sistemaVelhoEnergia []sistemaVelho.Al_Energia
	sistemaVelhoDB.DB.Order("data").Find(&sistemaVelhoEnergia)

	for _, energia := range sistemaVelhoEnergia {
		sistemaNovoEnergia := sistemaNovo.Energia{
			Data:       parseDate(energia.Data),
			Relogio1:   energia.Relogio1,
			Relogio2:   energia.Relogio2,
			Relogio3:   energia.Relogio3,
			ValorKwh:   energia.ValorKwh,
			ValorConta: energia.ValorConta,
		}
		sistemaNovoDB.DB.Create(&sistemaNovoEnergia)
	}
}
