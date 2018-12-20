package main
import (
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
	"log"
	"connect/lib/http"
	"connect/lib/udp"
	"connect/lib/tcp"
	"flag"
	"github.com/facebookgo/flagenv"
	"time"
	"fmt"
)
var (
	portBind, timeOutWeb, timeOutCheck int
)

type ViewData struct{
	Title string
}

type CheckStatus struct {
	ViewData
	KindTest []string
	IP string
	TCP [2]bool
	UDP [2]bool
	HTTP [2]bool
	HTTPS [2]bool
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := ViewData{
		Title: "Say Hello!",
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println("Error to load template ...")
	}
	tmpl.Execute(w, data)
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	var tcp_bool, udp_bool, http_bool, https_bool [2]bool
	r.ParseForm()
	service := r.Form["host"][0]
	kind := r.Form["type"]

	for _, v := range kind{
		if v == "tcp"{
			tcp_bool[0] = true
			tcp_bool[1] = tcp.TestTcpConnection(service, timeOutCheck)
		} else if v == "udp" {
			udp_bool[0] = true
			udp_bool[1] = udp.TestUdpConnection(service, timeOutCheck)
		} else if v == "http" {
			http_bool[0] = true
			http_bool[1] = http_check.TestHttpConnection(service, timeOutCheck)
		} else if v == "https" {
			https_bool[0] = true
			https_bool[1] = http_check.TestHttpsConnection(service, timeOutCheck)
		}
	}

	data := ViewData{
		Title: "Say Hello!",
	}

	data_in := CheckStatus{
		ViewData: data,
		KindTest: kind,
		IP: service,
		TCP: tcp_bool,
		UDP: udp_bool,
		HTTP: http_bool,
		HTTPS: https_bool,
	}
	tmpl, err := template.ParseFiles("templates/check.html")
	if err != nil {
		log.Println("Error to load template ...")
	}
	tmpl.Execute(w, data_in)
}


func main() {
	flag.IntVar(&portBind, "port", 8080, "Bind port web server")
	flag.IntVar(&timeOutWeb, "timeout-web", 30, "Timeout from web server")
	flag.IntVar(&timeOutCheck, "timeout-check", 5, "Timeout from check resurce")
	flagenv.Parse()
	flag.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/status", StatusHandler)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("Web Server to 0.0.0.0:", portBind)
	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%v", portBind),
		WriteTimeout: time.Duration(timeOutWeb) * time.Second,
		ReadTimeout:  time.Duration(timeOutWeb) * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}