package transporter

import "fmt"

type Truck struct {
	Speed    float64
	Capacity float64
}

func (t Truck) Deliver(destination string, weight float64) string {
	if weight > t.Capacity {
		return fmt.Sprintf("Truck cannot")
	}
	return fmt.Sprintf("truck delivers")

}

func (t Truck) CalculateETA(distance float64) float64 {
	return distance / t.Speed
}
