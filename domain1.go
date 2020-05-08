package domain1

import "log"

//
type Man struct {
	Name string
}

func (m *Man) TurnOnWashingMachine() {}

//
type WashingMachine struct {
	BrandName    string
	ModelName    string
	SerialNumber string
	Capacity     uint
}

func (m *WashingMachine) TurnOn() (err error) {
	log.Printf("Machine %q, serial: %s is turn On.", m.BrandName, m.SerialNumber)
	return nil
}

func (m *WashingMachine) TurnOf() (err error) {
	log.Printf("Machine %q, serial: %s is turn Off.", m.BrandName, m.SerialNumber)
	return nil
}

// WMCP implements Washing MAchine Controp Panel
type WMCP interface {
}
