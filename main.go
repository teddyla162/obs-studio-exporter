package main

import "C"
import (
    "fmt"
    "net/http"
    "os"
)

// This is called automatically by OBS when the plugin loads
//export obs_module_load
func obs_module_load() C.bool {
    // Write to a log file so we can see if this runs
    f, _ := os.Create("C:\\obs_plugin_debug.log")
    f.WriteString("obs_module_load was called!\n")
    f.Close()
    
    go StartMetricsServer()
    return C.bool(true)
}

//export Init
func Init() {
    // Plugin initialization
}

//export GetName
func GetName() *C.char {
    return C.CString("OBS Studio Exporter")
}

//export StartRecording
func StartRecording() {
    // Your recording logic here
}

//export StopRecording
func StopRecording() {
    // Your stop recording logic here
}

//export StartMetricsServer
func StartMetricsServer() {
    http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "# HELP obs_frames_total Total frames\n")
        fmt.Fprintf(w, "# TYPE obs_frames_total counter\n")
        fmt.Fprintf(w, "obs_frames_total 0\n")
        fmt.Fprintf(w, "# HELP obs_dropped_frames_total Dropped frames\n")
        fmt.Fprintf(w, "# TYPE obs_dropped_frames_total counter\n")
        fmt.Fprintf(w, "obs_dropped_frames_total 0\n")
    })
    http.ListenAndServe(":9407", nil)
}

func main() {}
