package models

import (
	"economicus/grpc/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"time"
)

type QuantOption struct {
	gorm.Model          `json:"-"`
	QuantID             uint         `gorm:"column:quant_id" json:"quant_id"`
	MainSectors         []MainSector `gorm:"constraint:OnDelete:CASCADE;" json:"main_sectors"`
	NetRevenue          IntPair      `gorm:"embedded;embeddedPrefix:net_revenue_" json:"net_revenue"`
	NetRevenueRate      DoublePair   `gorm:"embedded;embeddedPrefix:net_revenue_rate_" json:"net_revenue_rate"`
	NetProfit           IntPair      `gorm:"embedded;embeddedPrefix:net_profit_" json:"net_profit"`
	NetProfitRate       DoublePair   `gorm:"embedded;embeddedPrefix:net_profit_rate_" json:"net_profit_rate"`
	DERatio             DoublePair   `gorm:"embedded;embeddedPrefix:de_ratio_" json:"de_ratio"`
	Per                 DoublePair   `gorm:"embedded;embeddedPrefix:per_" json:"per"`
	Psr                 DoublePair   `gorm:"embedded;embeddedPrefix:psr_" json:"psr"`
	Pbr                 DoublePair   `gorm:"embedded;embeddedPrefix:pbr_" json:"pbr"`
	Pcr                 DoublePair   `gorm:"embedded;embeddedPrefix:pcr_" json:"pcr"`
	Activities          Activities   `gorm:"embedded;embeddedPrefix:activities_" json:"activities"`
	DividendYield       DoublePair   `gorm:"embedded;embeddedPrefix:dividend_yield_" json:"dividend_yield"`
	DividendPayoutRatio DoublePair   `gorm:"embedded;embeddedPrefix:dividend_payout_ratio_" json:"dividend_payout_ratio"`
	Roa                 DoublePair   `gorm:"embedded;embeddedPrefix:roa_" json:"roa"`
	Roe                 DoublePair   `gorm:"embedded;embeddedPrefix:roe_" json:"roe"`
	MarketCap           IntPair      `gorm:"embedded;embeddedPrefix:market_cap_" json:"market_cap"`
	StartDate           time.Time    `time_format:"2006-01-02T15:04:05.000Z" json:"start_date"`
	EndDate             time.Time    `time_format:"2006-01-02T15:04:05.000Z" json:"end_date"`
}

type MainSector struct {
	QuantOptionID uint   `json:"-"`
	Name          string `json:"name"`
}

func NewMainSectors(quantID uint, ms []string) []MainSector {
	var res []MainSector

	for _, val := range ms {
		res = append(res, MainSector{quantID, val})
	}
	return res
}

func (ms *MainSector) ToString() string {
	return ms.Name
}

type IntPair struct {
	Max int64 `json:"max"`
	Min int64 `json:"min"`
}

func (ip *IntPair) ToPB() *proto.IntPair {
	return &proto.IntPair{
		Max: ip.Max,
		Min: ip.Min,
	}
}

type DoublePair struct {
	Max float32 `json:"max"`
	Min float32 `json:"min"`
}

func (dp *DoublePair) ToPB() *proto.DoublePair {
	return &proto.DoublePair{
		Max: dp.Max,
		Min: dp.Min,
	}
}

type Activities struct {
	Operating DoublePair `gorm:"embedded;embeddedPrefix:operating_" json:"operating"`
	Investing DoublePair `gorm:"embedded;embeddedPrefix:investing_" json:"investing"`
	Financing DoublePair `gorm:"embedded;embeddedPrefix:financing_" json:"financing"`
}

func (a *Activities) ToPB() *proto.Activities {
	return &proto.Activities{
		Operating: a.Operating.ToPB(),
		Investing: a.Investing.ToPB(),
		Financing: a.Financing.ToPB(),
	}
}

func (q *QuantOption) ToRequest() *proto.QuantRequest {
	var sectors []string
	for _, ms := range q.MainSectors {
		sectors = append(sectors, ms.ToString())
	}
	return &proto.QuantRequest{
		MainSector:          sectors,
		NetRevenue:          q.NetRevenue.ToPB(),
		NetRevenueRate:      q.NetRevenueRate.ToPB(),
		NetProfit:           q.NetProfit.ToPB(),
		NetProfitRate:       q.NetProfitRate.ToPB(),
		DeRatio:             q.DERatio.ToPB(),
		Per:                 q.Per.ToPB(),
		Psr:                 q.Psr.ToPB(),
		Pbr:                 q.Pbr.ToPB(),
		Pcr:                 q.Pcr.ToPB(),
		Activities:          q.Activities.ToPB(),
		DividendYield:       q.DividendYield.ToPB(),
		DividendPayoutRatio: q.DividendPayoutRatio.ToPB(),
		Roa:                 q.Roa.ToPB(),
		Roe:                 q.Roe.ToPB(),
		MarketCap:           q.MarketCap.ToPB(),
		StartDate:           timestamppb.New(q.StartDate),
		EndDate:             timestamppb.New(q.EndDate),
	}
}

func (q *QuantOption) ToMap() map[string]interface{} {
	return nil
}

func (q *QuantOption) TableName() string {
	return "quant_options"
}
