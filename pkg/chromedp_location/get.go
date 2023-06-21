package chromedp_location

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"
	"url_location/utils"

	"github.com/chromedp/chromedp"
)

// func GetUrlLocation(navigateUrlStr string, remoteDebuggingUrl, remoteDebuggingPort, userAgent string, sleepSecondN int) (retText string, err error) {
// 	// 禁用chrome headless
// 	opts := append(
// 		chromedp.DefaultExecAllocatorOptions[:],
// 		chromedp.Flag("blink-settings", "imagesEnabled=false"), // 禁用图片加载
// 		// chromedp.Flag("headless", isHeadless),
// 		// chromedp.Flag("enable-automation", false), //

// 		chromedp.DisableGPU,
// 		// chromedp.WindowSize(1024, 768),
// 		// chromedp.UserAgent(userAgent),
// 		chromedp.Flag("remote-debugging-port", remoteDebuggingPort),
// 	)
// 	if len(userAgent) > 0 {
// 		opts = append(opts, chromedp.UserAgent(userAgent))
// 	}
// 	// allocCtx, cancel1 := chromedp.NewExecAllocator(context.Background(), opts...)
// 	allocCtx, cancel1 := chromedp.NewRemoteAllocator(context.Background(), remoteDebuggingUrl)
// 	defer cancel1()
// 	// 初始化chromedp的上下文，后续这个页面都使用这个上下文进行操作
// 	chromedpCtx, cancel2 := chromedp.NewContext(
// 		allocCtx,
// 		// 设置日志方法
// 		chromedp.WithLogf(log.Printf),
// 	)
// 	// 通常可以使用 defer cancel() 去取消
// 	// 但是在Windows环境下，我们希望程序能顺带关闭掉浏览器
// 	// 如果不希望浏览器关闭，使用cancel()方法即可
// 	defer cancel2()
// 	// defer chromedp.Cancel(ctx)

// 	timeOut := time.Second * 10
// 	// 创建新的chromedp上下文对象，超时时间的设置不分先后
// 	// 注意第二个返回的参数是cancel()
// 	timeCtx, cancel3 := context.WithTimeout(chromedpCtx, timeOut)
// 	defer cancel3()

// 	sleepTime := time.Duration(sleepSecondN) * time.Second
// 	var urlLocation string
// 	err = chromedp.Run(timeCtx,
// 		chromedp.Navigate(navigateUrlStr),

// 		chromedp.Sleep(sleepTime),
// 		chromedp.Location(&urlLocation),
// 	)
// 	if err != nil {
// 		err = fmt.Errorf("打开网址失败(%s):%w", navigateUrlStr, err)
// 		return "", err
// 	}

// 	return urlLocation, nil
// }

// func GetUrlLocationByRemoteDebug(navigateUrlStr string, remoteDebuggingUrl string, sleepSecondN int) (retText string, err error) {
// 	allocCtx, cancel1 := chromedp.NewRemoteAllocator(context.Background(), remoteDebuggingUrl)
// 	defer cancel1()
// 	// 初始化chromedp的上下文，后续这个页面都使用这个上下文进行操作
// 	chromedpCtx, cancel2 := chromedp.NewContext(
// 		allocCtx,
// 		// 设置日志方法
// 		chromedp.WithLogf(log.Printf),
// 		chromedp.NoRedirect(true),
// 	)
// 	// 通常可以使用 defer cancel() 去取消
// 	// 但是在Windows环境下，我们希望程序能顺带关闭掉浏览器
// 	// 如果不希望浏览器关闭，使用cancel()方法即可
// 	defer cancel2()
// 	// defer chromedp.Cancel(ctx)

// 	timeOut := time.Second * 10
// 	// 创建新的chromedp上下文对象，超时时间的设置不分先后
// 	// 注意第二个返回的参数是cancel()
// 	timeCtx, cancel3 := context.WithTimeout(chromedpCtx, timeOut)
// 	defer cancel3()

// 	// sleepTime := time.Duration(sleepSecondN) * time.Second
// 	var urlLocation string
// 	err = chromedp.Run(timeCtx,
// 		// chromedp.Navigate(navigateUrlStr),
// 		task(navigateUrlStr),
// 		// chromedp.Sleep(sleepTime),
// 		chromedp.Location(&urlLocation),
// 	)
// 	if err != nil {
// 		err = fmt.Errorf("打开网址失败(%s):%w", navigateUrlStr, err)
// 		return "", err
// 	}

// 	return urlLocation, nil
// }

