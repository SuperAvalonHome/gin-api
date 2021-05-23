package models

import (
	"github.com/jinzhu/gorm"
)

type SodexoUser struct {
	Id         	int				`json:"id" gorm:"type:int(11);primary_key"`
	Uid        	int				`json:"uid" gorm:"type:int(11);"`
	Uuid  		 	string		`json:"uuid" gorm:"type:varchar(64);"`
	Username 		string		`json:"username" gorm:"type:varchar(64);"`
	Mobile      int64			`json:"mobile" gorm:"type:char(11);"`
	CardInfo		string		`json:"card_info" gorm:"type:text;"`
	DefaultCard string		`json:"default_card" gorm:"type:varchar(64);"`
	IdSuffix		string		`json:"id_suffix" gorm:"type:char(4);"`
	BindTime		string		`json:"bind_time" gorm:"type:datetime;"`
	BindStatus	int				`json:"bind_status" gorm:"type:int(11);"`
}

type SodexoTrade struct {
	Id         		int				`json:"id" gorm:"type:int(11);primary_key"`
	PayId       	int				`json:"pay_id" gorm:"type:int(11);"`
	CardId  			int				`json:"card_id" gorm:"type:int(11);"`
	CardType 			string		`json:"card_type" gorm:"type:char(1);"`
	TradeNo      	string 		`json:"trade_no" gorm:"type:varchar(32);"`
	OutTradeNo		string		`json:"out_trade_no" gorm:"type:varchar(32);"`
	TradeAmount 	float32		`json:"trade_amount" gorm:"type:decimal(10,2);"`
	TradeStatus		string		`json:"trade_status" gorm:"type:varchar(10);"`
	TradeTime			string		`json:"trade_time" gorm:"type:datetime;"`
}


type SodexoRefund struct {
	Id         		int				`json:"id" gorm:"type:int(11);primary_key"`
	OrderId 			string 		`json:"order_id" gorm:"type:varchar(32);"`
	OutRefundNo		string 		`json:"out_refund_no" gorm:"type:varchar(32);"`
	TradeNo      	string 		`json:"trade_no" gorm:"type:varchar(32);"`
	TradeType     string 		`json:"trade_type" gorm:"type:varchar(32);"`
	Barcode     	string 		`json:"barcode" gorm:"type:varchar(32);"`
	TradeAmount 	float32 	`json:"trade_amount" gorm:"type:varchar(32);"`
	RefundAmount 	float32 	`json:"refund_amount" gorm:"type:varchar(32);"`
	RefundStatus	string 		`json:"refund_status" gorm:"type:varchar(32);"`
	CreateTime		string 		`json:"create_time" gorm:"type:datetime;"`
	TradeTime			string 		`json:"trade_time" gorm:"type:datetime;"`
}

func AddSodexoUser(data map[string]interface{}) error {
	sodexoUser := SodexoUser{
		Id:         		data["id"].(int),
		Uid:         		data["uid"].(int),
		Uuid:          	data["uuid"].(string),
		Username:       data["username"].(string),
		Mobile:     		data["mobile"].(int64),
		CardInfo:       data["card_info"].(string),
		DefaultCard: 		data["default_card"].(string),
		IdSuffix: 			data["id_suffix"].(string),
		BindTime: 			data["bind_time"].(string),
		BindStatus: 		data["bind_status"].(int),
	}

	if err := db.Create(&sodexoUser).Error; err != nil {
		return err
	}

	return nil
}

func AddSodexoTrade(data map[string]interface{}) error {
	sodexoTrade := SodexoTrade{
		Id:         			data["id"].(int),
		PayId:         		data["pay_id"].(int),
		CardId:          	data["card_id"].(int),
		CardType:       	data["card_type"].(string),
		TradeNo:     			data["trade_no"].(string),
		OutTradeNo:       data["out_trade_no"].(string),
		TradeAmount: 			data["trade_amount"].(float32),
		TradeStatus: 			data["trade_status"].(string),
		TradeTime: 				data["trade_time"].(string),
	}

	if err := db.Create(&sodexoTrade).Error; err != nil {
		return err
	}

	return nil
}


func AddSodexoRefund(data map[string]interface{}) error {
	sodexoRefund := SodexoRefund{
		Id:         				data["id"].(int),
		OrderId:         		data["order_id"].(string),
		OutRefundNo:        data["out_refund_no"].(string),
		TradeNo:       			data["trade_no"].(string),
		TradeType:     			data["trade_type"].(string),
		Barcode:       			data["barcode"].(string),
		TradeAmount: 				data["trade_amount"].(float32),
		RefundStatus: 			data["refund_status"].(string),
		CreateTime: 				data["create_time"].(string),
		TradeTime: 					data["trade_time"].(string),
	}

	if err := db.Create(&sodexoRefund).Error; err != nil {
		return err
	}

	return nil
}


// GetSodexoUser Get a single article based on ID
func GetSodexoUser(id int) (*SodexoUser, error) {
	var sodexoUser SodexoUser
	err := db.Where("id = ? ", id).First(&sodexoUser).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &sodexoUser, nil
}

// GetSodexoUser Get a single article based on ID
func GetSodexoTrade(order_id string) (*SodexoTrade, error) {
	var sodexoTrade SodexoTrade
	err := db.Where("order_id = ? ", order_id).First(&sodexoTrade).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &sodexoTrade, nil
}


// GetSodexoUser Get a single article based on ID
func GetSodexoRefund(order_id string) (*SodexoRefund, error) {
	var sodexoRefund SodexoRefund
	err := db.Where("order_id = ? ", order_id).First(&sodexoRefund).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &sodexoRefund, nil
}
