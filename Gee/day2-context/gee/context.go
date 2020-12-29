package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 代码最开头，给map[string]interface{}起了一个别名gee.H，构建JSON数据时，显得更简洁。
// Context目前只包含了http.ResponseWriter和*http.Request，
// 另外提供了对 Method 和 Path 这两个常用属性的直接访问。
// 提供了访问Query和PostForm参数的方法。
// 提供了快速构造String/Data/JSON/HTML响应的方法。

type H map[string]interface{} //给map[string]interface{}这个类型起了个别名H
//主要是为了构建JSON数据时能更加简洁

//Context目前只包含http.ResponseWriter和*http.Request
type Context struct {
	//初始对象
	Writer http.ResponseWriter
	Req    *http.Request
	//请求信息
	Path   string //端口号
	Method string //方法GET\POST
	//提供了对Method和Path的直接访问方法
	//参考base3中gee.go  Method 属于req
	//返回信息
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path, //URL路径
		Method: req.Method,
	}
}

// 提供了访问Query和PostForm参数的方法。
// 提供了快速构造String/Data/JSON/HTML响应的方法。

//快速访问FormValue的方法，获取POST的数据，实际上就是Request.FormValue
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

//快速访问Query的方法 相当于返回?后面的内容，
//如http://localhost:9999/hello?name=geektutu
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

//对Context结构体中StatusCode赋值
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)

}

//Set结构体中Writer的Header值（Header代表HTTP头域的键值对。）
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

//c.String实现在网页输出hello %s
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain") //一个key一个value
	//Header=>map[string]string =>Header[Content-Type]=text/plain
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
	//Write以有线格式将头域写入w。
	//Sprintf根据format参数生成格式化的字符串并返回该字符串。
	//Sprintf() 是把格式化字符串输出到指定的字符串中，可以用一个变量来接受，然后在打印
	// Write向连接中写入作为HTTP的一部分回复的数据。
	// 如果被调用时还未调用WriteHeader，本方法会先调用WriteHeader(http.StatusOK)
	// 如果Header中没有"Content-Type"键，
	// 本方法会使用包函数DetectContentType检查数据的前512字节，将返回值作为该键的值。
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	//调用刚才写好的SetHeader 函数将键值对存入头部
	//即存入c.ResponseWrite.Header().Set("Content-Type","application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

//然而并没有用到
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
	// Write向连接中写入作为HTTP的一部分回复的数据。
	// 如果被调用时还未调用WriteHeader，本方法会先调用WriteHeader(http.StatusOK)
	// 如果Header中没有"Content-Type"键，
	// 本方法会使用包函数DetectContentType检查数据的前512字节，将返回值作为该键的值。
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
