package quant

import "main/internal/core/pb"

type IntPair struct {
	Max int64 `json:"max"`
	Min int64 `json:"min"`
}

func (ip *IntPair) ToPB() *pb.IntPair {
	return &pb.IntPair{
		Max: ip.Max,
		Min: ip.Min,
	}
}

type DoublePair struct {
	Max float32 `json:"max"`
	Min float32 `json:"min"`
}

func (dp *DoublePair) ToPB() *pb.DoublePair {
	return &pb.DoublePair{
		Max: dp.Max,
		Min: dp.Min,
	}
}
