package main

import(
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"time"
)

func responseSize(url string){
	fmt.Println(url)
	resp,e := http.Get(url)
	if e!=nil{
		log.Fatal(e)
	}
	defer resp.Body.Close()
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(len(body))
}

func main(){
	go responseSize("https://www.inhatc.ac.kr")
	go responseSize("https://www.nate.com")
	go responseSize("https://www.google.com")
	go responseSize("https://www.harvard.edu")
	time.Sleep(time.Second * 5)
}

/*
package main
import(
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
)

resp,e := http.Get("https://www.inhatc.ac.kr")
if e != nil {
	log.Fatal(e)
}
defer resp.Body.Close()
body, e := ioutil.ReadAll(resp.Body)
if e != nil {
log.Fatal(e)
}
//fmt.Println(string(body))
fmt.Println(len(body))
}*/
