/* Temperature conversion Utility */
package main

import (
	"fmt"
	// "strconv"
);

func main() {
	// var temperatureF float64
	// temperatureK := 288.15
	// temperatureF = (temperatureK - 273.15 ) * 1.8 + 32

	// fmt.Println("Temperature in Ferenheit: ",temperatureF)

	// var temperatureKnew float64
	// var temperatureFnew float64 = 70
	// temperatureKnew = (temperatureFnew - 32 ) / 1.8 + 273.15

	// fmt.Println("Temperature in Kelvin: ", temperatureKnew)

	/* Math Functions */
	// fmt.Println("Square Root:", math.Sqrt(16)) // 4

	/* Type conversions */
	// Integer to Float
	// var tempInt int = 10;
	// var tempFloat float64 = float64(tempInt);
	// fmt.Println("Integer to Float:", tempFloat);
	// fmt.Printf("Integer to Float [Including decimals]: %.10f\n", tempFloat);
	// fmt.Printf("Data Type: %T\n", tempFloat) /* float 64 */
	// str := strconv.Itoa(80);
	// fmt.Println(str);
	// fmt.Printf("%T\n", str);
	
	// var myStr string = "42"
	// var myIntFromString, _ = strconv.Atoi(myStr);
	// fmt.Println("String to Integer:", myIntFromString);
	// fmt.Printf("%T\n", myIntFromString);

	/* + */ 
	/* String to Float */
	// var myFloatStr string = "3.14159"
	// var myFloatFromString, _ = strconv.ParseFloat(myFloatStr, 64)
	// fmt.Println("String to Float:", myFloatFromString)
	// fmt.Printf("%T\n", myFloatFromString)

	/* Converting Strings to Floats */
	// This accepts any of these as inputs 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False
	// var myBooleanString = "T"
	// var myBooleanFromString, _ = strconv.ParseBool(myBooleanString)
	// fmt.Println(myBooleanFromString)

	// const Agency string = "Fast Tracks";
	// const (
	// 	Agency1 = "Spectra"
	// 	Agency2 = "Gee"
	// 	Agency3 = "Mj"
	// )
	// fmt.Println(Agency)
	// fmt.Println(Agency1, Agency2, Agency3)
	// const (
	// 	Economy = iota
	// 	Compact
	// )
	// fmt.Println("Economy", Economy)
	// fmt.Println("Compact", Compact)
	/* Getting User's Input */
	// fmt.Println("What's your name?")
	// var name string;
	// fmt.Scanln(&name)

	// fmt.Printf("Hello, %v", name)

	// isFast := false
	// isSlow := true
	// if isFast {
	// 	fmt.Println("You're going too fast, slow down")
	// } else if isSlow {
	// 	fmt.Println("You're going too slow, This is dangerous too!")
	// } else {
	// 	fmt.Println("Thank you for driving within the speed limit!")
	// }
	const upperSpeedLimit, lowerSpeedLimit int = 120, 30
	speed := 120
	// if speed == upperSpeedLimit {
	// 	fmt.Println("You're driving at max allowed speed, Be careful")
	// } else if speed == lowerSpeedLimit {
	// 	fmt.Println("You're driving very very slow!")
	// } else if speed > upperSpeedLimit {
	// 	fmt.Println("You're going too fast, slow down!")
	// } else if speed < lowerSpeedLimit {
	// 	fmt.Println("You're going too slow, This is dangerous too!")
	// } else {
	// 	fmt.Println("Thank you for driving within speed limits!")
	// }
	if speed >= upperSpeedLimit {
		fmt.Println("You're going too fast, slow down!")
	} else if speed <= lowerSpeedLimit {
		fmt.Println("You're going too slow, This is dangerous too!")
	} else {
		fmt.Println("Thank you for driving within speed limits!")
	}
}