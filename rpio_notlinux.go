//+build !linux

package rpio

import "errors"

type Pin uint8

type State uint8

const (
	Low State = iota
	High
)


type Mode uint8

const (
	Input Mode = iota
	Output
	Clock
	Pwm
	Spi
)

func (p Pin) Mode(m Mode) {

}

func (p Pin) Low() {

}

func (p Pin) High() {

}

func (p Pin) Output() {

}

func (p Pin) Input() {

}

func (p Pin) Read() State {
	return Low
}

func Open() (error) {
	return errors.New("Invalid OS")
}

func Close() {

}