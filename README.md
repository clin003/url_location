url_location


//远程调试地址 ws://127.0.0.1:9222 http://127.0.0.1:9222

CHROMEDP_REMOTE_DEBUG_URL

//远程调试地址 端口 9222

CHROMEDP_REMOTE_DEBUG_PORT

//支持跳转的域名列表 

CHROMEDP_CAN_WAIT_DOMAIN_LIST

接口调用例子
```
package main

import (
   "fmt"
   "strings"
   "net/http"
   "io/ioutil"
)

func main() {

   url := "http://127.0.0.1:8080/api/v1/urlOriginAndLocation"
   method := "POST"

   payload := strings.NewReader(`{
    "url": "https://s.click.taobao.com/xxx",
    "retJson": false
}`)

   client := &http.Client {
   }
   req, err := http.NewRequest(method, url, payload)

   if err != nil {
      fmt.Println(err)
      return
   }
   req.Header.Add("Content-Type", "application/json")

   res, err := client.Do(req)
   if err != nil {
      fmt.Println(err)
      return
   }
   defer res.Body.Close()

   body, err := ioutil.ReadAll(res.Body)
   if err != nil {
      fmt.Println(err)
      return
   }
   fmt.Println(string(body))
}
```

## 交流群
[@baicai_dev](https://t.me/baicai_dev)
<!--
## 赞赏
[@一杯咖啡](https://3ae.cn/donations/)
![赞赏白菜林，多少不重要，1元也是支持](/zanalipay.jpg)  ![赞赏白菜林，多少不重要，1元也是支持](/zanweixin.jpg)
-->