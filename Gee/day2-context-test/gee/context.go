package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{} //构建json数据时用
//比如可以"username":(string类型)，也可以"age":(int类型)

type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key) //获得输入的json数据
}

//Query获得/hello?后面的name=后的字段
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

//用于显示网页状态
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

//（Header代表HTTP头域的键值对。）
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

//string调用
func (c *Context) String(code int, format string, values ...interface{} /*指不定是什么数据类型*/) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer) //转成json格式
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	//code是http.StatusOK,html是html代码
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
