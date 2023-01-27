package makerQueue

import (
	"fmt"
	"strconv"
	"testSpeedPoolChrome/internal/config"
	"testSpeedPoolChrome/internal/taskMaker"
)

func Cirkle(numTrade int, cfg *config.Config) {
	fmt.Println(`Запускаю поток` + strconv.Itoa(numTrade))

	taskMaker.WorkerMaker(numTrade, cfg)

	return
}
