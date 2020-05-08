package main

import (
	"github.com/RomanAvdeenko/domain1"
)

func main() {
	var (
		wm1   = &domain1.WashingMachine{BrandName: "Zanussi", ModelName: "zws 1040", SerialNumber: "8901752304958204532"}
		m1    = &domain1.Man{Name: "Roman Avdeenko"}
		wm1CP domain1.WMCP
	)

	wm1CP = wm1

	m1.TurnOnWashingMachine(wm1CP)
	m1.TurnOffWashingMachine(wm1CP)

}
