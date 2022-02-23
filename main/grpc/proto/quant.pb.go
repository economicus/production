// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: quant.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ChartData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartDate      *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start_date,proto3" json:"start_date,omitempty"`
	ProfitRateData []float32              `protobuf:"fixed32,2,rep,packed,name=profit_rate_data,proto3" json:"profit_rate_data,omitempty"`
}

func (x *ChartData) Reset() {
	*x = ChartData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChartData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChartData) ProtoMessage() {}

func (x *ChartData) ProtoReflect() protoreflect.Message {
	mi := &file_quant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChartData.ProtoReflect.Descriptor instead.
func (*ChartData) Descriptor() ([]byte, []int) {
	return file_quant_proto_rawDescGZIP(), []int{0}
}

func (x *ChartData) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *ChartData) GetProfitRateData() []float32 {
	if x != nil {
		return x.ProfitRateData
	}
	return nil
}

// 실험실에서 입력한 옵션값에 대한 데이터 결과
type QuantResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CumulativeReturn    float64    `protobuf:"fixed64,1,opt,name=cumulative_return,proto3" json:"cumulative_return,omitempty"`
	AnnualAverageReturn float64    `protobuf:"fixed64,2,opt,name=annual_average_return,proto3" json:"annual_average_return,omitempty"`
	WinningPercentage   float64    `protobuf:"fixed64,3,opt,name=winning_percentage,proto3" json:"winning_percentage,omitempty"`
	MaxLossRate         float64    `protobuf:"fixed64,4,opt,name=max_loss_rate,proto3" json:"max_loss_rate,omitempty"`
	HoldingsCount       int32      `protobuf:"varint,5,opt,name=holdings_count,proto3" json:"holdings_count,omitempty"`
	ChartData           *ChartData `protobuf:"bytes,6,opt,name=chart_data,proto3" json:"chart_data,omitempty"`
}

func (x *QuantResult) Reset() {
	*x = QuantResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuantResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuantResult) ProtoMessage() {}

func (x *QuantResult) ProtoReflect() protoreflect.Message {
	mi := &file_quant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuantResult.ProtoReflect.Descriptor instead.
func (*QuantResult) Descriptor() ([]byte, []int) {
	return file_quant_proto_rawDescGZIP(), []int{1}
}

func (x *QuantResult) GetCumulativeReturn() float64 {
	if x != nil {
		return x.CumulativeReturn
	}
	return 0
}

func (x *QuantResult) GetAnnualAverageReturn() float64 {
	if x != nil {
		return x.AnnualAverageReturn
	}
	return 0
}

func (x *QuantResult) GetWinningPercentage() float64 {
	if x != nil {
		return x.WinningPercentage
	}
	return 0
}

func (x *QuantResult) GetMaxLossRate() float64 {
	if x != nil {
		return x.MaxLossRate
	}
	return 0
}

func (x *QuantResult) GetHoldingsCount() int32 {
	if x != nil {
		return x.HoldingsCount
	}
	return 0
}

func (x *QuantResult) GetChartData() *ChartData {
	if x != nil {
		return x.ChartData
	}
	return nil
}

// 실험실에서 유저가 입력할 수 있는 퀀트 모델의 옵션
type QuantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 사업 분야
	MainSector []string `protobuf:"bytes,1,rep,name=main_sector,proto3" json:"main_sector,omitempty"`
	// 매출액
	NetRevenue *IntPair `protobuf:"bytes,2,opt,name=net_revenue,proto3" json:"net_revenue,omitempty"`
	// 매출액 증가율
	NetRevenueRate *DoublePair `protobuf:"bytes,3,opt,name=net_revenue_rate,proto3" json:"net_revenue_rate,omitempty"`
	// 당기 순이익
	NetProfit *IntPair `protobuf:"bytes,4,opt,name=net_profit,proto3" json:"net_profit,omitempty"`
	// 당기 순이익 증가율
	NetProfitRate *DoublePair `protobuf:"bytes,5,opt,name=net_profit_rate,proto3" json:"net_profit_rate,omitempty"`
	// 부채율
	DeRatio *DoublePair `protobuf:"bytes,6,opt,name=de_ratio,proto3" json:"de_ratio,omitempty"`
	// PER
	Per *DoublePair `protobuf:"bytes,7,opt,name=per,proto3" json:"per,omitempty"`
	// PSR
	Psr *DoublePair `protobuf:"bytes,8,opt,name=psr,proto3" json:"psr,omitempty"`
	// PBR
	Pbr *DoublePair `protobuf:"bytes,9,opt,name=pbr,proto3" json:"pbr,omitempty"`
	// PCR
	Pcr *DoublePair `protobuf:"bytes,10,opt,name=pcr,proto3" json:"pcr,omitempty"`
	// 현금 흐름
	Activities *Activities `protobuf:"bytes,11,opt,name=activities,proto3" json:"activities,omitempty"`
	// 현금 배당 수익률
	DividendYield *DoublePair `protobuf:"bytes,12,opt,name=dividend_yield,proto3" json:"dividend_yield,omitempty"`
	// 현금 배당 성향
	DividendPayoutRatio *DoublePair `protobuf:"bytes,13,opt,name=dividend_payout_ratio,proto3" json:"dividend_payout_ratio,omitempty"`
	// ROA
	Roa *DoublePair `protobuf:"bytes,14,opt,name=roa,proto3" json:"roa,omitempty"`
	// ROE
	Roe *DoublePair `protobuf:"bytes,15,opt,name=roe,proto3" json:"roe,omitempty"`
	// 시가 총액
	MarketCap *IntPair `protobuf:"bytes,16,opt,name=market_cap,proto3" json:"market_cap,omitempty"`
	// 검색 시작 날짜
	StartDate *timestamppb.Timestamp `protobuf:"bytes,17,opt,name=start_date,proto3" json:"start_date,omitempty"`
	// 검색 끝 날짜
	EndDate *timestamppb.Timestamp `protobuf:"bytes,18,opt,name=end_date,proto3" json:"end_date,omitempty"`
}

