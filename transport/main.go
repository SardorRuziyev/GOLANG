package main

import (
	"transport/transporter"
)

func main() {
	truck := transporter.Truck{Speed: 60, Capacity: 1000}
	drone := transporter.Drone{Speed: 120, Capacity: 10, BatteryLife: 2}
	ship := transporter.Ship{Speed: 30, Capacity: 5000}

	transporters := []transporter.Transporter{truck, drone, ship}

	truck.Deliver("Toshkent", 300)
	drone.Deliver("Samarqand", 180)
	ship.Deliver("Istanbul", 360)

	transporter.RecommendTransporter("Toshkent", transporters)
}
