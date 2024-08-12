package vehicles

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func TestCallNHSTA(t *testing.T) {
	v := &VehicleHandler{
		Client:      &http.Client{Timeout: 30 * time.Second},
		RateLimiter: rate.NewLimiter(rate.Every(time.Minute/5), 5),
	}
	res, err := v.callNHSTA("5UXWX7C5*BA")
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", res)
}

func TestRateCall(t *testing.T) {
	v := &VehicleHandler{
		Client:      &http.Client{Timeout: 30 * time.Second},
		RateLimiter: rate.NewLimiter(rate.Every(time.Minute/5), 5),
	}
	vin := "5UXWX7C5*BA"
	url := fmt.Sprintf("https://vpic.nhtsa.dot.gov/api/vehicles/DecodeVinValues/%s?format=json", vin)
	for i := 0; i < 10; i++ {
		t.Log(i)
		res, err := v.Do(url)
		if err != nil {
			t.Log(res.StatusCode)
			t.Error(err.Error())
		}
		if res.StatusCode == 429 {
			t.Logf("Rate limit reached after %d requests", i)
		}
	}
}
