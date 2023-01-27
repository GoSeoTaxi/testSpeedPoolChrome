package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"testSpeedPoolChrome/internal/config"
	"testSpeedPoolChrome/internal/makerQueue"
	"time"
)

func main() {

	fmt.Println("Starting app...")
	cfg, err := config.InitConfigsWithSendData()
	if err != nil {
		fmt.Printf("No load config %v", err)
	}

	// Количество потоков
	iConf, err := strconv.Atoi(cfg.StartCopy)
	if err != nil {
		fmt.Printf("No load config %v", err)
		iConf = 1
	}

	for i := 1; i <= iConf; i++ {
		go makerQueue.Cirkle(i, cfg)
		time.Sleep(2 * time.Second)
	}

	for {
		//		t1C := <-cfg.TaskSC
		//		fmt.Println(t1C)

		time.Sleep(1 * time.Second)
		//	fmt.Println(`Я работаю, не выключай меня`)
		/*		fmt.Println(cfg.TestComplite)
				fmt.Println(cfg.StartCopy)
				fmt.Println(strconv.Itoa(int(cfg.TFinish)))*/

		if cfg.TestComplite == true && cfg.StartCopy == strconv.Itoa(int(cfg.TFinish)) {
			fmt.Println(`считаем сколько будет файлов`)

			var countF int
			path := cfg.PathFilesToWrite
			extension := regexp.MustCompile(`.png$`)
			err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
				if extension.MatchString(path) {
					countF++
				}
				return nil
			})
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Всего открытых страниц: ", countF)

			break
		}

	}
	// Конец

}
