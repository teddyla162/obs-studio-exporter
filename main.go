package main

import "C"

// Export a simple function that OBS can call
//export Init
func Init() {
    // This will be called when OBS loads your plugin
}

//export GetName
func GetName() *C.char {
    return C.CString("OBS Studio Exporter")
}

func main() {
    // Required for c-shared build mode
}
