// Infinite web page server
// All requests to this web server will respond with an infinite HTML page.
// Interesting tool to see how HTTP clients cope with a misbehaving server.
package main
import (
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
)
const (
    Addr = ":8012"
)
func Robots(w http.ResponseWriter, req * http.Request) {
    // Fuck robots
    w.Header().Set("Content-type", "text/plain")
    fmt.Fprintf(w, "User-agent: *\nDisallow: /\n")
}
func Infinite(w http.ResponseWriter, req * http.Request) {
    log.Printf("<- [%s][%s][%s]\n", req.RemoteAddr, req.RequestURI, req.UserAgent())
    w.Header().Set("Connection", "Keep-Alive")
    w.Header().Set("Content-type", "text/html")
    w.Write([]byte("<!DOCTYPE html>\n<html>\n<body>\n"))
    for {
        w.Write([]byte("1234567890123456789012345678901234567890<br>\n"))
        time.Sleep(10 * time.Millisecond)
    }
    // Never reached
    return
}
func main() {
    addr:=Addr
    if len(os.Args)>1 {
        addr=os.Args[1]
    }
	http.HandleFunc("/robots.txt", Robots)
    http.HandleFunc("/", Infinite)
    fmt.Println("Listening on", addr)
    err := http.ListenAndServe(addr, nil)
    if err!=nil {
        fmt.Println(err)
    }
}
