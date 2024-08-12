package organisations

import (
	"errors"
	"motorq-assignment/internal/database"
	"motorq-assignment/internal/merrors"
	"motorq-assignment/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrgHandler struct {
	db *pgxpool.Pool
}

func Handler(db *pgxpool.Pool) *OrgHandler {
	return &OrgHandler{
		db: db,
	}
}

func (o *OrgHandler) GetOrganisation(c *gin.Context) {
	q := database.New(o.db)
	orgs, err := q.GetAllOrganisations(c)
	if errors.Is(err, pgx.ErrNoRows) {
		merrors.NotFound(c, "No Organisations found!")
		return
	} else if err != nil {
		merrors.InternalServer(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, utils.BaseResponse{
		Success:    true,
		Message:    "Organisations successfully retrieved",
		Data:       orgs,
		StatusCode: http.StatusOK,
	})
}

func (o *OrgHandler) CreateOrgansation(c *gin.Context) {
	var input struct {
		Name        string `json:"name" binding:"required"`
		Account     string `json:"account" binding:"required"`
		Website     string `json:"website" binding:"required"`
		FuelPolicy  int32  `json:"fuel_policy" binding:"gt=-1"`
		SpeedPolicy int32  `json:"speed_policy" binding:"required,gt=-1"`
		ParentID    *int64 `json:"parent_id"`
	}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		merrors.BadRequest(c, err.Error())
		return
	}

	q := database.New(o.db)
	org, err := q.CreateOrganisation(c, database.CreateOrganisationParams{
		Name:        input.Name,
		Account:     input.Account,
		Website:     input.Website,
		FuelPolicy:  input.FuelPolicy,
		SpeedPolicy: input.SpeedPolicy,
		ParentID:    input.ParentID,
	})
	if err != nil {
		merrors.InternalServer(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, utils.BaseResponse{
		Success:    true,
		Message:    "Organisation successfully added",
		Data:       org,
		StatusCode: http.StatusOK,
	})
}

func (o *OrgHandler) UpdateOrganisation(c *gin.Context) {
	var input struct {
		OrgID       int64  `json:"org_id" binding:"required"`
		Account     string `json:"account" binding:"required"`
		Website     string `json:"website" binding:"required"`
		FuelPolicy  int32  `json:"fuel_policy" binding:"gt=-1"`
		SpeedPolicy int32  `json:"speed_policy" binding:"required,gt=-1"`
	}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		merrors.BadRequest(c, err.Error())
		return
	}

	q := database.New(o.db)
	org, err := q.UpdateOrganisation(c, database.UpdateOrganisationParams{
		OrgID:       input.OrgID,
		Account:     input.Account,
		Website:     input.Website,
		FuelPolicy:  input.FuelPolicy,
		SpeedPolicy: input.SpeedPolicy,
	})
	if errors.Is(err, pgx.ErrNoRows) {
		merrors.NotFound(c, "No such organisation exists")
		return
	} else if err != nil {
		merrors.BadRequest(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, utils.BaseResponse{
		Success:    true,
		Message:    "Organisation successfully updated",
		Data:       org,
		StatusCode: http.StatusOK,
	})
}
