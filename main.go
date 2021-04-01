package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

//go:embed assets/*
var assets embed.FS

//go:embed www/index.html
var html []byte

func main()  {
	fs := http.FileServer(http.FS(assets))
	http.Handle("/assets/", fs)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		w.Write(html)
	})
	fmt.Println("the address is http://localhost:8080, to end press ctrl+c")
	openbrowser("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}