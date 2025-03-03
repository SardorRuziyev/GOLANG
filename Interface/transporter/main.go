package main

import "fmt"

func main() {
	truck := Truck{Speed: 60, Capacity: 1000}
	drone := Drone{Speed: 120, Capacity: 10, BatteryLife: 2}
	ship := Ship{Speed: 30, Capacity: 5000}

	transporters := []Transporter{truck, drone, ship}

	truck.Deliver("Toshkent", 300, 800)
	drone.Deliver("Samarqand", 180, 5)
	ship.Deliver("Istanbul", 360, 4000)

	fastest := FindFastestTransporter(transporters)
	fmt.Printf("Eng tez yetkazib beruvchi transport: %T\n", fastest)

}
