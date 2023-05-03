package main

import "fmt"

func main() {
	fmt.Print("================= Variables =================\n\n")

	var username string = "codewithwaqar"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n\n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n\n", isLoggedIn)

	var uintValue8 uint8 = 255
	fmt.Println(uintValue8)
	fmt.Printf("Variable is of type: %T\n\n", uintValue8)

	var uintValue uint = 500
	fmt.Println(uintValue)
	fmt.Printf("Variable is of type: %T\n\n", uintValue)

	var intValue8 int8 = 127
	fmt.Println(intValue8)
	fmt.Printf("Variable is of type: %T\n\n", intValue8)

	var intValue int = 500
	fmt.Println(intValue)
	fmt.Printf("Variable is of type: %T\n\n", intValue)

	var floatValue32 float32 = 255.1234588888
	fmt.Println(floatValue32)
	fmt.Printf("Variable is of type: %T\n\n", floatValue32)

	var floatValue64 float64 = 255.1234567891234588888
	fmt.Println(floatValue64)
	fmt.Printf("Variable is of type: %T\n\n", floatValue64)

	fmt.Print("================= Default values =================\n\n")

	var defaultValueString string // ""
	fmt.Println(defaultValueString)

	var defaultValueBool bool // false
	fmt.Println(defaultValueBool)

	var defaultValueInt int // 0
	fmt.Println(defaultValueInt)

	var defaultValueFloat float32 // 0
	fmt.Println(defaultValueFloat)
	fmt.Println()

	fmt.Print("================= Aliases =================\n\n")
	// byte --------> alias for uint8
	// rune --------> alias for int32
	var aliasByte byte = 255
	fmt.Println(aliasByte)
	fmt.Printf("Variable is of type: %T\n\n", aliasByte)

	var aliasRune rune = 500
	fmt.Println(aliasRune)
	fmt.Printf("Variable is of type: %T\n\n", aliasRune)

	fmt.Print("================= Implicit Type =================\n\n")

	var implicitType = "Impilicit Type"
	fmt.Println(implicitType)
	fmt.Println()

	fmt.Print("================= No var style =================\n\n")
	noVarStyle := "No var style"
	fmt.Println(noVarStyle)

	fmt.Print("================= Constants =================\n\n")
	const pi = 3.14
	fmt.Println(pi)
}
