package main

import (
	"fmt"
	"math"
)

type Converter struct{}

type Centimeter float64
type Feet float64
type Minutes float64
type Seconds float64
type Celsius float64
type Fahrenheit float64
type Radians float64
type Degrees float64
type Kilograms float64
type Pounds float64

func (cvr Converter) CentimeterToFeet(cm Centimeter) Feet {
	value := cm / 30.48
	return Feet(value)

}
func (cvr Converter) FeetToCentimeter(ft Feet) Centimeter {
	value := ft * 30.48
	return Centimeter(value)
}

func (cvr Converter) MinutestoSeconds(mins Minutes) Seconds {
	value := mins * 60
	return Seconds(value)
}

func (cvr Converter) SecondstoMinutes(secs Seconds) Minutes {
	value := secs / 60
	return Minutes(value)
}

func (cvr Converter) CelstoFahr(cels Celsius) Fahrenheit {
	value := ((cels * 1.8) + 32)
	return Fahrenheit(value)
}

func (cvr Converter) FahrtoCels(fahr Fahrenheit) Celsius {
	value := ((fahr - 32) / 1.8)
	return Celsius(value)
}

func (cvr Converter) RadstoDegs(rads Radians) Degrees {
	value := ((rads * 180) / math.Pi)
	return Degrees(value)
}

func (cvr Converter) DegstoRads(degs Degrees) Radians {
	value := ((degs / 180) * math.Pi)
	return Radians(value)
}

func (cvr Converter) KilostoPounds(kilos Kilograms) Pounds {
	value := kilos * 2.2046
	return Pounds(value)
}

func (cvr Converter) PoundstoKilos(lbs Pounds) Kilograms {
	value := lbs * 0.45359
	return Kilograms(value)
}

func main() {
	cvr := Converter{}
	fmt.Println("30 centimeters is equal to : ", cvr.CentimeterToFeet(30), "feet")
	fmt.Println("20 feet is equal to : ", cvr.FeetToCentimeter(20), "centimeters")
	fmt.Println("45 minute(s) is/are equal to : ", cvr.MinutestoSeconds(45), "seconds")
	fmt.Println("590 second(s) is/are equal to : ", cvr.SecondstoMinutes(590), "minutes")
	fmt.Println("37.5 degree(s) Celsius is equal to : ", cvr.CelstoFahr(37.5), "degree(s) Fahreinheit")
	fmt.Println("101.1 degree(s) Fahrenheit is equal to : ", cvr.FahrtoCels(101.1), "degree(s) Celsius")
	fmt.Println("3.3 radians is equal to : ", cvr.RadstoDegs(3.3), "degrees")
	fmt.Println("295.8 degrees is equal to : ", cvr.DegstoRads(295.8), "radians")
	fmt.Println("99.9 kilograms is equal to : ", cvr.KilostoPounds(99.9), "pounds")
	fmt.Println("782.6 pounds is equal to : ", cvr.PoundstoKilos(782.6), "kilograms")

}
