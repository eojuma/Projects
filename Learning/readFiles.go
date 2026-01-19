package main
import (
	"fmt"
	"os"
)

func main() {
	file,err:=os.ReadFile("sample.txt")
	if err !=nil{
		fmt.Println("Error reading file",err)
		return
	}
	fmt.Println(string(file))
}