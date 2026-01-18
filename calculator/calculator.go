package main
import ("fmt"
		"strconv"
		"os"
)
func main() {
if len(os.Args) !=4{
	fmt.Println("Usage:Enter first number(a),the operation(*,-,/,+)the second number(n)")
	return 
}
operation:=os.Args[2]

a,error1:=strconv.Atoi(os.Args[1])
n,error2:=strconv.Atoi(os.Args[3])

if error1 !=nil || error2 !=nil{
	fmt.Println("Enter valid  numbers")
	return 
}

if operation =="-"{
	fmt.Println(subtract(a,n))
}else if operation =="+"{
	fmt.Println(add(a,n))
}else if operation =="x"{ // Shells (like Bash or Zsh) expands * before Go sees it,so i decided to use X for user inputs but the multiplication operation remains the same.
	fmt.Println(multiply(a,n))
}else if operation =="/"{
	if n==0{
		fmt.Println("Math error:cannot divide a number by 0")
	}else{

		fmt.Println(divide(a,n))
	}
}else if operation =="%"{
	if n==0{
		fmt.Println("Math error:cannot find the modulus of a number divided by 0")
	}else{

		fmt.Println(modulus(a,n))
	}
}else{
	fmt.Println("Invalid operation")
}
}

func add(a,n int)int{
	return a+n
}
func subtract(a,n int)int{
	return a-n
}
func multiply(a,n int)int{
	return a*n
}
func divide(a,n int)int{
	return a/n
}
func modulus(a,n int)int{
	return a%n
}