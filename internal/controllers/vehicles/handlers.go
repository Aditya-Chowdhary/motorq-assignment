package vehicles

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

type VehicleHandler struct {
	db *pgxpool.Pool
}

func Handler(db *pgxpool.Pool) *VehicleHandler {
	return &VehicleHandler{
		db: db,
	}
}

func (v *VehicleHandler) DecodeVehicle(c *gin.Context) {
	var input struct {
		VIN string `uri:"vin" binding:"required,alphanum,len=17"`
	}

	err := c.BindUri(&input)
	if err != nil {
		merrors.BadRequest(c, err.Error())
		return
	}

	result, err := callNHSTA(input.VIN)
	if errors.Is(err, ErrNoCar) {
		merrors.NotFound(c, ErrNoCar.Error())
		return
	} else if err != nil {
		merrors.InternalServer(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, utils.BaseResponse{
		Success:    true,
		Message:    "Details retrieved successfully",
		Data:       result,
		StatusCode: http.StatusOK,
	})
}

func (v *VehicleHandler) GetVehicle(c *gin.Context) {
	var input struct {
		VIN string `uri:"vin" binding:"required,alphanum,len=17"`
	}

	err := c.BindUri(&input)
	if err != nil {
		merrors.BadRequest(c, err.Error())
		return
	}

	q := database.New(v.db)
	vehicle, err := q.GetVehicle(c, input.VIN)
	if errors.Is(err, pgx.ErrNoRows) {
		merrors.NotFound(c, ErrNoCar.Error())
		return
	} else if err != nil {
		merrors.InternalServer(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, utils.BaseResponse{
		Success:    true,
		Message:    "Details retrieved successfully",
		Data:       vehicle,
		StatusCode: http.StatusOK,
	})
}

func (v *VehicleHandler) CreateVehicle(c *gin.Context) {
	var input struct {
		VIN   string `json:"vin" binding:"required,alphanum,len=17"`
		OrgID int64  `json:"org" binding:"required"`
	}

	result, err := callNHSTA(input.VIN)
	if errors.Is(err, ErrNoCar) {
		merrors.NotFound(c, ErrNoCar.Error())
		return
	} else if err != nil {
		merrors.InternalServer(c, err.Error())
		return
	}

	q := database.New(v.db)
	vehicle, err := q.CreateVehicle(c, database.CreateVehicleParams{
		Vin:          input.VIN,
		OrgID:        input.OrgID,
		Manufacturer: result.Manufacturer,
		Make:         result.Make,
		Model:        result.Model,
		Year:         int32(result.ModelYear),
	})
	if err != nil {
		merrors.InternalServer(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, utils.BaseResponse{
		Success:    true,
		Message:    "Vehicle created successfully",
		Data:       vehicle,
		StatusCode: http.StatusOK,
	})
}
