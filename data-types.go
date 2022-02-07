package main  
import "fmt"  
func main() {  
   var i int  
   var f float64  
   var b bool  
   var s string  
   fmt.Printf("%T %T %T %T\n",i,f,b,s) // Prints type of the variable  
   fmt.Printf("%v   %v      %v  %q     \n", i, f, b, s) //prints initial value of the variable  
}  

//types of variable = %T

//intial values
// int = %v
// float = %v
// bool = %v
// string = %q

//These are the 25 keywords for Go-code:

// break,	default,	func,	interface,	select
// case	defer,	go	,map	,struct
// chan,	else,	goto,	package,	switch
// const,	fallthrough,	If,	range,	type
// continue,	for,	import,	return,	var

//Programs consist of keywords, constants, variables, operators, types and functions.

// The following delimiters are used in constructs such as parentheses ( ), brackets [ ] and braces { }.

// The following punctuation characters . , ; : and ... are used

// append,	bool,	byte,	cap,	close,	complex,	complex64,	complex128,	uint16,
// copy,	false,	float32,	float64,	imag,	int,	int8,	int16,	uint32,
// int32,	int64,	iota,	len	make,	new,	nil,	panic,	uint64
// print,	println,	real,	recover,	string,	true,	uint,	uint8,	Uintptr