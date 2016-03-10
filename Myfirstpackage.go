//Package greetings shows the greetings
package Myfirstpackage

//GreetingsString is a global variable
var Myfirstpackagestring = "Hello World"

//PrintGreetings is a global function
func PrintMyfirstpackage(name string) string {
	return Myfirstpackagestring + "-" + name
}
