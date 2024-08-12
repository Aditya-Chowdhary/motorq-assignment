package vehicles

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type NHSTADetails struct {
	Make         string `json:"Make"`
	Manufacturer string `json:"Manufacturer"`
	Model        string `json:"Model"`
	ModelYear    int    `json:"ModelYear"`
}

type NHSTAResponse struct {
	Results []struct {
		AdditionalErrorText string `json:"AdditionalErrorText"`
		Make                string `json:"Make"`
		Manufacturer        string `json:"Manufacturer"`
		Model               string `json:"Model"`
		ModelYear           string `json:"ModelYear"`
	} `json:"Results"`
}

var ErrNoCar = errors.New("no car associated with the given VIN")

func (v *VehicleHandler) Do(url string) (*http.Response, error) {
	ctx := context.Background()
	err := v.RateLimiter.Wait(ctx)
	if err != nil {
		return nil, err
	}
	res, err := v.Client.Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (v *VehicleHandler) callNHSTA(vin string) (*NHSTADetails, error) {
	url := fmt.Sprintf("https://vpic.nhtsa.dot.gov/api/vehicles/DecodeVinValues/%s?format=json", vin)

	r, err := v.Do(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	var response NHSTAResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	if response.Results[0].AdditionalErrorText != "" {
		return nil, ErrNoCar
	}
	year, err := strconv.Atoi(response.Results[0].ModelYear)
	if err != nil {
		return nil, err
	}

	details := NHSTADetails{
		Make:         response.Results[0].Make,
		Manufacturer: response.Results[0].Manufacturer,
		Model:        response.Results[0].Model,
		ModelYear:    year,
	}
	return &details, nil
}
