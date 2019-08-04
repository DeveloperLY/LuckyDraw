package main

import (
	"fmt"
	"log"
	logger "luckydraw/log"
)

func main() {
	parse()
	logger.InitLogConfig()
	InitAwardPool()

	Start()
}

func drawPrize(awardBatchs []AwardBatch) {
	awardBatch := WinPrize("DeveloperLY")

	if awardBatch == nil {
		fmt.Println("很遗憾，您未能抽中奖品")
	}

	log.Println("恭喜您抽中:", awardBatch.GetName())
}