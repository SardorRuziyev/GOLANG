package transporter

import "fmt"

type Transporter interface {
	Deliver(destination string, weight float64) string
	CalculateETA(distance float64) float64
}

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

type Drone struct {
	Speed       float64
	Capacity    float64
	BatteryLife float64
}

func (d Drone) Deliver(destination string, weight float64) string {
	if weight > d.Capacity {
		return fmt.Sprintf("Drone cant")
	}
	if d.CalculateETA(destination) > d.BatteryLife {
		return fmt.Sprintf("drone cannot")
	}
	return fmt.Sprintf("drone can deliver")
}

func (d Drone) CalculateETA(distance float64) float64 {
	return distance / d.Speed
}

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
func distanceTo(destination string) float64 {
	distances := map[string]float64{
		"Toshkent":  400.0,
		"Samarqand": 600.0,
		"Istanbul":  5000.0,
	}
	return distances[destination]
}
