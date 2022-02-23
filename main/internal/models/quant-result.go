package models

import (
	"economicus/grpc/proto"
	"fmt"
	"time"
)

type ChartData struct {
	StartDate       time.Time `time_format:"2006-01-02T15:04:05.000Z" json:"start_date"`
	ProfitRateData  []float32 `json:"profit_rate_data"`
	ProfitKospiData []float32 `json:"profit_kospi_data"`
}

type QuantResult struct {
	QuantID             uint      `json:"-"`
	CumulativeReturn    float64   `json:"cumulative_return"`
	AnnualAverageReturn float64   `json:"annual_average_return"`
	WinningPercentage   float64   `json:"winning_percentage"`
	MaxLossRate         float64   `json:"max_loss_rate"`
	HoldingsCount       int32     `json:"holdings_count"`
	ChartData           ChartData `json:"chart_data"`
}

func NewQuantResultFromProto(pb *proto.QuantResult) *QuantResult {
	return &QuantResult{
		CumulativeReturn:    pb.CumulativeReturn,
		AnnualAverageReturn: pb.AnnualAverageReturn,
		WinningPercentage:   pb.WinningPercentage,
		MaxLossRate:         pb.MaxLossRate,
		HoldingsCount:       pb.HoldingsCount,
		ChartData: ChartData{
			StartDate:      pb.ChartData.StartDate.AsTime(),
			ProfitRateData: pb.ChartData.ProfitRateData,
		},
	}
}

func (qr *QuantResult) AddKospiData() error {
	var kospiVal []float32

	dataLen := len(qr.ChartData.ProfitRateData)
	startTime := qr.ChartData.StartDate
	idx := kospiData.Date[startTime]
	if idx < dataLen {
		return fmt.Errorf("kospi data doesn't match")
	}

	for i := idx; i > idx-dataLen; i-- {
		kospiVal = append(kospiVal, kospiData.IndexVal[i])
	}
	qr.ChartData.ProfitKospiData = kospiVal
	return nil
}
