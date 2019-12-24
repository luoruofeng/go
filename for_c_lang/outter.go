package for_c_lang

//void SayHello(char *name);
import "C"

func OutterApp(){
	C.SayHello(C.CString("use other c file in the same dir, hello.c\n"))
}