package main

import (
	"./fcgiclient"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "ok!!!")
		fmt.Println("ok!")

		env := make(map[string]string)
		url := "/index.php"
		env["DOCUMENT_ROOT"] = "D:/APM/APMServ5.2.6/nginx-1.5.9/html"
		env["SCRIPT_FILENAME"] = env["DOCUMENT_ROOT"] + url
		env["SERVER_SOFTWARE"] = "gomws/1.0"
		env["REMOTE_ADDR"] = "127.0.0.1"
		env["SERVER_PROTOCOL"] = "HTTP/1.1"

		reqParams := ""

		if len(reqParams) != 0 {
			env["CONTENT_LENGTH"] = strconv.Itoa(len(reqParams))
			env["REQUEST_METHOD"] = "POST"
			//env["PHP_VALUE"] = "allow_url_include = On\ndisable_functions = \nsafe_mode = Off\nauto_prepend_file = php://input"
		} else {
			env["REQUEST_METHOD"] = "GET"
		}

		fcgi, err := fcgiclient.New("127.0.0.1", 9000)
		if err != nil {
			fmt.Printf("err: %v", err)
		}

		stdout, stderr, err := fcgi.Request(env, reqParams)
		if err != nil {
			fmt.Printf("err: %v", err)
		}

		var cutLine = "-----0vcdb34oju09b8fd-----\n"
		if strings.Contains(string(stdout), cutLine) {
			stdout = []byte(strings.SplitN(string(stdout), cutLine, 2)[0])
		}

		fmt.Printf("%s", stdout)
		fmt.Fprintf(w, string(stdout))
		if len(stderr) > 0 {
			fmt.Printf("%s", stderr)
		}

	}) // 设置访问的路由
	err := http.ListenAndServe(":8000", nil) //设置监听的端口
	if err != nil {
		log.Fatal("listen and server:", err)
	}
}
