package framework

import (
    "log"
    "net/http"
)

// 框架核心结构
type Core struct {
    router map[string]ControllerHandler
}

// 初始化框架核心结构
func NewCore() *Core {
    return &Core{router: map[string]ControllerHandler{}}
}

func (c *Core) Get(url string, handler ControllerHandler) {
    c.router[url] = handler
}

// 框架核心结构实现Handler接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
    log.Println("请求进入")
    ctx := NewContext(request, response)

    router := c.router["foo"]
    if router == nil {
        return
    }
    log.Println("请求结束")

    router(ctx)
}
