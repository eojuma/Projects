package main
import (
	"os"
	"fmt"
)

func main() {
	file,error:=os.Create("Output.txt")
	if error !=nil{
		fmt.Println("Error create the file",error)
		return
	}
	defer file.Close()
	fmt.Println(file.Name(),"created successfully")
} 