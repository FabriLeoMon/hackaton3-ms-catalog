package models

type Product struct {
    Title       string   `json:"title" bigquery:"title"`
    Description string   `json:"description" bigquery:"description"`
    Images      []string `json:"images" bigquery:"images"`
    Price       float32  `json:"price" bigquery:"price"`
}

type Item struct {
    Stock   int32   `json:"stock"`
    Product Product `json:"product"`
}