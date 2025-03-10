package main

import (
// "errors"
"fmt"
// "time"
// "unicode/utf8"
// "strings"
// "math/rand"
// "sync"
// "encoding/json"
// "os"
)

//Structs and interfaces
type gasEngine struct {
	mpg uint16
	gallons uint16
	// ownerInfo owner //or
	// owner
}

type electricEngine struct {
	mpkwh uint16
	kwh uint16
}

// type owner struct {
// 	name string
// }

//Generics with struct
type car [T gasEngine | electricEngine]struct{
	carMake string
	carModel string
	engine T
}

// //Inferring the type of Generic parameter
// type contactInfo struct {
// 	Name string
// 	Email string
// }

// type purchaseInfo struct {
// 	Name string
// 	Price float32
// 	Amount int
// }

// //struct methods
// func (e gasEngine) milesLeft() uint16 {
// 	return e.gallons*e.mpg
// }

// func (e electricEngine) milesLeft() uint16 {
// 	return e.mpkwh*e.kwh
// }

// //Interface
// type engine interface {
// 	milesLeft() uint16
// }

// func canMakeIt(e engine, miles uint16) {
// 	if miles<=e.milesLeft() {
// 		fmt.Println("You can make it there")
// 	} else {
// 		fmt.Println("Need to fuel up first")
// 	}
// }

// //Go Routines
// var m = sync.RWMutex{}
// var wg = sync.WaitGroup{}
// var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
// var results = []string{}

// //Program that mocks checking for sales for chicken fingerd at walmart, costco and wholefoods
// var MAX_CHICKEN_PRICE float32 = 5 
// var MAX_TOFU_PRICE float32 = 5 

