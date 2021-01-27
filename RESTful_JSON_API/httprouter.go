package main

import (
  "fmt"
  "github.com/julienschmidt/httprouter"
  "net/http"
  "log"
  "os"
  "net"
  "strconv"
  "strings"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  fmt.Fprint(w, "Welcom!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func System(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  if ps.ByName("info") == "name" {
    name, err := os.Hostname()
    if err != nil {
      panic(err)
    }
    fmt.Fprintf(w, "System name is %s!\n", name)
  }
  if ps.ByName("info") == "ip" {
    ifaces, err := net.Interfaces()
    if err != nil {
      panic(err)
    }

    var ips []string
    for _, iface := range ifaces {
      if iface.Flags&net.FlagUp == 0 {
        continue // interface down
      }
      if iface.Flags&net.FlagLoopback != 0 {
        continue // loopback interface
      }
      addrs, err := iface.Addrs()
      if err != nil {
        log.Fatal(err)
        continue
      }
      for _, addr := range addrs {
        var ip net.IP
        switch v := addr.(type) {
        case *net.IPNet:
          ip = v.IP
        case *net.IPAddr:
          ip = v.IP
        }
        ip = ip.To4()
        if ip == nil {
          continue // not ipv4
        }
        name := iface.Name
        fmt.Printf(name+" "+ip.String()+"\n")
        ips = append(ips, name+" "+ip.String())
      }
    }
    all_ip := ""
    for i, s := range ips {
      all_ip = all_ip+"\n"+strconv.Itoa(i)+" "+s
    }

    ipString := strings.Join(ips, "\n")
    fmt.Fprintf(w, "System ip:\n"+ipString)
  }
}

func main() {
  router := httprouter.New()
  router.GET("/", Index)
  router.GET("/hello/:name", Hello)
  router.GET("/system/:info", System)

  log.Fatal(http.ListenAndServe(":8080", router))
}