func getChromedpRemoteDebugCtx(remoteDebuggingUrl, remoteDebuggingPort string) (context.Context, context.CancelFunc) {
	allocCtx, _ := chromedp.NewRemoteAllocator(context.Background(), remoteDebuggingUrl)
	// defer cancel()

	// 初始化chromedp的上下文，后续这个页面都使用这个上下文进行操作
	chromedpCtx, _ := chromedp.NewContext(
		allocCtx,
		// 设置日志方法
		chromedp.WithLogf(log.Printf),
		// chromedp.NoRedirect(true),
	)
	// 通常可以使用 defer cancel() 去取消
	// 但是在Windows环境下，我们希望程序能顺带关闭掉浏览器
	// 如果不希望浏览器关闭，使用cancel()方法即可
	// defer cancel()

	// defer chromedp.Cancel(ctx)

	timeOut := time.Second * 10
	// 创建新的chromedp上下文对象，超时时间的设置不分先后
	// 注意第二个返回的参数是cancel()
	// timeCtx, cancel := context.WithTimeout(chromedpCtx, timeOut)
	// defer cancel()
	return context.WithTimeout(chromedpCtx, timeOut)
}
func getChromedpDebugCtx(remoteDebuggingUrl, remoteDebuggingPort string) (context.Context, context.CancelFunc) {
	// 禁用chrome headless
	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("blink-settings", "imagesEnabled=false"), // 禁用图片加载
		// chromedp.Flag("headless", isHeadless),
		// chromedp.Flag("enable-automation", false), //

		chromedp.DisableGPU,
		// chromedp.WindowSize(1024, 768),
		// chromedp.UserAgent(userAgent),
		chromedp.Flag("remote-debugging-port", remoteDebuggingPort),
	)
	// if len(userAgent) > 0 {
	// 	opts = append(opts, chromedp.UserAgent(userAgent))
	// }
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	// allocCtx, cancel1 := chromedp.NewRemoteAllocator(context.Background(), remoteDebuggingUrl)
	// defer cancel()
	// 初始化chromedp的上下文，后续这个页面都使用这个上下文进行操作
	chromedpCtx, _ := chromedp.NewContext(
		allocCtx,
		// 设置日志方法
		chromedp.WithLogf(log.Printf),
	)
	// 通常可以使用 defer cancel() 去取消
	// 但是在Windows环境下，我们希望程序能顺带关闭掉浏览器
	// 如果不希望浏览器关闭，使用cancel()方法即可
	// defer cancel()
	// defer chromedp.Cancel(ctx)

	timeOut := time.Second * 30
	// 创建新的chromedp上下文对象，超时时间的设置不分先后
	// 注意第二个返回的参数是cancel()
	// timeCtx, cancel := context.WithTimeout(chromedpCtx, timeOut)
	// defer cancel()

	return context.WithTimeout(chromedpCtx, timeOut)
}
func GetUrlLocation(navigateUrlStr string, remoteDebuggingUrl, remoteDebuggingPort string, sleepSecondN int) (string, error) {
	if retText, err := getUrlLocationByCtxFunc(navigateUrlStr, remoteDebuggingUrl, remoteDebuggingPort, sleepSecondN, getChromedpRemoteDebugCtx); err != nil {
		err = fmt.Errorf("GetUrlLocation.getUrlLocationByCtxFunc.getChromedpRemoteDebugCtx:%w", err)
		fmt.Println(err)
		if retText, err1 := getUrlLocationByCtxFunc(navigateUrlStr, remoteDebuggingUrl, remoteDebuggingPort, sleepSecondN, getChromedpDebugCtx); err1 != nil {
			err1 = fmt.Errorf("GetUrlLocation.getUrlLocationByCtxFunc.getChromedpDebugCtx:%w", err1)
			fmt.Println(err1)
			return "", fmt.Errorf("%w\n%w", err, err1)
		} else {
			return retText, nil
		}
	} else {
		return retText, nil
	}

}
func getUrlLocationByCtxFunc(
	navigateUrlStr string, remoteDebuggingUrl, remoteDebuggingPort string, sleepSecondN int,
	ctxFun func(remoteDebuggingUrl, remoteDebuggingPort string) (context.Context, context.CancelFunc),
) (retOriginUrl string, err error) {
	// allocCtx, cancel := chromedp.NewRemoteAllocator(context.Background(), remoteDebuggingUrl)
	// defer cancel()

	// // 初始化chromedp的上下文，后续这个页面都使用这个上下文进行操作
	// chromedpCtx, cancel := chromedp.NewContext(
	// 	allocCtx,
	// 	// 设置日志方法
	// 	chromedp.WithLogf(log.Printf),
	// 	// chromedp.NoRedirect(true),
	// )
	// // 通常可以使用 defer cancel() 去取消
	// // 但是在Windows环境下，我们希望程序能顺带关闭掉浏览器
	// // 如果不希望浏览器关闭，使用cancel()方法即可
	// defer cancel()

	// // defer chromedp.Cancel(ctx)

	// timeOut := time.Second * 10
	// // 创建新的chromedp上下文对象，超时时间的设置不分先后
	// // 注意第二个返回的参数是cancel()
	// timeCtx, cancel := context.WithTimeout(chromedpCtx, timeOut)
	// defer cancel()

	timeCtx, cancel := ctxFun(remoteDebuggingUrl, remoteDebuggingPort)
	defer cancel()
	// ch := chromedp.WaitNewTarget(timeCtx, func(info *target.Info) bool {
	// 	return info.URL != ""
	// })

	if err = chromedp.Run(timeCtx,
		chromedp.Navigate(navigateUrlStr),
		// task(navigateUrlStr),
		// chromedp.Sleep(sleepTime),
	// chromedp.Location(&urlLocation),
	); err != nil {
		err = fmt.Errorf("打开网址失败(%s):%w", navigateUrlStr, err)
		return "", err
	}
	// newCtx, cancel := chromedp.NewContext(timeCtx, chromedp.WithTargetID(<-ch))
	// defer cancel()
	// fmt.Println("hello(newCtx)")
	// sleepTime := time.Duration(sleepSecondN) * time.Second
	var urlLocation string
	// defer func() {
	// 	fmt.Println("defer")
	// }()

	cUrlLocation := make(chan string)
	defer close(cUrlLocation)
	c := make(chan string)
	defer close(c)
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	go func() { //ctx context.Context
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recover", err)
				return
			}
		}()
		for {
			select {
			case <-timeCtx.Done():
				return
			case s, ok := (<-c):
				if ok {
					fmt.Println("received", s)
					if utils.UnionMsgUrlCheck(s) {
						cUrlLocation <- s
						return
					}
					fmt.Println(ok, time.Now().String())
				} else {
					fmt.Println("close(c)")
					return
				}
				// default:
			}
		}
	}() //ctx
	if err = chromedp.Run(timeCtx,
		// chromedp.Navigate(navigateUrlStr),
		task(navigateUrlStr, c),
		// chromedp.Sleep(sleepTime),
		chromedp.Location(&urlLocation),
		// tasks,
	); err != nil {
		err = fmt.Errorf("获取Location失败(%s):%w", navigateUrlStr, err)
		return "", err
	}
	if utils.UnionMsgUrlCheck(urlLocation) {
		return urlLocation, nil
	}
	urlLocation2 := <-cUrlLocation
	fmt.Println("cUrlLocation.urlLocation2", urlLocation2)
	return urlLocation2, nil
}

