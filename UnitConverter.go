package main

import "fmt"
import "math"
//Declare all types, converter struct, and unit types
type Converter struct {}
type (
	Minutes float64
	Seconds float64
	Milliseconds float64
	Centimeter float64
	Feet float64
	Celsius float64
	Fahrenheit float64
	Radian float64
	Degree float64
	Kilogram float64
	Pounds float64
)
func (cvr Converter) MinutesToSeconds(min Minutes) Seconds {
	//unit conversion code
	sec := Seconds(min * 60)
	return sec

}
func (cvr Converter) SecondsToMinutes(sec Seconds) Minutes {
	//unit conversion code
	min := Minutes(sec / 60)
	//convert to integer to split value into minutes and seconds
	myInt := int(sec) 
	myMin := myInt / 60
	mySec := myInt % 60
	//print out minutes an seconds
	fmt.Printf("(%dmins : %dseconds)", myMin, mySec)
	return min
}
func (cvr Converter) SecondsToMilliseconds(sec Seconds) Milliseconds {
	//unit conversion code
	ms := Milliseconds(sec * 1000)
	return ms

}
func (cvr Converter) MillisecondsToSeconds(ms Milliseconds) Seconds {
	//unit conversion code
	sec := Seconds(ms / 1000)
	return sec

}
func (cvr Converter) CentimeterToFeet(c Centimeter) Feet {
	//unit conversion code
	feet := Feet(c/30.48)
	return feet

}
func (cvr Converter) FeetToCentimeter(ft Feet) Centimeter {
	//unit conversion code
	cent := Centimeter(ft * 30.48)
	return cent
}
func (cvr Converter) CelsiusToFahrenheit(Cel Celsius) Fahrenheit {
	//unit conversion code
	fht := Fahrenheit((Cel * 9/5)+32)
	return fht

}
func (cvr Converter) FahrenheitToCelsius(fht Fahrenheit) Celsius {
	//unit conversion code
	Cel := Celsius((fht - 32) * 5/9)
	return Cel

}
func (cvr Converter) RadianToDegree(rad Radian) Degree {
	//import exact PI value from math
	PI := Radian(math.Pi)
	//unit conversion code
	deg := Degree(rad * (180/PI))
	return deg

}
func (cvr Converter) DegreeToRadian(deg Degree) Radian {
	//import exact PI value from math
	PI := Degree(math.Pi)
	//unit conversion code
	rad := Radian(deg * (PI/180))
	return rad

}
func (cvr Converter) KilogramToPounds(kg Kilogram) Pounds {
	//unit conversion code
	pds := Pounds(kg * 2.20462)
	return pds
}
func (cvr Converter) PoundsToKilogram(pds Pounds) Kilogram {
	//unit conversion code
	kg := Kilogram(pds / 2.20462)
	return kg
}

func main() {
	cvr := Converter{}
	fmt.Println(cvr.MinutesToSeconds(34.7))
	fmt.Println(cvr.SecondsToMinutes(500))
	fmt.Println(cvr.SecondsToMilliseconds(12))
	fmt.Println(cvr.MillisecondsToSeconds(2.5))
	fmt.Println(cvr.CentimeterToFeet(7))
	fmt.Println(cvr.FeetToCentimeter(198))
	fmt.Println(cvr.CelsiusToFahrenheit(56))
	fmt.Println(cvr.FahrenheitToCelsius(34))
	fmt.Println(cvr.RadianToDegree(90))
	fmt.Println(cvr.DegreeToRadian(76))
	fmt.Println(cvr.KilogramToPounds(16))
	fmt.Println(cvr.PoundsToKilogram(18))
}