package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

// NewServer 创建一个新的negroni服务器，使用mux作为路由
func NewServer() *negroni.Negroni {
	// 使用formatter格式化输出JSON数据
	formatter := render.New(render.Options{IndentJSON: true})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

// 初始化路由
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/hello/{id}", helloHandler(formatter)).Methods(http.MethodGet)
}

// 对于URL"/hello/{id}"的路由
func helloHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]
		// 格式化输出JSON
		formatter.JSON(w, http.StatusOK, struct {
			Hello string `json:"hello"`
		}{id})
	}
}
