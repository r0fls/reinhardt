package models

import (
	"github.com/r0fls/reinhardt/src/model"
)

func Models() model.Model {
	M := model.NewModel("Customers")
	M["Customers"].IntegerField("Age")
	M["Customers"].CharacterField("Name")
	M.AddModel("Invoices")
	M["Invoices"].IntegerField("Amount")
	//model.Connect("postgres", "postgres", "localhost", "235711", "test")
	return M
}
