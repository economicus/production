package models

import "time"

// RegisterRequest is a type for register
type RegisterRequest struct {
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Name     string    `json:"name"`
	Nickname string    `json:"nickname"`
	Birth    time.Time `json:"birth"`
}

// QuantRequest is a type for create quant model
type QuantRequest struct {
	Name                string     `json:"name"`
	QuantID             uint       `json:"quant_id"`
	MainSectors         []string   `json:"main_sectors"`
	NetRevenue          IntPair    `json:"net_revenue"`
	NetRevenueRate      DoublePair `json:"net_revenue_rate"`
	NetProfit           IntPair    `json:"net_profit"`
	NetProfitRate       DoublePair `json:"net_profit_rate"`
	DERatio             DoublePair `json:"de_ratio"`
	Per                 DoublePair `json:"per"`
	Psr                 DoublePair `json:"psr"`
	Pbr                 DoublePair `json:"pbr"`
	Pcr                 DoublePair `json:"pcr"`
	Activities          Activities `json:"activities"`
	DividendYield       DoublePair `json:"dividend_yield"`
	DividendPayoutRatio DoublePair `json:"dividend_payout_ratio"`
	Roa                 DoublePair `json:"roa"`
	Roe                 DoublePair `json:"roe"`
	MarketCap           IntPair    `json:"market_cap"`
	StartDate           time.Time  `time_format:"2006-01-02T15:04:05.000Z" json:"start_date"`
	EndDate             time.Time  `time_format:"2006-01-02T15:04:05.000Z" json:"end_date"`
}

func (qr *QuantRequest) ToQuantOption() *QuantOption {
	return &QuantOption{
		QuantID:             qr.QuantID,
		MainSectors:         NewMainSectors(qr.QuantID, qr.MainSectors),
		NetRevenue:          qr.NetRevenue,
		NetRevenueRate:      qr.NetRevenueRate,
		NetProfit:           qr.NetProfit,
		NetProfitRate:       qr.NetProfitRate,
		DERatio:             qr.DERatio,
		Per:                 qr.Per,
		Psr:                 qr.Psr,
		Pbr:                 qr.Pbr,
		Pcr:                 qr.Pcr,
		Activities:          qr.Activities,
		DividendYield:       qr.DividendYield,
		DividendPayoutRatio: qr.DividendPayoutRatio,
		Roa:                 qr.Roa,
		Roe:                 qr.Roe,
		MarketCap:           qr.MarketCap,
		StartDate:           qr.StartDate,
		EndDate:             qr.EndDate,
	}
}
