package main

import "fmt"

// package level constants
const pi float32 = 3.14
const minutesInHour uint8 = 60
const hoursInDay uint = 24
const comp = 2 + 3i

func main() {
	fmt.Println("Package constants:")
	fmt.Println("---------------------------")
	fmt.Println(pi)
	fmt.Println(minutesInHour)
	fmt.Println(hoursInDay)
	fmt.Println(comp)

	explicitTypedConstants()
	inferredTypeConstants()
}

func explicitTypedConstants() {
	fmt.Println("Explicitly typed constants:")
	fmt.Println("---------------------------")
	const printingEnabled bool = true
	const myPi float32 = 3.14
	const myMinutesInHour uint8 = 60
	const myHoursInDay uint = 8
	const number int = 4
	const aBigNumber uint16 = 64353
	const aBigNegNumber int32 = -67353
	const aBiggerNumber float32 = 909938892178012238923475684 // too big for int64 or even uint64
	const aBiggerNegNumber float32 = -909938892178012238923475684
	const favoriteLanguage string = "Go Language"

	printTypeAndValue(printingEnabled)
	printTypeAndValue(myPi)
	printTypeAndValue(myMinutesInHour)
	printTypeAndValue(myHoursInDay)
	printTypeAndValue(number)
	printTypeAndValue(aBigNumber)
	printTypeAndValue(aBigNegNumber)
	printTypeAndValue(aBiggerNumber)
	printTypeAndValue(aBiggerNegNumber)
	printTypeAndValue(favoriteLanguage)
}

func inferredTypeConstants() {
	fmt.Println("\nInferred type constants:")
	fmt.Println("---------------------------")
	const printingEnabled = true
	const myPi = 3.14
	const myMinutesInHour = 60
	const myHoursInDay = 8
	const number = 4
	const aBigNumber = 66353
	const aBigNegNumber = -66353
	const aBiggerNumber = 909938892178012238923475684.0 // too big for int64 or even uint64
	const aBiggerNegNumber = -909938892178012238923475684.0
	const favoriteLanguage = "Go Language"

	printTypeAndValue(printingEnabled)
	printTypeAndValue(myPi)
	printTypeAndValue(myMinutesInHour)
	printTypeAndValue(myHoursInDay)
	printTypeAndValue(number)
	printTypeAndValue(aBigNumber)
	printTypeAndValue(aBigNegNumber)
	printTypeAndValue(aBiggerNumber)
	printTypeAndValue(aBiggerNegNumber)
	printTypeAndValue(favoriteLanguage)
}

func printTypeAndValue(t interface{}) {
	fmt.Printf("value: %v, type: %T\n", t, t)
}
