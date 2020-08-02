package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main(){
	f,err:=os.Create("Name.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer f.Close()
	r:=strings.NewReader("Dr. Strange")
	io.Copy(f,r)
}
