package schema

import "time"

type GroupMenu struct {
	Id          uint16 `gorm:"primarykey;AUTO_INCREMENT;"`
	Name        string `gorm:"type:varchar(50);unique;"`
	Description string `gorm:"type:text;"`
	Active      bool   `gorm:"default:true"`
	Type        uint8  `gorm:"default:0;"`
	Menu        []Menu `gorm:"foreignKey:GroupId;"`
}

func (GroupMenu) TableName() string {
	return "group_menu"
}

type Menu struct {
	Id          uint16 `gorm:"primarykey;AUTO_INCREMENT"`
	Name        string `gorm:"type:varchar(50);unique;"`
	Description string `gorm:"type:text;"`
	GroupId     uint16
	Price       float64 `gorm:"type:decimal;"`
	MaxDisc     float32 `gorm:"type:float(2);"`
	MixDisc     float32 `gorm:"type:float(2);"`
	Picture     string  `gorm:"type:varchar(255)"`
	Active      bool
}

func (Menu) TableName() string {
	return "menu"
}

type Transaction struct {
	Id        uint16 `gorm:"primaryKey;AUTO_INCREMENT"`
	InvoiceId string `gorm:"type:varchar(50);unique"`
	TableId   uint8
	Table     TableTransaction `gorm:"foreignKey:TableId;"`
	Date      time.Time
	Void      bool  `gorm:"default:false"`
	Type      uint8 `gorm:"default:0"`
	ClientId  uint16
	UserId    string            `gorm:"type:varchar(50)"`
	Item      []TransactionItem `gorm:"foreignKey:InvoiceRefer;references:InvoiceId;"`
}

func (Transaction) TableName() string {
	return "transaction"
}

type TransactionItem struct {
	Id           uint16 `gorm:"primary_key;AUTO_INCREMENT"`
	InvoiceRefer string
	MenuId       uint16
	Menu         Menu    `gorm:"foreignKey:MenuId;"`
	Price        int64   `gorm:"type:decimal;"`
	Void         bool    `gorm:"default:false"`
	Disc         float32 `gorm:"type:float(2);default:0"`
	Date         time.Time
}

func (TransactionItem) TableName() string {
	return "transaction_item"
}

type TableTransaction struct {
	Id          uint8  `gorm:"primary_key;AUTO_INCREMENT;"`
	Name        string `gorm:"type:varchar(20);unique"`
	Description string `gorm:"type:text"`
	Active      bool   `gorm:"default:true"`
	Type        uint8  `gorm:"default:0"`
}

func (TableTransaction) TableName() string {
	return "table_transactions"
}
