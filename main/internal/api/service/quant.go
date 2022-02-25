package service

import (
	"context"
	"economicus/internal/api/repository"
	ecoerror "economicus/internal/error"
	"economicus/internal/models"
	"economicus/proto"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type QuantService struct {
	repo repository.QuantRepositoryFactory
}

func NewQuantService(repo repository.QuantRepositoryFactory) *QuantService {
	return &QuantService{
		repo: repo,
	}
}

func (s *QuantService) GetAllQuants(userID uint, option *models.QueryOption) (models.Quants, error) {
	return s.repo.GetAllQuants(userID, option)
}

func (s *QuantService) GetFollowingsQuants(userID uint, option *models.QueryOption) (models.Quants, error) {
	return s.repo.GetFollowingsQuants(userID, option)
}

func (s *QuantService) GetQuant(quantID uint) (*models.Quant, error) {
	return s.repo.GetQuant(quantID)
}

func (s *QuantService) GetMyQuants(userID uint) (models.Quants, error) {
	return s.repo.GetMyQuants(userID)
}

func (s *QuantService) CreateQuant(userID uint, request *models.QuantRequest) (*models.QuantResult, error) {
	if err := s.repo.CheckModelName(request.Name); err != nil {
		return nil, err
	}

	quant := models.NewQuant(userID, request.Name)
	quantID, err := s.repo.CreateQuant(quant)
	if err != nil {
		return nil, err
	}

	request.QuantID = quantID
	option := request.ToQuantOption()

	if err = s.repo.CreateQuantOption(option); err != nil {
		return nil, err
	}

	protoData := option.ToRequest()
	conn, err := grpc.Dial("172.17.0.1:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := proto.NewQuantClient(conn)
	ctx, f := context.WithTimeout(context.Background(), time.Minute*3)
	defer f()

	res, err := client.Request(ctx, protoData)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	quantResult := models.NewQuantResultFromProto(res)
	if err = quantResult.AddKospiData(); err != nil {
		return nil, err
	}

	return quantResult, nil
}

func (s *QuantService) UpdateQuant(userID uint, request map[string]interface{}) error {
	quantID, exist := request["quant_id"]
	if !exist {
		return ecoerror.ErrInvalidJson
	}

	delete(request, "quant_id")
	if err := s.repo.CheckQuantPermission(userID, uint(quantID.(float64))); err != nil {
		return err
	}

	request["updated_at"] = time.Now()

	return s.repo.UpdateQuant(uint(quantID.(float64)), request)
}

func (s *QuantService) UpdateQuantOption(userID uint, option *models.QuantOption) error {
	if err := s.repo.CheckQuantPermission(userID, option.QuantID); err != nil {
		return err
	}

	return s.repo.UpdateQuantOption(option)
}

func (s *QuantService) DeleteQuant(userID, quantID uint) error {
	if err := s.repo.CheckQuantPermission(userID, quantID); err != nil {
		return err
	}

	return s.repo.DeleteQuant(quantID)
}
