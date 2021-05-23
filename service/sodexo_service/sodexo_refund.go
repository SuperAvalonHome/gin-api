package sodexo_service

import (
	"github.com/SuperAvalonHome/gin-api/models"
)

type SodexoRefund struct {
	Id         		int
	OrderId 			string
	OutTradeNo		string
	TradeNo      	string
	TradeType     string
	Barcode     	string
	TradeAmount 	float32
	RefundAmount 	float32
	RefundStatus	string
	CreateTime		string
	TradeTime			string
}

func (r *SodexoRefund) Add() error {
	sodexoRefund := map[string]interface{}{
		"id":          		r.Id,
		"order_id":       r.OrderId,
		"out_refund_no":  r.OutTradeNo,
		"trade_no":     	r.TradeNo,
		"trade_type":     r.TradeType,
		"barcode": 				r.Barcode,
		"trade_amount": 	r.TradeAmount,
		"refund_amount":  r.RefundAmount,
		"refund_status":  r.RefundStatus,
		"create_time":  	r.CreateTime,
		"trade_time":  		r.TradeTime,
	}
	if err := models.AddSodexoRefund(sodexoRefund); err != nil {
		return err
	}
	return nil
}

func (r *SodexoRefund) Get() (*models.SodexoRefund, error) {
	var sodexoRefund *models.SodexoRefund
	sodexoRefund, err := models.GetSodexoRefund(r.OrderId)
	if err != nil {
		return nil, err
	}
	return sodexoRefund, nil
}