func (x *QuantRequest) Reset() {
	*x = QuantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quant_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuantRequest) ProtoMessage() {}

func (x *QuantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_quant_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuantRequest.ProtoReflect.Descriptor instead.
func (*QuantRequest) Descriptor() ([]byte, []int) {
	return file_quant_proto_rawDescGZIP(), []int{2}
}

func (x *QuantRequest) GetMainSector() []string {
	if x != nil {
		return x.MainSector
	}
	return nil
}

func (x *QuantRequest) GetNetRevenue() *IntPair {
	if x != nil {
		return x.NetRevenue
	}
	return nil
}

func (x *QuantRequest) GetNetRevenueRate() *DoublePair {
	if x != nil {
		return x.NetRevenueRate
	}
	return nil
}

func (x *QuantRequest) GetNetProfit() *IntPair {
	if x != nil {
		return x.NetProfit
	}
	return nil
}

func (x *QuantRequest) GetNetProfitRate() *DoublePair {
	if x != nil {
		return x.NetProfitRate
	}
	return nil
}

func (x *QuantRequest) GetDeRatio() *DoublePair {
	if x != nil {
		return x.DeRatio
	}
	return nil
}

func (x *QuantRequest) GetPer() *DoublePair {
	if x != nil {
		return x.Per
	}
	return nil
}

func (x *QuantRequest) GetPsr() *DoublePair {
	if x != nil {
		return x.Psr
	}
	return nil
}

func (x *QuantRequest) GetPbr() *DoublePair {
	if x != nil {
		return x.Pbr
	}
	return nil
}

func (x *QuantRequest) GetPcr() *DoublePair {
	if x != nil {
		return x.Pcr
	}
	return nil
}

func (x *QuantRequest) GetActivities() *Activities {
	if x != nil {
		return x.Activities
	}
	return nil
}

func (x *QuantRequest) GetDividendYield() *DoublePair {
	if x != nil {
		return x.DividendYield
	}
	return nil
}

func (x *QuantRequest) GetDividendPayoutRatio() *DoublePair {
	if x != nil {
		return x.DividendPayoutRatio
	}
	return nil
}

func (x *QuantRequest) GetRoa() *DoublePair {
	if x != nil {
		return x.Roa
	}
	return nil
}

func (x *QuantRequest) GetRoe() *DoublePair {
	if x != nil {
		return x.Roe
	}
	return nil
}

func (x *QuantRequest) GetMarketCap() *IntPair {
	if x != nil {
		return x.MarketCap
	}
	return nil
}

func (x *QuantRequest) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *QuantRequest) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

type Activities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 영업 현금 흐름
	Operating *DoublePair `protobuf:"bytes,1,opt,name=operating,proto3" json:"operating,omitempty"`
	// 투자 현금 흐름
	Investing *DoublePair `protobuf:"bytes,2,opt,name=investing,proto3" json:"investing,omitempty"`
	// 재무 현금 흐름
	Financing *DoublePair `protobuf:"bytes,3,opt,name=financing,proto3" json:"financing,omitempty"`
}

func (x *Activities) Reset() {
	*x = Activities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quant_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activities) ProtoMessage() {}

