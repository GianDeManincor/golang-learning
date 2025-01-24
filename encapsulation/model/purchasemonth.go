package model

import "time"

type PurchaseOfTheMonth struct {
	Date   time.Time
	Market string
	itens  []string
}

func Initializer() PurchaseOfTheMonth {
	return PurchaseOfTheMonth{
		Date:   time.Now(),
		Market: "Serv lar",
		itens:  []string{"Arroz", "Feijão", "Macarrão", "Cebola", "Chocolate"},
	}
}
