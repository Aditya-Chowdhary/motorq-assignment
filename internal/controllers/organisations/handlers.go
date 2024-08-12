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
	var org database.Organisation
	if input.ParentID == nil {
		org, err = q.CreateOrganisation(c, database.CreateOrganisationParams{
			Name:        input.Name,
			Account:     input.Account,
			Website:     input.Website,
			FuelPolicy:  input.FuelPolicy,
			SpeedPolicy: input.SpeedPolicy,
			ParentID:    input.ParentID,
		})
	} else {
		org, err = q.CreateOrganisationThroughParent(c, database.CreateOrganisationThroughParentParams{
			Name:        input.Name,
			Account:     input.Account,
			Website:     input.Website,
			FuelPolicy:  input.FuelPolicy,
			SpeedPolicy: input.SpeedPolicy,
			ParentID:    *input.ParentID,
		})
	}

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
		OrgID       int64   `json:"org_id" binding:"required"`
		Account     *string `json:"account"`
		Website     *string `json:"website"`
		FuelPolicy  *int32  `json:"fuel_policy" binding:"gt=-1"`
		SpeedPolicy *int32  `json:"speed_policy" binding:"gt=-1"`
	}
	var fuelPolicy int32

	err := c.ShouldBindJSON(&input)
	if err != nil {
		merrors.BadRequest(c, err.Error())
		return
	}

	q := database.New(o.db)
	org, err := q.GetOrganisation(c, input.OrgID)
	if errors.Is(err, pgx.ErrNoRows) {
		merrors.NotFound(c, "No such organisation exists")
		return
	} else if err != nil {
		merrors.InternalServer(c, err.Error())
		return
	}
	if nil == input.Account {
		input.Account = &org.Account
	}
	if nil == input.Website {
		input.Website = &org.Website
	}
	if nil == input.FuelPolicy {
		fuelPolicy = org.FuelPolicy
	} else if org.OrgID != org.FuelSetBy {
		merrors.BadRequest(c, "you cannot update your fuel policy (controlled by parent)")
		return
	}
	if nil == input.SpeedPolicy {
		input.SpeedPolicy = &org.SpeedPolicy
	}

	tx, err := o.db.Begin(c)
	if err != nil {
		merrors.InternalServer(c, err.Error())
		return
	}
	defer tx.Rollback(c)
	qtx := q.WithTx(tx)

	var row int64
	org, err = qtx.UpdateOrganisation(c, database.UpdateOrganisationParams{
		OrgID:       input.OrgID,
		Account:     *input.Account,
		Website:     *input.Website,
		FuelPolicy:  fuelPolicy,
		SpeedPolicy: *input.SpeedPolicy,
	})

	if input.FuelPolicy != nil && *input.FuelPolicy != fuelPolicy {
		row, err = qtx.UpdateFuelPolicy(c, database.UpdateFuelPolicyParams{
			FuelPolicy: fuelPolicy,
		})
	}

	if err != nil {
		merrors.BadRequest(c, err.Error())
		return
	}

	err = tx.Commit(c)
	if err != nil {
		merrors.InternalServer(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, utils.BaseResponse{
		Success:    true,
		Message:    "Organisation successfully updated",
		Data:       gin.H{"org": org, "rows_updated": row},
		StatusCode: http.StatusOK,
	})
}
