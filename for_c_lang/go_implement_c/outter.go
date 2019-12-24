package go_implement_c

import "fmt"

//#include <hello.h>
import "C"

//export SayHelloImpl
func SayHelloImpl(s *C.char){
	fmt.Println(C.GoString(s))
}

func OutterApp(){
	SayHelloImpl(C.CString("golang implement header of C, hello.h"))
}