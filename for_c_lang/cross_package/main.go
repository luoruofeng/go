package main

import cgo_help_dir "go_demo/for_c_lang/cross_package/cgo_helper_dir"

//static const char* cs = "hello";
import "C"

func main() {
	//运行将会报错，因为C.cs是*main.C.char,而方法接受*cgo_helper_dir.C.char
	cgo_help_dir.PrintCString(C.cs)
}
