package main

import (
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/config/cmd"
	"github.com/tddey01/luffy/day017/demo/src/share/config"
	"github.com/tddey01/luffy/day017/demo/src/share/utils/path"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// 定义网关请求
func main() {
	//
	max := http.NewServeMux()
	//
	max.HandleFunc("/", handlerRPC)
	log.Println("Lsten on:8888")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("启动失败", err)
	}

}

func handlerRPC(w http.ResponseWriter, r *http.Request) {
	log.Println("hadlerRPC...")
	// 1 正常请求
	if r.URL.Path == "/" {
		_, err := w.Write([]byte("server ..."))
		if err != nil {
			fmt.Println(err, "server")
		}
		return
	}
	// 2 RPC请求 跨域请求
	if origin := r.Header.Get("Origin"); true {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,AccessToken,Authorization,Token,Origin,X-Token,X-Client,Accept-Encoding, X-Requested-With, If-Modified-Since, Pragma, Last-Modified, Cache-Control, Expires, Content-Type, X-E4M-With")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	handleJSONRPC(w, r)
	return
}

func handleJSONRPC(w http.ResponseWriter, r *http.Request) {
	service, method := path.PathToReceiver(config.Namespace, r.URL.Path)
	log.Println("server " + service)
	log.Println("method" + method)
	// 去取请求体
	br, _ := ioutil.ReadAll(r.Body)
	request := json.RawMessage(br)
	var response json.RawMessage
	req := (*cmd.DefaultOptions().Client).NewJsonRewquest(service, method, &request)
	ctx := path.RequestToContext(r)
	err := (*cmd.DefaultOptions().Client).Call(ctx, req, &request)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, _ := response.MarshalJSON()
	// 设置响应头
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	_, err = w.Write(b)
	if err != nil {
		fmt.Println(err)
	}

}
