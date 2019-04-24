//+build !linux

package rpio

import "errors"

type Pin uint8

func (p Pin) Low() {

}

func (p Pin) High() {

}

func (p Pin) Output() {

}

func Open() (error) {
	return errors.New("Invalid OS")
}

func Close() {

}