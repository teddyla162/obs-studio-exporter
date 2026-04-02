package main

import "C"

//export Init
func Init() {
    // Add your plugin initialization
}

//export GetName
func GetName() *C.char {
    return C.CString("OBS Studio Exporter")
}

// Add more exported functions here
//export StartExport
func StartExport() {
    // Your export logic
}

func main() {}
