package main

import (
	"emailn/internal/domain/campaign"

	"github.com/go-playground/validator/v10"
)

func main() {
	campaign := campaign.Campaign{}
	validate := validator.New()
	err := validate.Struct(campaign)
	println(err)
	if err == nil {
		println("Nenhum Erro")
	} else {
		validationErrors := err.(validator.ValidationErrors)
		for _, v := range validationErrors {
			println(v.Error())
		}
	}
}
