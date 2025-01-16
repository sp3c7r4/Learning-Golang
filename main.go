/* Temperature conversion Utility */
package main

import "fmt"

// "crypto/rand"
// "fmt"
// "math/rand"
// "strings"
// "strings"
// "strconv";

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
	// const upperSpeedLimit, lowerSpeedLimit int = 120, 30
	// speed := 120
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
	// if speed >= upperSpeedLimit {
	// 	fmt.Println("You're going too fast, slow down!")
	// } else if speed <= lowerSpeedLimit {
	// 	fmt.Println("You're going too slow, This is dangerous too!")
	// } else {
	// 	fmt.Println("Thank you for driving within speed limits!")
	// }

	/* Logical Operators */
	/* Write a program to classify the days of a Week */

	// isWeekend := true 
	// isHoliday := false
	// if isWeekend && isHoliday {
	// 	fmt.Println("It's a holiday on a weekend!!")
	// } else if isWeekend && !isHoliday {
	// 	fmt.Println("It's a regular weekend day.")
	// } else if !isWeekend && isHoliday {
	// 	fmt.Println("It's a weekday but a holiday")
	// } else {
	// 	fmt.Println("It's a regular day!")
	// }

	// var isWeekend bool = true
	// var isHoliday bool = false;
	// isDayoff := isWeekend || isHoliday;

	// if isDayoff {
	// 	fmt.Println("It's a day off!")
	// }

	// /* Program that interpretes color meanings */
	// color := "yellow"
	// switch color {
	// case "red":
	// 	fmt.Println("Red: Passion and Energy")
	// case "green":
	// 	fmt.Println("Green: Growth and Harmony")
	// case "blue":
	// 	fmt.Println("Blue: Calm and Trust")
	// case "yellow":
	// 	fmt.Println("Yellow: Optimism and Happiness")
	// default:
	// 	fmt.Println("Unknown Color")
	// } 
	/* Day checker Program */
	// day := "Tue"
	// switch day {
	// case "Mon", "Tue", "Wed", "Thu", "Fri":
	// 	fmt.Println("Week Day")
	// case "Sat", "Sun":
	// 	fmt.Println("Weekend")
	// default:
	// 	fmt.Println("Invalid day")
	// }
	// age := 25;
	// switch {
	// case age < 18:
	// 	fmt.Println("Minor")
	// case age >= 18 && age < 60:
	// 	fmt.Println("Adult")
	// case age >= 60:
	// 	fmt.Println(("Senior"))
	// }

	/* 
	User access Permissions 
		+ Admin [ Full Access ]
		+ Editor [ Publish, write, view]
		+ Contributor [ Write, View ]
		+ Viewer [ View ]
	*/
		// userRole := "editor";
		// fmt.Printf("Access permissions for role '%s':\n", userRole)
		// switch userRole {
		// case "admin":
		// 	fmt.Println("- Full system access")
		// 	/* The fallthrough keyword tells the program to continue to the next case without checking this conditon */
		// 	fallthrough;
		// case "editor":
		// 	fmt.Println("- Can publish content")
		// 	fallthrough;
		// case "contributor":
		// 	fmt.Println("- Can write content")
		// 	fallthrough;
		// case "viewer":
		// 	fmt.Println("- Can view content")
		// default:
		// 	fmt.Println("- No access")
		// }

		/* Looping[for] */
		// for i := 0; i < 5; i++ {
		// 	fmt.Println("Simple loop iteration:",i)
		// }

		// 	for j := 0; j < 10; j++ {
		// 		if j%2 == 0 { // Skips even Numbers
		// 			/* Skips the remaining body of the Iteration and Starts another iteration instance */
		// 			continue;
		// 		}
		// 		fmt.Println("Odd number:",j)
		// 	}
	// i := 0
	// for {
	// 	fmt.Println("Infinite loop, iteration:", i)
	// 	i++
	// 	/* Break: Allows the loop to Exit Immediately */
	// 	if i == 3 { break; } // Breaks Infinite Loop
	// }

	/* Another Loop Example */
	// y := 0
	// for y < 3 {
	// 	fmt.Println("Loop Example:", y)
	// 	y++
	// }

	/* Rock Paper Scissors */
	// var computerChoice string;
	// Game()

	/* Arrays */
	// Declaring an Array
	// var bodyTypes [3]string;
	// fmt.Println(bodyTypes) // Returns an Array with 3 Empty Strings
	
	// /* Adding Array items using the Indexing Syntax */
	// bodyTypes[0] = "Sedan"
	// bodyTypes[1] = "SUV"
	// bodyTypes[2] = "Convertible"
	// /* Declaring and Initializing Array in one go */
	// bodyTypes := [3]string{"Sedan","SUV","Convertible"}
	// fmt.Println(bodyTypes)

	// /* Accessing/Getting individual Items of an Array */
	// fmt.Println(bodyTypes[0])
	// fmt.Println(bodyTypes[1])
	// fmt.Println(bodyTypes[2])

	/* 2D Arrays */
	// 2-Dimensional (2D) Arrays
	// var carFleet [3][2]string;
	// carFleet[0] = [2]string{"5 Spectra Gee", "unlimited Brains"}
	// carFleet[1] = [2]string{"3 SUVs Available", "4 SUVs Booked"}
	// carFleet[2] = [2]string{"1 SUVs Available", "1 Convertible Booked"}

	// fmt.Println(carFleet)

	// /* Slices */
	// fuelTypes := []string{"Electric", "Gasoline", "Hybrid"}
	// fmt.Println(fuelTypes)

	// /*
	// append => This allows us to add items to slices
	// + Note: append doesn't modify the original slice it creates a new slice with the new data that has been added
	// */
	// fuelTypes = append(fuelTypes, "Diesel")
	// fmt.Println(fuelTypes)
	// /*
	// Append also allows us to add multiple data to a slice
	// */

	// fuelTypes = append(fuelTypes, "Spectra", "Gee")
	// fmt.Println(fuelTypes)

	/* Another Example */
	/*
	make => This is a built-in function and we can use it to initialize a slice with a redefined length
	*/
	// fuelTypes := make([]string, 3)
	// fmt.Println(len(fuelTypes))
	// fmt.Println("The slice is nil:", fuelTypes == nil)
	// fmt.Println(fuelTypes)

	// fuelTypes := []string{"Gasoline", "Diesel", "Electric", "Hybrid", "Hydrogen"}
	// popular := fuelTypes[0:2];
	// fmt.Println("Popular:",popular)

	// clean := fuelTypes[2:]
	// fmt.Println("Clean:",clean)

	// electricHybrid := fuelTypes[2:4]
	// fmt.Println("Electric Hybrid:",electricHybrid)
  /* Maps */
	// make(map[keyDataType]valueDataType)
	// carInventory := make(map[string]int)
	// carInventory["Sedan"] = 25
	// carInventory["SUV"] = 15
	// carInventory["Convertible"] = 10

	// fmt.Println("Car Inventory:",carInventory)

	/* Initializing maps in one statement */
	// carInventory := map[string]int{
	// 	"Sedan":25,
	// 	"SUV":15,
	// 	"Convertible":10,
	// }

	// fmt.Println("Types of card:",len(carInventory))
	// carInventory["Sedan"] = 20
	
	//Removing Entries from a Map we use the delete() keyword
	// delete(carInventory, "SUV")
	
	// Checking if a Key Exists
	// numberOfSedans, sedansFound := carInventory["Sedan"]
	// if sedansFound { fmt.Printf("Sedans Exists in our Map | We have %v sedans\n", numberOfSedans) }
	
	/* Short Form */
	// if numberOfSedans, ok := carInventory["Sedan"]; ok {
	// 	fmt.Printf("Sedans Exists | We have %v Sedans\n", numberOfSedans)
	// }

	// fmt.Printf("We have %v Sedans\n", numberOfSedans)
	
	// If the key doesn't exist you get the 0 value of the value type
	// fmt.Printf("We have %v Coupes\n", carInventory["Coupe"])
	
	// fmt.Println("Car Inventory:",carInventory)
	
	/* Clearing all Map Values using the clear() function */
	// clear(carInventory)
	// fmt.Println("Cleared Car Inventory:",carInventory)
	
} 