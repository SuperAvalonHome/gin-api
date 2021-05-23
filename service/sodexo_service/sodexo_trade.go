package sodexo_service

import (
	"github.com/SuperAvalonHome/gin-api/models"
)

type SodexoTrade struct {
	Id         		int
	PayId       	int
	CardId  			int
	CardType 			string
	OrderId 			string
	TradeNo      	int64
	OutTradeNo		string
	TradeAmount 	float32
	TradeStatus		string
	TradeTime			string
}

func (t *SodexoTrade) Add() error {
	sodexoTrade := map[string]interface{}{
		"id":          		t.Id,
		"pay_id":         t.PayId,
		"card_id":        t.CardId,
		"card_type":      t.CardType,
		"order_id":     	t.OrderId,
		"trade_no":      	t.TradeNo,
		"out_trade_no": 	t.OutTradeNo,
		"trade_amount": 	t.TradeAmount,
		"trade_status":   t.TradeStatus,
		"trade_time":    	t.TradeTime,
	}
	if err := models.AddSodexoTrade(sodexoTrade); err != nil {
		return err
	}
	return nil
}

func (t *SodexoTrade) Get() (*models.SodexoTrade, error) {
	var sodexoTrade *models.SodexoTrade
	sodexoTrade, err := models.GetSodexoTrade(t.OrderId)
	if err != nil {
		return nil, err
	}
	return sodexoTrade, nil
}
