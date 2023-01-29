package taskMaker

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"path/filepath"
	"strconv"
	"sync/atomic"
	"testSpeedPoolChrome/internal/config"
	"time"
)

func WorkerMaker(numTrade int, cfg *config.Config) {

	func() {

		pathChrome := filepath.Join(cfg.Path, strconv.Itoa(numTrade))

		//		pathChrome := cfg.Path + strconv.Itoa(numTrade)

		fmt.Println(`++++====++++`)
		fmt.Println(pathChrome)
		fmt.Println(`++++====++++`)

		opts := append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("user-data-dir", pathChrome),
			//	chromedp.Flag("no-sandbox", true),
			chromedp.Flag("no-first-run", true),
			chromedp.Flag("headless", !cfg.Debug),
			//	chromedp.Flag("headless", false),
			chromedp.Flag("disable-gpu", false),
			chromedp.Flag("enable-automation", false),
			chromedp.Flag("disable-extensions", false),
			//	chromedp.WindowSize(912, 1368),
			chromedp.WindowSize(1366, 768),
		)

		// create context.allocCtx
		allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
		defer cancel()

		// create context.ctx
		ctx, cancel2 := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
		defer cancel2()

		/*		time.AfterFunc(120*time.Second, func() {
				fmt.Println("Завершаем зависший процесс")
				cancel2()
			})*/

		err := chromedp.Run(ctx,
			chromedp.Navigate("https://google.com"),
		)

		if err != nil {
			fmt.Println(err)
		}

		for {

			if !cfg.TestComplite {

				func() {

					t1C, errC := <-cfg.TaskSC
					if !errC {
						fmt.Println(`Ожидаем завершения...`)
						cfg.STRMutex.Lock()
						cfg.TestComplite = true
						cfg.STRMutex.Unlock()
						return
					} else {

						var buf []byte
						//	fmt.Println(t1C)
						err = chromedp.Run(ctx,
							chromedp.Navigate(t1C),
							chromedp.Screenshot(`body`, &buf, chromedp.ByQuery),
						)
						if err != nil {
							fmt.Println(err)
							return
						}
						fmt.Println(`Отработал ` + strconv.Itoa(numTrade) + " URL=" + t1C)

						err = scrineWriter(buf, cfg)

						buf = nil

						time.Sleep(100 * time.Millisecond)

					}

				}()

			} else {
				atomic.AddUint32(&cfg.TFinish, 1)
				cancel2()
				break
			}
		}

		//runWorker(cfg),

		return

	}()

}
