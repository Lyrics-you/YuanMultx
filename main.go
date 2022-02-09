package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"yuanmultx/serverremote"
)

func OpenConfig(configname string) string {
	var content strings.Builder
	//打开文件
	file, err := os.Open(fmt.Sprintf("./config/%s.conf", configname))
	if err != nil {
		log.Printf("file %s.conf open failed! err:%v", configname, err)
	}
	//及时关闭 file 句柄，否则会有内存泄漏
	defer file.Close()
	//创建一个 *Reader ， 是带缓冲的
	reader := bufio.NewReader(file)
	isAdd := false
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束
		content.WriteString(str)

		if !isAdd {
			regexp, _ := regexp.Compile(".*?server_remote.*?")
			if len(regexp.FindString(str)) != 0 {
				content.WriteString(serverremote.ServerRemote())
				content.WriteString("\n")
				isAdd = true
				log.Printf("Add [server_remote] success! :%s", serverremote.ServerRemote())
			}
		}
		if err == io.EOF { //io.EOF 表示文件的末尾
			break
		}
	}
	log.Printf("Load %s success!", configname)
	return content.String()
}
func GetConfig(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()       // 解析参数，默认是不会解析的
	fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path : ", r.URL.Path)
	// fmt.Println("scheme", r.URL.Scheme)
	var configname string
	if len(r.Form) == 0 || len(r.Form["name"][0]) == 0 {
		configname = "Sabrina"
	} else {
		configname = r.Form["name"][0]
	}

	content := OpenConfig(configname)
	fmt.Fprint(w, content) // 这个写入到 w 的是输出到客户端的
	log.Printf("Get %s success!", configname)
}
func main() {
	http.HandleFunc("/getconfig", GetConfig)  // 设置访问的路由
	err := http.ListenAndServe(":39790", nil) // 设置监听的端口
	if err != nil {
		log.Println("ListenAndServe: ", err)
	}
}
