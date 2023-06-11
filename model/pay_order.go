package model

import (
	"go4pay/pkg/db"
	"time"

	log "go4pay/pkg/logger"
)

type PayOrder struct {
	Id         int64   `xorm:"'id' pk autoincr"`
	CreateTime string  `xorm:"'create_time'"`
	UpdateTime string  `xorm:"'update_time'"`
	DeleteTime string  `xorm:"'delete_time'"`
	OrderId    string  `xorm:"'order_id'"`
	Channel    string  `xorm:"'channel'"`
	Name       string  `xorm:"'name'"`
	Comment    string  `xorm:"'comment'"`
	Value      float64 `xorm:"'value'"`
	ValueType  string  `xorm:"'value_type'"`
	Identifier string  `xorm:"'identifier'"`
	Data       string  `xorm:"'data'"`
	Status     string  `xorm:"'status'"`
	Fee        float64 `json:"fee"`
	ValueIn    float64 `json:"valueIn"`

	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updateAt"`
}

func CreateOrUpdatePayOrder(req *PayOrder) (err error) {

	var order PayOrder
	has, err := db.GetDB().Table("pay_order").Where("order_id = ?", req.OrderId).Get(&order)
	if err != nil {
		log.Errorf("get pay order err:%v", err)
		return
	}

	if has {
		// update
		_, err = db.GetDB().Table("pay_order").Cols("status", "updated_at").Where("order_id = ?", req.OrderId).Update(req)
		if err != nil {
			log.Errorf("update pay order err:%v", err)
			return
		}
	} else {
		// insert
		_, err = db.GetDB().Table("pay_order").Insert(req)
		if err != nil {
			log.Errorf("create pay order err:%v", err)
			return
		}
	}
	return

}
