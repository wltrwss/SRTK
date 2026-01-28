package items

import "time"

type Position struct {
	Barcode         string     `json:"barcode"`
	Name            string     `json:"name"`
	UnitMeasurement string     `json:"unit_measurement"`
	Quantity        float64    `json:"quantity"`
	Price_buy       float64    `json:"price_buy"`
	Price_sell      float64    `json:"price_sell"`
	Date            *time.Time `json:"date"`
}

type Scanner struct {
	Status bool `json:"status"`
}

type Signal struct {
	Status bool `json:"status"`
}
