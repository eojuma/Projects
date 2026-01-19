package main
import (
"os"
"fmt"
"io"

)
func main() {
	file,error:=os.Open("sample.txt")
	if error !=nil{
		fmt.Println("Error opening file",error)
		return 
	}

	buffer:=make([]byte,1024)
	for {
		n,err:=file.Read(buffer)
		if err==io.EOF{
			break
		}
		if err !=nil{
			fmt.Println("Error reading this file",err)
			return
		}
		fmt.Println(string(buffer[:n]))
	}
}