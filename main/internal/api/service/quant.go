package service

import (
	"main/internal/api/repo"
	g "main/internal/conf/grpc"
	"main/internal/core/model"
	"main/internal/core/model/request"
	"main/internal/core/model/response"
	"time"
)

type QuantService struct {
	repo *repo.QuantRepo
	grpc *g.Quant
}

func NewQuantService(repo *repo.QuantRepo) *QuantService {
	return &QuantService{
		repo: repo,
		grpc: g.New(),
	}
}

func (s *QuantService) GetAllQuants(userID uint, option *model.Query) (model.Quants, error) {
	return s.repo.GetAllQuants(userID, option)
}

func (s *QuantService) GetQuant(quantID uint) (*model.Quant, error) {
	return s.repo.GetQuant(quantID)
}

func (s *QuantService) GetMyQuants(userID uint) (model.Quants, error) {
	return s.repo.GetMyQuants(userID)
}

func (s *QuantService) CreateQuant(userID uint, req *request.QuantC) (*response.QuantResponse, error) {
	if err := s.repo.CheckModelName(req.Name); err != nil {
		return nil, err
	}

	quant := model.NewQuant(userID, req.Name)
	quantID, err := s.repo.CreateQuant(quant)
	if err != nil {
		return nil, err
	}

	req.QuantID = quantID
	option := model.NewQuantOption(req)

	if err = s.repo.CreateQuantOption(option); err != nil {
		return nil, err
	}

	return s.getQuantResponse(option)
}

func (s *QuantService) getQuantResponse(req *model.QuantOption) (*response.QuantResponse, error) {
	gReq := req.ToRequest()
	result, err := s.grpc.Request(gReq)
	if err != nil {
		return nil, err
	}

	resp := response.NewQuantResultFromPB(result)
	if err = resp.AddKospiData(); err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *QuantService) UpdateQuant(userID uint, req *request.QuantE) error {
	q, err := s.GetQuant(req.QuantID)
	if err != nil {
		return err
	}

	if err = repo.CheckPermission(userID, q); err != nil {
		return err
	}

	reqBody := model.ToMap(req)
	reqBody["updated_at"] = time.Now()

	return s.repo.UpdateQuant(q.ID, reqBody)
}

func (s *QuantService) UpdateQuantOption(userID uint, req *model.QuantOption) error {
	q, err := s.GetQuant(req.QuantID)
	if err != nil {
		return err
	}

	if err = repo.CheckPermission(userID, q); err != nil {
		return err
	}

	return s.repo.UpdateQuantOption(q.ID, req.ToMap())
}

func (s *QuantService) DeleteQuant(userID, quantID uint) error {
	q, err := s.GetQuant(quantID)
	if err != nil {
		return err
	}

	if err = repo.CheckPermission(userID, q); err != nil {
		return err
	}

	return s.repo.DeleteQuant(quantID)
}
