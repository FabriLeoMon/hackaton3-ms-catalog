package handlers

import (
    "github.com/Set-Forget/Hackaton3/ms-catalog/repositories"
    "github.com/gofiber/fiber/v2"
)

type CatalogHandler struct {
    productRepo *repositories.ProductRepository
}

func NewCatalogHandler(productRepo *repositories.ProductRepository) *CatalogHandler {
    return &CatalogHandler{
        productRepo: productRepo,
    }
}

func (h *CatalogHandler) GetCatalog(c *fiber.Ctx) error {
    items, err := h.productRepo.GetCatalog()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to fetch catalog",
        })
    }

    return c.JSON(items)
}