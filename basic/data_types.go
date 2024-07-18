package basic

import "fmt"

/*
Data Types : https://www.geeksforgeeks.org/data-types-in-go/
for print integers: %b-base2; %o-base8; %d-base10
*/
func DataTypes() {
	// Basic types : Numbers : Integers : signed integers
	// (int8, int16, int32(rune), int64)
	signedIntegers()

	// Basic types : Numbers : Integers : unsigned integers
	// (uint8(byte), uint16, uint32, uint64, uintptr)
	unsignedIntegers()

	// Basic types : Numbers : Floating-Point numbers
	// (float32, float64)
	floatingNumbers()

	// Basic types : Numbers : Complex numbers
	// (complex64, complex128)
	complexNumbers()

	// Basic types : Booleans
	// (true, false)
	booleans()

	// Basic types : Strings :  a sequence of immutable bytes
	stringTypes()
}

func stringTypes() {
	str := "Hello world!"
	fmt.Printf("%T %s len:%d\n", str, str, len(str))
	str2 := "Hi user!"
	fmt.Println("String concatenation : ", str+str2)
}

func booleans() {
	st1 := "Hello world"
	st2 := "Hello world!"
	result := st1 == st2
	fmt.Printf("%T %v\n", result, result)

	st2 = "Hello world"
	result = st1 == st2
	fmt.Printf("%T %v\n", result, result)
}

func complexNumbers() {
	var x complex64 = complex(6, 2)
	fmt.Printf("%T %v\n", x, x)

	var y complex128 = complex(6, 2)
	fmt.Printf("%T %v\n", y, y)
}

func floatingNumbers() {
	var x float32 = 3.14
	fmt.Printf("%T %f\n", x, x)

	var y float64 = 3.14
	fmt.Printf("%T %f\n", y, y)
}

func unsignedIntegers() {
	var x uint = 4
	fmt.Printf("%T %d\n", x, x)

	var y uint8 = 4
	fmt.Printf("%T %d\n", y, y)

	var z uint16 = 4
	fmt.Printf("%T %d\n", z, z)

	var i uint32 = 4
	fmt.Printf("%T %d\n", i, i)

	var j uint64 = 4
	fmt.Printf("%T %d\n", j, j)
}

func signedIntegers() {
	var x int = 4
	fmt.Printf("%T %d\n", x, x)

	var y int8 = 4
	fmt.Printf("%T %d\n", y, y)

	var z int16 = 4
	fmt.Printf("%T %d\n", z, z)

	var i int32 = 4
	fmt.Printf("%T %d\n", i, i)

	var j int64 = 4
	fmt.Printf("%T %d\n", j, j)
}
