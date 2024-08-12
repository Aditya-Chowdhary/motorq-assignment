package vehicles

import (
	"testing"
)

func TestCallNHSTA(t *testing.T) {
	res, err := callNHSTA("HFISYFHRJS847JDY7")
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", res)
}
