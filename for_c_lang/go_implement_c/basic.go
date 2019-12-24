package go_implement_c

import "fmt"

//void SayHello(char *name);
import "C"

//export SayHello
func SayHello(s *C.char){
	fmt.Println(C.GoString(s))
}

func BasicApp() {
	SayHello(C.CString("Implement C"))
}
