url_location


//远程调试地址

CHROMEDP_REMOTE_DEBUG_URL

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