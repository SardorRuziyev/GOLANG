package transporter

import "fmt"

type Ship struct {
	Speed    float64
	Capacity float64
}

func (s Ship) Deliver(destination string, weight float64) string {
	if weight > s.Capacity {
		return fmt.Sprintf("Ship cannot")
	}
	return fmt.Sprintf("Ship deliver")
}

func (s Ship) CalculateETA(distance float64) float64 {
	return distance / s.Speed
}
