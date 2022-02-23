package repository

import (
	"economicus/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuantRepository_GetAllQuantsWithEmptyOption(t *testing.T) {
	a := assert.New(t)
	o := &models.QueryOption{}
	quants, err := quantRepo.GetAllQuants(uint(1), o)
	a.NoError(err, err)
	a.NotEmpty(quants, "quants should not be empty")
}

func TestQuantRepository_GetAQuant(t *testing.T) {
	a := assert.New(t)
	quant, err := quantRepo.GetQuant(uint(1))
	a.NoError(err, err)
	a.Equal(quant.Name, "quant_ex00")
}

func TestQuantRepository_GetMyQuants(t *testing.T) {
	a := assert.New(t)
	quants, err := quantRepo.GetMyQuants(uint(1))
	a.NoError(err, err)
	a.NotEmpty(quants, "quants of user 1 should not be empty")
}

func TestQuantRepository_CreateQuant_With_ValidData(t *testing.T) {
	a := assert.New(t)
	quant := models.NewQuant(1, "Quant example")
	_, err := quantRepo.CreateQuant(quant)
	a.NoError(err, err)
}

func TestQuantRepository_CreateQuant_With_InvalidData_Duplicated_Name(t *testing.T) {
	a := assert.New(t)
	quant := models.NewQuant(1, "Quant example")
	_, err := quantRepo.CreateQuant(quant)
	a.Error(err, "error should be generated")
}

func TestQuantRepository_UpdateQuant_With_ValidData(t *testing.T) {
	a := assert.New(t)
	data := map[string]interface{}{
		"name":        "edited_name",
		"profit_rate": 0.0,
		"description": "edited name",
	}
	err := quantRepo.UpdateQuant(uint(1), data)
	a.NoError(err, err)
}

func TestQuantRepository_UpdateQuant_With_InvalidData_Duplicated_Name(t *testing.T) {
	a := assert.New(t)
	data := map[string]interface{}{
		"name":        "quant_example",
		"profit_rate": 0.0,
		"description": "edited name",
	}
	err := quantRepo.UpdateQuant(uint(1), data)
	a.Error(err, err)
}

func TestQuantRepository_DeleteQuant(t *testing.T) {
	a := assert.New(t)
	err := quantRepo.DeleteQuant(uint(5))
	a.NoError(err, err)
}
