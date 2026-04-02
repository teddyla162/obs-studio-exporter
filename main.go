package main

import "C"
import (
    "fmt"
    "net/http"
)

// This is called automatically by OBS when the plugin loads
//export obs_module_load
func obs_module_load() C.bool {
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
        // Get stats from OBS
        fmt.Fprintf(w, "# HELP obs_frames_total Total frames\n")
        fmt.Fprintf(w, "# TYPE obs_frames_total counter\n")
        fmt.Fprintf(w, "obs_frames_total 0\n")
        fmt.Fprintf(w, "# HELP obs_dropped_frames_total Dropped frames\n")
        fmt.Fprintf(w, "# TYPE obs_dropped_frames_total counter\n")
        fmt.Fprintf(w, "obs_dropped_frames_total 0\n")
        fmt.Fprintf(w, "# HELP obs_bitrate_bytes_total Total bytes sent\n")
        fmt.Fprintf(w, "# TYPE obs_bitrate_bytes_total counter\n")
        fmt.Fprintf(w, "obs_bitrate_bytes_total 0\n")
    })
    http.ListenAndServe(":9407", nil)
}

func main() {
    // Required for c-shared build mode
}
