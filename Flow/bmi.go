package main

import "fmt"

func main() {

	var weight float32
	var height float32
	fmt.Println("Enter weight in pounds")
	fmt.Scan(&weight)
	fmt.Println("Enter height in inches")
	fmt.Scan(&height)

	const KILOGRAMS_PER_POUND = 0.45359237
	const METERS_PER_INCH = 0.0254

	var weightInKilograms = weight * KILOGRAMS_PER_POUND
	var heightInMeters = height * METERS_PER_INCH
	var bmi = weightInKilograms /
		(heightInMeters * heightInMeters)

	fmt.Println("BMI is ", bmi)

	if bmi < 18.5 {
		fmt.Println("underweight")
	} else if bmi < 25 {
		fmt.Println("Normal")
	} else if bmi < 30 {
		fmt.Println("Overweight")
	} else {
		fmt.Println("Obese")
	}
}