func main() {
	//     // Data types
		// var intnum int16 = 32767
	    // intnum = intnum + 1
	    // fmt.Println(intnum)

	//     var floatNum float64 = 12345678.9
	//     fmt.Println(floatNum)

	//     var floatNum32 float32 = 10.1
	//     var intNum32 int32 = 2
	//     var result float32 = floatNum32 + float32(intNum32)
	//     fmt.Println(result)

	//     var intNum1 int = 3
	//     var intNum2 int = 2
	//     fmt.Println(intNum1/intNum2)
	//     fmt.Println(intNum1%intNum2)

	//     //string types
	//     var myString1 string = "Hello World"
	//     fmt.Println(myString1)

	//     var myString2 string = "Hello \nWorld"
	//     fmt.Println(myString2)

	//     var myString3 string = `Hello
	// World`
	//     fmt.Println(myString3)

	// var myString4 string = "Hello" + " " + "World"
	// fmt.Println(myString4)

	//     fmt.Println(len("Chris"))
	//     fmt.Println(utf8.RuneCountInString("Chris"))

	//     var myRune rune = 'c'
	//     fmt.Println(myRune)

	//     var myBoolean bool = false
	//     fmt.Println(myBoolean)

	//     var myVar = "no type"
	//     fmt.Println(myVar)

	//     myVar1 := "no type"
	//     fmt.Println(myVar1)

	//     var var2, var3 int = 1, 2
	//     fmt.Println(var2, var3)

	//     var4, var5 := 1, 2
	//     fmt.Println(var4, var5)

	// printme()

	// var printValues string = "Hello Earth"
	// printMeToo(printValues)

	// var numerator int = 10
	// var denominator int = 2
	// var result, reminder, err = intDivision(numerator, denominator)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else if reminder == 0 {
	// 	fmt.Printf("The result of the integer division is %v\n", result)
	// } else {
	// 	fmt.Printf("The result of the integer division is %v with reminder %v\n", result, reminder)
	// }

	// //Using switch
	// switch {
	// case err != nil:
	//     fmt.Println(err.Error())
	// case reminder == 0:
	//     fmt.Printf("The result of the integer is %v\n", result)
	// default:
	//     fmt.Printf("The result of the integer division is %v with reminder %v", result, reminder)
	// }

	// //using conditional swith

	// switch reminder {
	// case 0:
	//     fmt.Printf("The division was exact")
	// case 1,2:
	//     fmt.Printf("The division was close")
	// default:
	//     fmt.Printf("The division was not close")
	// }

	// //Arrays, slices, maps, and loops
	// //Arrays
	// var intArr [3]int32
	// intArr[1] = 123
	// fmt.Println(intArr[0])
	// fmt.Println(intArr[1:3])

	// //Memory location using address-of operator (&)
	// fmt.Println(&intArr[0])
	// fmt.Println(&intArr[1])
	// fmt.Println(&intArr[2])

	// var intArr1 [3]int32 = [3]int32{1, 2, 3}
	// fmt.Println(intArr1)

	// intArr2 := [3]int{1, 2, 3}
	// fmt.Println(intArr2)

	// intArr3 := [...]int{1, 2, 3, 4}
	// fmt.Println(intArr3)

	// //Slices
	// var intSlice []int32 = []int32{4,5,6}
	// fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	// fmt.Println(intSlice)
	// intSlice = append(intSlice, 7)
	// fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	// fmt.Println(intSlice)

	// var intSlice2 []int32 = []int32{8,9}
	// intSlice2 = append(intSlice, intSlice2...)
	// fmt.Printf("The length is %v with capacity %v\n", len(intSlice2), cap(intSlice2))
	// fmt.Println(intSlice2)

	// var intSlice3 []int32 = make([]int32, 4, 8)
	// fmt.Println(intSlice3)

	// //Maps
	// var myMap map[string]uint8 = make(map[string]uint8)
	// fmt.Println(myMap)

	// var myMap1 = map[string]uint8 {"Adam":23, "Sarah":45}
	// fmt.Println(myMap1)
	// fmt.Println(myMap1["Adam"])
	// // delete(myMap1, "Adam")
	// // fmt.Println(myMap1)
	// var age, ok = myMap1["Chris"]
	// if ok {
	// 	fmt.Printf("The age is %v", age)
	// } else {
	// 	fmt.Println("Invalid name")
	// }

	// //Loops
	// for name, age := range myMap1 {
	// 	fmt.Printf("Name: %v, Age: %v\n", name, age)
	// }

	// //Iterating over arrays
	// for i, v := range intArr1 {
	// 	fmt.Printf("Index: %v, Value: %v\n", i, v)
	// }

	// // While loop type
	// var i int = 0
	// for i < 10 {
	// 	fmt.Print(i)
	// 	i = i + 1
	// }

	// var j int = 0
	// for {
	// 	if j >= 6 {
	// 		break
	// 	}
	// 	fmt.Print(j)
	// 	j = j + 1
	// }

	// for k:=0; k<5; k++ {
	// 	fmt.Print(k)
	// }

	// for l:=range 15 {
	// 	fmt.Print(l)
	// }
	// fmt.Println("")

	// //Speed test
	// var n int = 1000000
	// var testSlice = []int{}
	// var testSlice1 = make([]int, 0, n)

	// fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))
	// fmt.Printf("Total time with preallocation: %v\n", timeLoop(testSlice1, n))

	// //Strings, Runes, Byte
	// //Using String
	// var myString = "résumé"
	// var indexed = myString[0]
	// fmt.Println(myString)
	// fmt.Printf("%v, %T\n", indexed, indexed)
	// for i, v := range myString {
	// 	fmt.Println(i, v)
	// }
	// fmt.Printf("The length of 'myString' is %v", len(myString))

	// //Using Rune
	// var myString1 = []rune("résumé")
	// var indexed1 = myString1[0]
	// fmt.Println(myString1)
	// fmt.Printf("%v, %T\n", indexed1, indexed1)
	// for i, v := range myString1 {
	// 	fmt.Println(i, v)
	// }
	// fmt.Printf("The length of 'myString1' is %v", len(myString1))

	// var myRune1 = "a"
	// fmt.Printf("\nmyRune1 = %v", myRune1)

	// //String Building
	// var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	// //Instead of this
	// var catStr = ""
    // start := time.Now()
	// for i := range strSlice {
	// 	catStr += strSlice[i]
	// }
	// fmt.Printf("\nTime taken to concatenate %v is %v", catStr, time.Since(start))

	// //Do this
	// var strBuilder strings.Builder
	// start1 := time.Now()
	// for i := range strSlice {
	// 	strBuilder.WriteString(strSlice[i])
	// }
	// var catStr1 = strBuilder.String()
	// fmt.Printf("\nTime taken to string build %v is %v", catStr1, time.Since(start1))

	// //Structs and interfaces
	// var myEngine gasEngine = gasEngine{25, 15, owner{"Alex"}, owner{"Jimmie"}}
	// myEngine.mpg = 30
	// fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name, myEngine.name)

	// //Anonymous structs
	// var myEngine1 = struct {
	// 	mpg uint8
	// 	gallons uint8
	// }{25, 20}
	// fmt.Println("mpg:", myEngine1.mpg)
	// fmt.Println("gallons:", myEngine1.gallons)

	// //struct methods
	// fmt.Printf("Total miles left in tank: %v\n", myEngine.milesLeft())

	// //Interfaces
	// canMakeIt(myEngine, 60)

	// var myEngine2 electricEngine = electricEngine{25, 15}
	// canMakeIt(myEngine2, 376)

	// //Pointers
	// var p *int32 = new(int32)
	// fmt.Printf("The value p points to is: %v and p is: %v", *p, p)
	// *p = 5
	// fmt.Printf("\nThe value p points to is: %v and p is: %v", *p, p)
	// var i int32
	// fmt.Printf("\nThe value of i is: %v and address is: %v", i, &i)
	// p = &i
	// *p = 10
	// fmt.Printf("\nThe value p points to is: %v and p is: %v", *p, p)
	// fmt.Printf("\nThe value of i is: %v and address is: %v", i, &i)
	// var k int32 = 20
	// i = k
	// fmt.Printf("\nThe value p points to is: %v and p is: %v", *p, p)
	// fmt.Printf("\nThe value of i is: %v and address is: %v", i, &i)

	// //slices contain pointers under the hood
	// var slice = []int32{1, 2, 3}
	// var sliceCopy = slice
	// sliceCopy[2] = 4
	// fmt.Println(slice)
	// fmt.Println(sliceCopy)

	// //Pointers in Functions
	// var thing1 = [5]float64{1,2,3,4,5}
	// fmt.Printf("\nThe memory allocation of the thing1 array is: %p", &thing1)
	// //Without Pointers
	// var result [5]float64 = square(thing1)
	// fmt.Printf("\nThe result is: %v", result)
	// fmt.Printf("\nThe value of thing1 is: %v", thing1)

	// //With Pointers
	// var result1 [5]float64 = squared(&thing1)
	// fmt.Printf("\nThe result1 is: %v", result1)
	
	// fmt.Printf("\nThe value of thing1 is: %v", thing1)

	// //Go Routines
	// t0 := time.Now()
	// for i := 0; i < len(dbData); i++ {
	// 	wg.Add(1)
	// 	go dbCall(i)
	// }
	// wg.Wait()
	// fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	// fmt.Printf("\nThe results are %v", results)

	// //Channels
	// //Unbuffered Channel
	// var c = make(chan int)
	// go process(c)
	// for i:= range c {
	// 	fmt.Println(i)
	// }

	// //Buffered Channel
	// var ch = make(chan int, 5)
	// go processed(ch)
	// for i:= range ch {
	// 	fmt.Println(i)
	// 	time.Sleep(time.Second*1)
	// }

	// //Program that mocks checking for sales for chicken fingerd at walmart, costco and wholefoods
	// var chickenChannel = make(chan string)
	// var tofuChannel = make(chan string)
	// var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	// for i := range websites {
	// 	go checkChickenPrices(websites[i], chickenChannel)
	// 	go checkTofuPrices(websites[i], tofuChannel)
	// }
	// sendMessage(chickenChannel, tofuChannel)

	// //Generics
	// var intSlice4 = []int{1,2,3}
	// fmt.Println(sumSlice(intSlice4))

	// var float32Slice = []float32{1,2,3.6}
	// fmt.Println(sumSlice(float32Slice))

	// //Generics using type any
	// var intSlice5 = []int{}
	// fmt.Println(isEmpty(intSlice5))

	// var float32Slice2 = []float32{1,2,3}
	// fmt.Println(isEmpty(float32Slice2))

	// //Inferring the type of Generic parameter
	// var contacts []contactInfo = loadJSON[contactInfo]("./contactInfo.json")
	// fmt.Printf("\n%+v", contacts)

	// var purchases []purchaseInfo = loadJSON[purchaseInfo]("./purchaseInfo.json")
	// fmt.Printf("\n%+v", purchases)

	//Generics with struct
	var gasCar = car[gasEngine] {
		carMake: "Honda",
		carModel: "Civic",
		engine: gasEngine{
			gallons: 12,
			mpg: 40,
		},
	}
	var electricCar = car[electricEngine] {
		carMake: "Tesla",
		carModel: "Model 3",
		engine: electricEngine{
			kwh: 57,
			mpkwh: 4,
		},
	}

	fmt.Println(gasCar, electricCar)

	}


