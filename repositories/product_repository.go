package repositories

import (
    "database/sql"
    "encoding/json"
    "log"

    "github.com/Set-Forget/Hackaton3/ms-catalog/models"
)

type ProductRepository struct {
    db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
    return &ProductRepository{
        db: db,
    }
}

func (r *ProductRepository) GetCatalog() ([]models.Item, error) {
    query := `
        SELECT i.stock, 
                json_build_object(
                   'title', p.title,
                   'description', p.description,
                   'images', p.images,
                   'price', p.price
               ) as product
        FROM public.stock i
        JOIN public."Products" p ON i.product_id = p.id
    `

    rows, err := r.db.Query(query)
    if err != nil {
        log.Printf("Error executing query: %v", err)
        return nil, err
    }
    defer rows.Close()

    var items []models.Item
    for rows.Next() {
        var item models.Item
        var productJSON []byte
        
        err := rows.Scan(&item.Stock, &productJSON)
        if err != nil {
            log.Printf("Error scanning row: %v", err)
            continue
        }

        err = json.Unmarshal(productJSON, &item.Product)
        if err != nil {
            log.Printf("Error unmarshaling product: %v", err)
            continue
        }

        items = append(items, item)
    }

    return items, nil
}