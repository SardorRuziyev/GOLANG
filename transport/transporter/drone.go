package transporter

import (
	"fmt"
)

type Drone struct {
	Speed       float64
	Capacity    float64
	BatteryLife float64
}

func (d Drone) Deliver(destination string, weight float64) string {
	if weight > d.Capacity {
		return fmt.Sprintf("Drone cant")
	}

	return fmt.Sprintf("drone can deliver")
}

func (d Drone) CalculateETA(distance float64) float64 {
	return distance / d.Speed
}
