// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

type Organisation struct {
	OrgID       int64  `json:"org_id"`
	Name        string `json:"name"`
	Account     string `json:"account"`
	Website     string `json:"website"`
	FuelPolicy  int32  `json:"fuel_policy"`
	FuelSetBy   int64  `json:"fuel_set_by"`
	SpeedPolicy int32  `json:"speed_policy"`
	SpeedSetBy  int64  `json:"speed_set_by"`
	ParentID    *int64 `json:"parent_id"`
}

type Vehicle struct {
	VehicleID    int64  `json:"vehicle_id"`
	Vin          string `json:"vin"`
	OrgID        int64  `json:"org_id"`
	Manufacturer string `json:"manufacturer"`
	Make         string `json:"make"`
	Model        string `json:"model"`
	Year         int32  `json:"year"`
}
