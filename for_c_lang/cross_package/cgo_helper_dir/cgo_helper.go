package cgo_help_dir

//#include <stdio.h>
import "C"

type CChar C.char

func (p *CChar)GoString() string{
	return C.GoString((*C.char)(p))
}

func PrintCString(cs *C.char){
	C.puts(cs)
}

func PrintCStringRight(cs string){
	C.puts(C.CString(cs))
}