// // Functions and structures
//  func printme() {
// 	fmt.Println("Hello World")
// }

// func printMeToo(printValue string) {
// 	fmt.Println(printValue)
// }

// func intDivision(numerator int, denominator int) (int, int, error) {
// 	var err error
// 	if denominator == 0 {
// 		err = errors.New("cannot divide by 0")
// 		return 0, 0, err
// 	}
// 	var result int = numerator / denominator
// 	var reminder int = numerator % denominator
// 	return result, reminder, nil
// }

// func timeLoop(slice []int, n int) time.Duration {
// 	var t0 = time.Now()
// 	for len(slice) < n {
// 		slice = append(slice, 1)
// 	}
// 	return time.Since(t0)
// }

// //Pointers in Functions
// //Without Pointers
// func square(thing2 [5]float64) [5]float64 {
// 	fmt.Printf("\nThe memory allocation of the thing2 array is: %p", &thing2)
// 	for i := range thing2 {
// 		thing2[i] = thing2[i]*thing2[i]
// 	}
// 	return thing2
// }

// //With Pointers
// func squared(thing3 *[5]float64) [5]float64 {
// 	fmt.Printf("\nThe memory allocation of the thing3 array is: %p", thing3)
// 	for i := range thing3 {
// 		thing3[i] = thing3[i]*thing3[i]
// 	}
// 	return *thing3
// }

