package taskMaker

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"testSpeedPoolChrome/internal/config"
	"testSpeedPoolChrome/internal/constants"
	"time"
)

func scrineWriter(b []byte, cfg *config.Config) (err error) {

	//cfg.PathFiles + "\\" + "screenshot" + strconv.Itoa(time.Now().Nanosecond()) + ".png"

	//	pathF := filepath.Join(os.Getenv("HOMEDRIVE")+os.Getenv("HOMEPATH"), "AppData", "Local", "Yandex", "YandexBrowser", "Application", "browser.exe")

	tr, err := strconv.Atoi(cfg.StartCopy)
	if err != nil {
		tr = 999999
	}
	rand.Seed(int64(time.Now().Nanosecond()))

	pathF := filepath.Join(cfg.PathFilesToWrite, constants.NameScreens+strconv.Itoa(rand.Intn(100000*tr))+"_"+strconv.Itoa(time.Now().Nanosecond())+".png")

	cfg.STRMutex.Lock()
	err = os.WriteFile(pathF, b, 0644)
	cfg.STRMutex.Unlock()

	if err != nil {
		fmt.Println(err)
	}
	return err
}
