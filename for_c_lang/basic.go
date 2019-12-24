package for_c_lang

/*
	#include <stdio.h>

	static void SayHello(const char *name){
		puts(name);
	}
*/
import "C"

func BasicApp(){
	C.SayHello(C.CString("Hello,World!\n"))
}