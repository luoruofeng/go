package go_implement_c

import "fmt"

//void Hello(_GoString_ s);
import "C"

//export Hello
func Hello(s string){
	fmt.Println(s)
}

func PredefineApp(){
	Hello("_GoString_ is available after go1.10 version")
}