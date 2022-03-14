package util

import "testing"

func TestStringToTime(t *testing.T) {
	var convertedTime = StringToTime("2022-03-14T09:21:30")

	if convertedTime.Year() != 2022 {
		t.Errorf("Espera que o ano seja 2022")
	}
	if convertedTime.Month() != 03 {
		t.Errorf("Espera que o mês seja 03")
	}
	if convertedTime.Hour() != 9 {
		t.Errorf("Espera que a hora seja 9")
	}
}
