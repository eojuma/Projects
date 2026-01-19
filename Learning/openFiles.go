package main
import (
	"os"
	"fmt"
)
func main() {
	data,err:=os.Open("sample.txt")
	if err !=nil{
		fmt.Println("Error opening",err)
		return
	}
	defer data.Close()
	fmt.Println("Successfully opened:",data.Name())
}