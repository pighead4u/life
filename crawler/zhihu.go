package crawler

import (
	"github.com/parnurzeal/gorequest"
	"fmt"
)

func FetchTest() {
	request := gorequest.New()
	response, body, _ := request.Get("http://zhuanlan.zhihu.com/api/columns/jixin").End()

	fmt.Println("%s, %s", response, body)

}
