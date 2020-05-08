package domain1

import (
	"log"
	"math/rand"
	"time"
)

//
type Man struct {
	Name string
}

func (m *Man) TurnOnWashingMachine(cp WMCP) {
	cp.TurnOn()
}
func (m *Man) TurnOffWashingMachine(cp WMCP) {
	cp.TurnOff()
}

//
type WashingMachine struct {
	BrandName    string
	ModelName    string
	SerialNumber string
	Capacity     uint
}

func (m *WashingMachine) TurnOn() (err error) {
	log.Printf("Washing machine serial: %s turn On.", m.SerialNumber)
	m.makeWashing()
	return nil
}

func (m *WashingMachine) TurnOff() (err error) {
	log.Printf("Waching machine serial: %s turn Off.", m.SerialNumber)
	return nil
}
func (m *WashingMachine) makeWashing() (err error) {
	log.Printf("Waching machine serial: %s start washing.", m.SerialNumber)
	time.Sleep(time.Second*4 + time.Second*time.Duration(rand.Intn(4)))
	log.Printf("Waching machine serial: %s finish washing.", m.SerialNumber)
	return nil
}

// WMCP implements Washing MAchine Controp Panel
type WMCP interface {
	TurnOn() error
	TurnOff() error
}