func task(urlStr string, c chan<- string) chromedp.Tasks {
	return chromedp.Tasks{
		checkUnionUrl(urlStr, c),
	}
}

func checkUnionUrl(urlStr string, c chan<- string) chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		// if err = chromedp.Navigate(urlStr).Do(ctx); err != nil {
		// 	fmt.Println("checkUnionUrl(err0)", err)
		// 	return
		// }
		// var currUrl string
		// if err = chromedp.Evaluate(`window.location.href`, &currUrl).Do(ctx); err != nil {
		// 	fmt.Println("checkUnionUrl(err1)", err)
		// 	return
		// }
		// log.Println("window.location.href: ", currUrl)
		// if unionMsgUrlCheck(currUrl) {
		// 	fmt.Println("unionMsgUrlCheck(currUrl)", currUrl)
		// 	chromedp.Stop()
		// }
		var referrerUrl string
		if err = chromedp.Evaluate(`window.document.referrer`, &referrerUrl).Do(ctx); err != nil {
			fmt.Println("checkUnionUrl(referrerUrl)", err)
			return
		}
		if len(referrerUrl) > 0 && strings.HasPrefix(referrerUrl, "http") {
			c <- referrerUrl
		}
		// s := fmt.Sprintf("window.document.referrer:%s", referrerUrl)
		// c <- s
		// log.Println("window.document.referrer: ", referrerUrl)
		if utils.UnionMsgUrlCheck(referrerUrl) {
			fmt.Println("unionMsgUrlCheck(referrerUrl)", referrerUrl)
			return
		}

		var locationUrl string
		if err = chromedp.Location(&locationUrl).Do(ctx); err != nil {
			fmt.Println("checkUnionUrl(err2)", err)
			return
		}
		// s1 := fmt.Sprintf("chromedp.Location:%s", locationUrl)
		// c <- s1
		// log.Println("chromedp.Location: ", locationUrl)
		if strings.EqualFold(locationUrl, urlStr) {
			return checkUnionUrl(urlStr, c).Do(ctx)
		}
		if len(locationUrl) > 0 && strings.HasPrefix(locationUrl, "http") {
			c <- locationUrl
		}
		if utils.UnionMsgUrlCheck(locationUrl) {
			// fmt.Println("unionMsgUrlCheck(locationUrl)", locationUrl)
			return
		}
		return
	}
}
