/* Temperature conversion Utility */
package main

import (
	"fmt"
	"strconv"
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
	str := strconv.Itoa(80);
	fmt.Println(str);
	fmt.Printf("%T\n", str);
	
	var myStr string = "42"
	var myIntFromString, _ = strconv.Atoi(myStr);
	fmt.Println("String to Integer:", myIntFromString);
	fmt.Printf("%T\n", myIntFromString);

	/* + */ 
	
}