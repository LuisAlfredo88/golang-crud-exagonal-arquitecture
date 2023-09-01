package entity

type Product struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Weight      float32 `json:"weight"`
	Color       string  `json:"color"`
}
