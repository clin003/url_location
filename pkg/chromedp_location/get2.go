package chromedp_location

// import (
// 	"context"
// 	"fmt"

// 	"net"
// 	"time"

// 	"github.com/chromedp/chromedp"
// )

// func GetUrlRedirect(urlStr string, sleepTime time.Duration) (string, error) {
// 	timeCtx, cancel := context.WithTimeout(GetChromeCtx(false), 30*time.Second)
// 	defer cancel()

// 	var urlLocation string
// 	err := chromedp.Run(timeCtx,
// 		chromedp.Navigate(urlStr),
// 		// chromedp.WaitVisible(`//ul[@class="news-list"]`),
// 		chromedp.Sleep(sleepTime),
// 		chromedp.Location(&urlLocation),
// 	)
// 	if err != nil {
// 		err = fmt.Errorf("打开网址失败(%s):%w", urlStr, err)
// 		// log.Println("打开网址失败1：", urlStr, err.Error())
// 		return "", err
// 	}
// 	// log.Println("正在跳转Url：", urlLocation)
// 	return urlLocation, nil
// }

// //检查是否有9222端口，来判断是否运行在linux上
// func checkChromePort() bool {
// 	addr := net.JoinHostPort("", "9222")
// 	conn, err := net.DialTimeout("tcp", addr, 1*time.Second)
// 	if err != nil {
// 		return false
// 	}
// 	defer conn.Close()
// 	return true
// }

// // ChromeCtx 使用一个实例
// var ChromeCtx context.Context

// func GetChromeCtx(focus bool) context.Context {
// 	if ChromeCtx == nil || focus {
// 		allocOpts := chromedp.DefaultExecAllocatorOptions[:]
// 		allocOpts = append(allocOpts,
// 			chromedp.DisableGPU,
// 			chromedp.Flag("blink-settings", "imagesEnabled=false"), // 禁用图片加载
// 			chromedp.UserAgent(`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36`),
// 			chromedp.Flag("accept-language", `zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6`),
// 		)

// 		if checkChromePort() {
// 			// 不知道为何，不能直接使用 NewExecAllocator ，因此增加 使用 ws://127.0.0.1:9222/ 来调用
// 			c, _ := chromedp.NewRemoteAllocator(context.Background(), "ws://127.0.0.1:9222/")
// 			ChromeCtx, _ = chromedp.NewContext(c)
// 		} else {
// 			c, _ := chromedp.NewExecAllocator(context.Background(), allocOpts...)
// 			ChromeCtx, _ = chromedp.NewContext(c)
// 		}
// 	}

// 	return ChromeCtx
// }
