/*
Avtomobilni boshqarish tizimi
Vazifa:
Avtomobilni boshqarish uchun VehicleController interfeysini yarating, unda quyidagi metodlar bo‘lsin:

StartEngine() string
Drive(distance float64) string
StopEngine() string
Keyin ushbu interfeysni implement qiluvchi quyidagi transport vositalarini yozing:

Car: benzinli avtomobil.
ElectricCar: elektr avtomobil.
Truck: yuk tashuvchi transport vositasi.
Har bir transport vositasining o‘ziga xos haydash uslubi va yoqilg‘i turi bo‘lsin.
*/
package main

import "fmt"

type VehicleController interface {
	StartEngine() string
	Drive(distance float64) string
	StopEngine() string
}
type Car struct {
	Fueltype string
}

func (c *Car) StartEngine() string {
	return fmt.Sprintf("Car: %s starts engine", c.Fueltype)
}
func (c *Car) Drive(distance float64) string {
	return fmt.Sprintf("Car: Driving %.2f km with gasoline.", distance, c.Fueltype)

}
func (c *Car) StopEngine() string {
	return "engine stopped"
}

type ElectricCar struct {
	BatteryType string
}

func (e *ElectricCar) StartEngine() string {
	return fmt.Sprintf("E-Car: %s starts engine", e.BatteryType)
}

func (e *ElectricCar) Drive(distance float64) string {
	return fmt.Sprintf("E-Car: Driving %.2f km with gasoline.", distance, e.BatteryType)

}
func (e *ElectricCar) StopEngine() string {
	return "engine stopped"

}

type Truck struct {
	Fueltype string
}

func (t *Truck) StartEngine() string {
	return fmt.Sprintf("Car: %s starts engine", t.Fueltype)
}
func (t *Truck) Drive(distance float64) string {
	return fmt.Sprintf("Car: Driving %.2f km with gasoline.", distance, t.Fueltype)

}
func (t *Truck) StopEngine() string {
	return "engine stopped"
}

func main() {
	var vehicle VehicleController
	fmt.Println("Choose a vehicle: 1 - Car, 2 - ElectricCar, 3 - Truck")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		vehicle = &Car{Fueltype: "benzin"}
	case 2:
		vehicle = &ElectricCar{BatteryType: "L-ion"}
	case 3:
		vehicle = &Truck{Fueltype: "d"}
	default:
		fmt.Println("Invalid choice!")
		return
	}
	fmt.Println(vehicle.StartEngine())
	fmt.Println(vehicle.Drive(120.5))
	fmt.Println(vehicle.StopEngine())
}