// // Go Routines
//  func dbCall(i int) {
// 	var delay float32 = 2000
// 	time.Sleep(time.Duration(delay)*time.Millisecond)
// 	save(dbData[i])
// 	log()
// 	wg.Done()
// }

// func save(result string) {
// 	m.Lock()
// 	results = append(results, result)
// 	m.Unlock()
// }

// func log() {
// 	m.RLock()
// 	fmt.Printf("\nThe current results are: %v", results)
// 	m.RUnlock()
// }

// //Unbuffered Channel
// func process(c chan int) {
// 	defer close(c)
// 	for i:=0; i<5; i++ {
// 		c <- i
// 	}
// }

// //Buffered Channel
// func processed(c chan int) {
// 	defer close(c)
// 	for i:=0; i<5; i++ {
// 		c <- i
// 	}
// 	fmt.Println("Exiting Process")
// }

// //Program that mocks checking for sales for chicken fingerd at walmart, costco and wholefoods
// func checkChickenPrices(website string, chickenChannel chan string) {
// 	for {
// 		time.Sleep(time.Second*1)
// 		fmt.Println("Checking Chicken Price")
// 		var chickenPrice = rand.Float32()*20
// 		if chickenPrice <= MAX_CHICKEN_PRICE {
// 			chickenChannel <- website
// 			fmt.Println("Found suitable chicken price")
// 			break
// 		}
// 	}
// }

// func checkTofuPrices(website string, tofuChannel chan string) {
// 	for {
// 		time.Sleep(time.Second*1)
// 		fmt.Println("Checking Tofu Price")
// 		var tofuPrice = rand.Float32()*20
// 		if tofuPrice <= MAX_TOFU_PRICE {
// 			fmt.Println("Found suitable tofu price")
// 			tofuChannel <- website
// 			break
// 		}
// 	}
// }

// func sendMessage(chickenChannel chan string, tofuChannel chan string) {
// 	select {
// 	case website := <-chickenChannel:
// 		fmt.Printf("\nText Sent: Found a deal on chicken at %v", website)
// 	case website := <-tofuChannel:
// 		fmt.Printf("\nEmail Sent: Found a deal on Tofu at %v", website)
// 	}
	
// }

// //Generics
// func sumSlice[T int | float32 | float64](slice []T) T {
// 	var sum T
// 	for _, v := range slice {
// 		sum += v
// 	}
// 	return sum
// }

// //Generics using type any
// func isEmpty[T any](slice []T) bool {
// 	return len(slice)==0
// }

// //Inferring the type of Generic parameter
// func loadJSON[T contactInfo | purchaseInfo](filePath string) []T {
// 	data, _ = os.ReadFile(filePath)

// 	var loaded = []T{}
// 	json.Unmarshal(data, &loaded)

// 	return loaded
// }
