package config

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"testSpeedPoolChrome/internal/constants"
	"time"
)

type Config struct {
	StartCopy string `env:"START_COPY"`
	Debug     bool   `env:"SERVER_DEBUG"`
	Path      string `env:"TEMP_PATH"`
	PathFiles string `env:"FILES_PATH"`
	//	URLImport    string `env:"URLS_IMPORT"`
	PathFilesToWrite  string
	StartApp          time.Time
	TaskSC            chan string
	TFinish           uint32
	TestComplite      bool
	STRMutex          sync.Mutex
	TimeDurectionTest string `env:"TIME_DUR"`
}

func initConfig() (*Config, error) {
	var cfg Config

	cfg.TaskSC = make(chan string)

	flag.StringVar(&cfg.TimeDurectionTest, "timeD", strconv.Itoa(constants.TimeDur), "timeD=3600")
	flag.StringVar(&cfg.StartCopy, "t", "1", "t=1")
	flag.BoolVar(&cfg.Debug, "debug", false, "debug=true")
	flag.StringVar(&cfg.Path, "path", constants.TPath, "path=/tmp/Profiles")
	flag.StringVar(&cfg.PathFiles, "pathFiles", constants.PathFiles, "pathFiles=/tmp/folder")
	//	flag.StringVar(&cfg.URLImport, "urlI", constants.UrlI, "urlI=http://avtozzzapchasti.ru/rest/get_items/")

	flag.Parse()

	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	//записываем время старта программы.
	cfg.StartApp = time.Now()

	//проверяем папку для профилей
	if _, err := os.Stat(cfg.Path); os.IsNotExist(err) {
		fmt.Println("Папки для файлов нет, поэтому мы её создаём.")
		err = os.Mkdir(cfg.PathFiles, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	//проверяем папку для картинок
	if _, err := os.Stat(cfg.PathFiles); os.IsNotExist(err) {
		fmt.Println("Папки для файлов нет, поэтому мы её создаём.")
		err = os.Mkdir(cfg.PathFiles, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

	}

	sub1 := filepath.Join(cfg.PathFiles, cfg.StartApp.Format("2006_01_02_15_04_05"))
	if _, err = os.Stat(sub1); os.IsNotExist(err) {
		errC := os.Mkdir(sub1, os.ModePerm)
		if errC != nil {
			log.Fatal(`ERR create folder`)
		}
	}
	cfg.PathFilesToWrite = sub1

	return &cfg, nil
}

func InitConfigsWithSendData() (*Config, error) {

	cfg, err := initConfig()
	if err != nil {
		fmt.Printf("No load config %v", err)
	}

	go func(cfg *Config) {
		t1 := cfg.StartApp
		//Ждём перед стартом
		/*		timeOut, err := strconv.Atoi(cfg.StartCopy)
				if err != nil {
					timeOut = 10
				}
		*/
		timeDur, err := strconv.Atoi(cfg.TimeDurectionTest)
		if err != nil {
			timeDur = constants.TimeDur
		}

		//time.Sleep(time.Duration(timeOut+2) * time.Second)
		for {
			if t1.Add(time.Duration(timeDur) * time.Second).After(time.Now()) {

				rand.Seed(time.Now().UnixNano())
				//cfg.TaskSC <- constants.URLsList[rand.Intn(len(constants.URLsList))]
				cfg.TaskSC <- "https://www.google.com/search?q=" + strconv.Itoa(time.Now().Nanosecond())
			} else {
				close(cfg.TaskSC)
				break
			}
		}
		fmt.Println(`Закрываем канал с задачками`)
	}(cfg)

	return cfg, nil
}
