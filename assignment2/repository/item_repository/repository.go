package item_repository

import (
	"assignment2/entity"
	"assignment2/pkg/errs"

	"gorm.io/gorm"
)

type ItemRepository interface {
	GetItemByItemCode(itemCode string, txs ...*gorm.DB) (*entity.Item, errs.MessageErr)
	UpdateItemByItemCode(itemCode string, payload *entity.Item, txs ...*gorm.DB) (*entity.Item, errs.MessageErr)
}
