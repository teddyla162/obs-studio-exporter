package main

import "C"

//export Init
func Init() {
    // Plugin initialization
}

//export GetName  
func GetName() *C.char {
    return C.CString("OBS Exporter")
}

func main() {
    // Required for c-shared build
}
