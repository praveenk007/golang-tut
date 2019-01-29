package main

import (
	"fmt"
)

const kmphMultiple float64 = 1.6094

const u16BitMax float64 = 65535

type car struct {
	gasPedal    uint16 // 0 - 65535
	breakPedal  uint16
	steerWheel  int16 // -32000 - +32000
	topSpeedKmh float64
}

func (c car) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / u16BitMax)
}

func main() {
	car := car{gasPedal: 1234, breakPedal: 2345, steerWheel: 232, topSpeedKmh: 200}
	fmt.Println(car.gasPedal)
	fmt.Println(car.kmh())
}
