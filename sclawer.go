package main

import (
	"fetchhtml"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

var (
	HtmlTemplate_Folder string = "template"
	CurrentPath, _             = os.Getwd()
)

type MyRoute struct {
}

func (mr *MyRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Serve static files

	if match, _ := regexp.MatchString("/static/", r.URL.Path); match {
		// Supply file server for static css and js
		http.StripPrefix("/static/", http.FileServer(http.Dir("static")))

		// Read server static file
		urlpath := r.URL.Path[1:]
		//log.Println(urlpath)
		data, err := ioutil.ReadFile(string(urlpath))

		// Set content type according file type
		if err == nil {
			var contentType string

			if strings.HasSuffix(urlpath, ".css") {
				contentType = "text/css"
			} else if strings.HasSuffix(urlpath, ".html") {
				contentType = "text/html"
			} else if strings.HasSuffix(urlpath, ".js") {
				contentType = "application/javascript"
			} else if strings.HasSuffix(urlpath, ".png") {
				contentType = "image/png"
			} else if strings.HasSuffix(urlpath, ".svg") {
				contentType = "image/svg+xml"
			} else {
				contentType = "text/plain"
			}

			w.Header().Add("Content-Type", contentType)
			//log.Println(w.Header())
			// Write file content to responder
			w.Write(data)
		}
	} else if match, _ := regexp.MatchString("index", r.URL.Path); match {
		if r.Method == "GET" {
			//fmt.Println("-----------get", r.Form)
			//sayhelloName(w, r)
			t, err := template.ParseFiles(path.Join(CurrentPath, HtmlTemplate_Folder, "listitem.gtpl"))
			//t, err := template.ParseFiles(path.Join(HtmlTemplate_Folder, "index.html"))
			if err != nil {
				log.Println(err)
			}

			log.Println(t.Execute(w, nil))
		} else if r.Method == "POST" {
			var (
				pagecount    int = 1
				commentcount int = 1
				votecount        = 1
				locals           = make(map[string]interface{})
			)
			r.ParseForm()
			//fmt.Println("-----------post", r.Form)
			if len(r.Form["pagecount"][0]) != 0 {
				pagecount, _ = strconv.Atoi(r.Form["pagecount"][0])
			}
			if len(r.Form["commentcount"][0]) != 0 {
				commentcount, _ = strconv.Atoi(r.Form["commentcount"][0])
			}
			if len(r.Form["votecount"][0]) != 0 {
				votecount, _ = strconv.Atoi(r.Form["votecount"][0])
				if votecount > 100 || votecount < 1 {
					votecount = 3
				}
			}
			fetchhtml.HandelAllUrl(pagecount, commentcount, votecount)
			//fmt.Printf("%#v\n", fetchhtml.ItemList)
			locals["itemlist"] = fetchhtml.ItemList
			t, err := template.ParseFiles(path.Join(HtmlTemplate_Folder, "listitem.gtpl"))
			if err != nil {
				log.Println(err)
			}

			t.Execute(w, locals)
		}
	}

}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Printf("%#v\n", r)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func test(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(path.Join(HtmlTemplate_Folder, "index.html"))
	if err != nil {
		log.Println(err)
	}
	log.Println(t.Execute(w, nil))
}
func main() {
	myroute := &MyRoute{}
	err := http.ListenAndServe(":9090", myroute) //设置监听的端口

	/*
		//http.HandleFunc("/login", sayhelloName)      //设置访问的路由
		//http.Handle("static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
		http.HandleFunc("/index", test)
		//fs := http.FileServer(http.Dir("static/"))
		//http.Handle("/", fs)
		http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
		//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
		err := http.ListenAndServe(":9090", nil)
	*/

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
