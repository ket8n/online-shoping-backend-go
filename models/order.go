package models

type Order struct {
    ID        string `uri:"id" binding:"required"`
    ItemID    int    `form:"item_id"`
    Quantity  int    `form:"quantity"`
}