func (x *Activities) ProtoReflect() protoreflect.Message {
	mi := &file_quant_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activities.ProtoReflect.Descriptor instead.
func (*Activities) Descriptor() ([]byte, []int) {
	return file_quant_proto_rawDescGZIP(), []int{3}
}

func (x *Activities) GetOperating() *DoublePair {
	if x != nil {
		return x.Operating
	}
	return nil
}

func (x *Activities) GetInvesting() *DoublePair {
	if x != nil {
		return x.Investing
	}
	return nil
}

func (x *Activities) GetFinancing() *DoublePair {
	if x != nil {
		return x.Financing
	}
	return nil
}

type IntPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 최솟값
	Min int64 `protobuf:"varint,1,opt,name=min,proto3" json:"min,omitempty"`
	// 최댓값
	Max int64 `protobuf:"varint,2,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *IntPair) Reset() {
	*x = IntPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quant_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntPair) ProtoMessage() {}

func (x *IntPair) ProtoReflect() protoreflect.Message {
	mi := &file_quant_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntPair.ProtoReflect.Descriptor instead.
func (*IntPair) Descriptor() ([]byte, []int) {
	return file_quant_proto_rawDescGZIP(), []int{4}
}

func (x *IntPair) GetMin() int64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *IntPair) GetMax() int64 {
	if x != nil {
		return x.Max
	}
	return 0
}

type DoublePair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 최솟값
	Min float32 `protobuf:"fixed32,1,opt,name=min,proto3" json:"min,omitempty"`
	// 최댓값
	Max float32 `protobuf:"fixed32,2,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *DoublePair) Reset() {
	*x = DoublePair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quant_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoublePair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoublePair) ProtoMessage() {}

func (x *DoublePair) ProtoReflect() protoreflect.Message {
	mi := &file_quant_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoublePair.ProtoReflect.Descriptor instead.
func (*DoublePair) Descriptor() ([]byte, []int) {
	return file_quant_proto_rawDescGZIP(), []int{5}
}

func (x *DoublePair) GetMin() float32 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *DoublePair) GetMax() float32 {
	if x != nil {
		return x.Max
	}
	return 0
}

var File_quant_proto protoreflect.FileDescriptor

var file_quant_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x72, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x3a, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2a,
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x02, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa1, 0x02, 0x0a, 0x0b, 0x51,
	0x75, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x75,
	0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x12, 0x34, 0x0a, 0x15, 0x61, 0x6e, 0x6e, 0x75,
	0x61, 0x6c, 0x5f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x15, 0x61, 0x6e, 0x6e, 0x75, 0x61, 0x6c, 0x5f,
	0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x12, 0x2e,
	0x0a, 0x12, 0x77, 0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x12, 0x77, 0x69, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x24,
	0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x6f, 0x73, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x6f, 0x73, 0x73, 0x5f,
	0x72, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x68, 0x6f, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x68, 0x6f,
	0x6c, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x0a,
	0x63, 0x68, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0xf6,
	0x06, 0x0a, 0x0c, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x12, 0x30, 0x0a, 0x0b, 0x6e, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49,
	0x6e, 0x74, 0x50, 0x61, 0x69, 0x72, 0x52, 0x0b, 0x6e, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x76, 0x65,
	0x6e, 0x75, 0x65, 0x12, 0x3d, 0x0a, 0x10, 0x6e, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x76, 0x65, 0x6e,
	0x75, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x69, 0x72,
	0x52, 0x10, 0x6e, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x5f, 0x72, 0x61,
	0x74, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49,
	0x6e, 0x74, 0x50, 0x61, 0x69, 0x72, 0x52, 0x0a, 0x6e, 0x65, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x74, 0x12, 0x3b, 0x0a, 0x0f, 0x6e, 0x65, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x0f,
	0x6e, 0x65, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x12,
	0x2d, 0x0a, 0x08, 0x64, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x50, 0x61, 0x69, 0x72, 0x52, 0x08, 0x64, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x23,
	0x0a, 0x03, 0x70, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x03,
	0x70, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x03, 0x70, 0x73, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x50,
	0x61, 0x69, 0x72, 0x52, 0x03, 0x70, 0x73, 0x72, 0x12, 0x23, 0x0a, 0x03, 0x70, 0x62, 0x72, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x03, 0x70, 0x62, 0x72, 0x12, 0x23, 0x0a,
	0x03, 0x70, 0x63, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x03, 0x70,
	0x63, 0x72, 0x12, 0x31, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e,
	0x64, 0x5f, 0x79, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x69, 0x72,
	0x52, 0x0e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x5f, 0x79, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x47, 0x0a, 0x15, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x5f, 0x70, 0x61, 0x79,
	0x6f, 0x75, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x50, 0x61,
	0x69, 0x72, 0x52, 0x15, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x64, 0x5f, 0x70, 0x61, 0x79,
	0x6f, 0x75, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x23, 0x0a, 0x03, 0x72, 0x6f, 0x61,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x03, 0x72, 0x6f, 0x61, 0x12, 0x23,
	0x0a, 0x03, 0x72, 0x6f, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x03,
	0x72, 0x6f, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x63, 0x61,
	0x70, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x49, 0x6e, 0x74, 0x50, 0x61, 0x69, 0x72, 0x52, 0x0a, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f,
	0x63, 0x61, 0x70, 0x12, 0x3a, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x36, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x65,
	0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x22, 0x9f, 0x01, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x09, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x2f, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x09, 0x69,
	0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x2f, 0x0a, 0x09, 0x66, 0x69, 0x6e, 0x61,
	0x6e, 0x63, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x09,
	0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x22, 0x2d, 0x0a, 0x07, 0x49, 0x6e, 0x74,
	0x50, 0x61, 0x69, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x22, 0x30, 0x0a, 0x0a, 0x44, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x50, 0x61, 0x69, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x32, 0x3d, 0x0a, 0x05, 0x51, 0x75,
	0x61, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x61, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_quant_proto_rawDescOnce sync.Once
	file_quant_proto_rawDescData = file_quant_proto_rawDesc
)

