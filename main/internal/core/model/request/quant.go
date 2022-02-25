package request

import (
	"main/internal/core/model/quant"
	"time"
)

type QuantC struct {
	Name                string           `json:"name"`
	QuantID             uint             `json:"quant_id"`
	MainSectors         []string         `json:"main_sectors"`
	NetRevenue          quant.IntPair    `json:"net_revenue"`
	NetRevenueRate      quant.DoublePair `json:"net_revenue_rate"`
	NetProfit           quant.IntPair    `json:"net_profit"`
	NetProfitRate       quant.DoublePair `json:"net_profit_rate"`
	DERatio             quant.DoublePair `json:"de_ratio"`
	Per                 quant.DoublePair `json:"per"`
	Psr                 quant.DoublePair `json:"psr"`
	Pbr                 quant.DoublePair `json:"pbr"`
	Pcr                 quant.DoublePair `json:"pcr"`
	Activities          quant.Activities `json:"activities"`
	DividendYield       quant.DoublePair `json:"dividend_yield"`
	DividendPayoutRatio quant.DoublePair `json:"dividend_payout_ratio"`
	Roa                 quant.DoublePair `json:"roa"`
	Roe                 quant.DoublePair `json:"roe"`
	MarketCap           quant.IntPair    `json:"market_cap"`
	StartDate           time.Time        `time_format:"2006-01-02T15:04:05.000Z" json:"start_date"`
	EndDate             time.Time        `time_format:"2006-01-02T15:04:05.000Z" json:"end_date"`
}

type QuantE struct {
	QuantID     uint   `json:"quant_id"`
	Active      bool   `json:"active"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
