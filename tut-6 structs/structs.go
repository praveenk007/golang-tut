package main

import (
	"fmt"
)

type car struct {
	gasPedal    uint16 // 0 - 65535
	breakPedal  uint16
	steerWheel  int16 // -32000 - +32000
	topSpeedKmh float64
}

func main() {
	car := car{gasPedal: 1234, breakPedal: 2345, steerWheel: 232, topSpeedKmh: 200}
	fmt.Println(car.gasPedal)
}
