package main

import (
	"fmt"
	// "io/ioutil"
	// "net/http"
	"net/url"
)

func main() {
	const myurl string = "https://www.google.com/search?q=hello&sca_esv=6b5001a3b62373fd&sxsrf=AHTn8zpxHw5EGSCtIPUL_tBJj2uAmZegKA%3A1746979720875&source=hp&ei=iMsgaLaiM9Gw4-EPuquqsQE&iflsig=ACkRmUkAAAAAaCDZmNL7UEkkk3tgeKv_A6TZVSRo-mX3&ved=0ahUKEwi2hsPN5puNAxVR2DgGHbqVKhYQ4dUDCBk&uact=5&oq=hello&gs_lp=Egdnd3Mtd2l6IgVoZWxsbzIIEAAYgAQYsQMyCBAuGIAEGLEDMgUQABiABDIIEAAYgAQYsQMyCxAuGIAEGLEDGNQCMggQABiABBixAzIFEAAYgAQyCxAAGIAEGLEDGIoFMgoQLhiABBgCGMsBMggQLhiABBixA0i7HFD4C1ihEnABeACQAQCYAZMBoAG8BaoBAzAuNbgBA8gBAPgBAZgCBqACzAWoAgrCAgcQIxgnGOoCwgIEECMYJ8ICChAjGIAEGCcYigXCAgoQABiABBhDGIoFwgILEAAYgAQYkQIYigXCAgsQLhiABBixAxiDAcICDhAuGIAEGLEDGNEDGMcBwgINEC4YgAQYsQMYQxiKBZgDBfEF5TLQM1nTYoeSBwMxLjWgB5EzsgcDMC41uAfHBQ&sclient=gws-wiz"
	fmt.Printf("web request ")
	// response , err := http.Get(myurl)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Print(response)
	// defer response.Body.Close()
	// data , err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// content := string(data)
	// 	fmt.Printf("response content is %s",content)
	parsedurl, _ := url.Parse(myurl)
	fmt.Println(parsedurl.Scheme)
	fmt.Println(parsedurl.Host)
	fmt.Println(parsedurl.Path)
	fmt.Println(parsedurl.Port())
	fmt.Printf("query \t")
	// fmt.Println(parsedurl.Query().Get("q"))
}