func file_quant_proto_rawDescGZIP() []byte {
	file_quant_proto_rawDescOnce.Do(func() {
		file_quant_proto_rawDescData = protoimpl.X.CompressGZIP(file_quant_proto_rawDescData)
	})
	return file_quant_proto_rawDescData
}

var file_quant_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_quant_proto_goTypes = []interface{}{
	(*ChartData)(nil),             // 0: proto.ChartData
	(*QuantResult)(nil),           // 1: proto.QuantResult
	(*QuantRequest)(nil),          // 2: proto.QuantRequest
	(*Activities)(nil),            // 3: proto.Activities
	(*IntPair)(nil),               // 4: proto.IntPair
	(*DoublePair)(nil),            // 5: proto.DoublePair
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_quant_proto_depIdxs = []int32{
	6,  // 0: proto.ChartData.start_date:type_name -> google.protobuf.Timestamp
	0,  // 1: proto.QuantResult.chart_data:type_name -> proto.ChartData
	4,  // 2: proto.QuantRequest.net_revenue:type_name -> proto.IntPair
	5,  // 3: proto.QuantRequest.net_revenue_rate:type_name -> proto.DoublePair
	4,  // 4: proto.QuantRequest.net_profit:type_name -> proto.IntPair
	5,  // 5: proto.QuantRequest.net_profit_rate:type_name -> proto.DoublePair
	5,  // 6: proto.QuantRequest.de_ratio:type_name -> proto.DoublePair
	5,  // 7: proto.QuantRequest.per:type_name -> proto.DoublePair
	5,  // 8: proto.QuantRequest.psr:type_name -> proto.DoublePair
	5,  // 9: proto.QuantRequest.pbr:type_name -> proto.DoublePair
	5,  // 10: proto.QuantRequest.pcr:type_name -> proto.DoublePair
	3,  // 11: proto.QuantRequest.activities:type_name -> proto.Activities
	5,  // 12: proto.QuantRequest.dividend_yield:type_name -> proto.DoublePair
	5,  // 13: proto.QuantRequest.dividend_payout_ratio:type_name -> proto.DoublePair
	5,  // 14: proto.QuantRequest.roa:type_name -> proto.DoublePair
	5,  // 15: proto.QuantRequest.roe:type_name -> proto.DoublePair
	4,  // 16: proto.QuantRequest.market_cap:type_name -> proto.IntPair
	6,  // 17: proto.QuantRequest.start_date:type_name -> google.protobuf.Timestamp
	6,  // 18: proto.QuantRequest.end_date:type_name -> google.protobuf.Timestamp
	5,  // 19: proto.Activities.operating:type_name -> proto.DoublePair
	5,  // 20: proto.Activities.investing:type_name -> proto.DoublePair
	5,  // 21: proto.Activities.financing:type_name -> proto.DoublePair
	2,  // 22: proto.Quant.Request:input_type -> proto.QuantRequest
	1,  // 23: proto.Quant.Request:output_type -> proto.QuantResult
	23, // [23:24] is the sub-list for method output_type
	22, // [22:23] is the sub-list for method input_type
	22, // [22:22] is the sub-list for extension type_name
	22, // [22:22] is the sub-list for extension extendee
	0,  // [0:22] is the sub-list for field type_name
}

func init() { file_quant_proto_init() }
func file_quant_proto_init() {
	if File_quant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_quant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChartData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_quant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuantResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_quant_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuantRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_quant_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Activities); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_quant_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntPair); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_quant_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoublePair); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_quant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_quant_proto_goTypes,
		DependencyIndexes: file_quant_proto_depIdxs,
		MessageInfos:      file_quant_proto_msgTypes,
	}.Build()
	File_quant_proto = out.File
	file_quant_proto_rawDesc = nil
	file_quant_proto_goTypes = nil
	file_quant_proto_depIdxs = nil
}
