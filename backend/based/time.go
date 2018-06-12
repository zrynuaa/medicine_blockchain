package based

import (
	"time"
	"fmt"
)

func StartTimer() {
	for {
		now := time.Now()
		fmt.Println(now)
		next := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute()+1, 0, 0, now.Location())
		t := time.NewTimer(next.Sub(now))
		<-t.C
		//fmt.Println("xxxxxxxxxxxxxxxxxxxx")
		go func() {
			timeToDo(uint64(next.Unix()))
		}()
	}
}

func timeToDo(ts uint64){
	nowheight := getLastBlockHeight()
	b := new(Block)
	b.DataHash = getOneMinuteBeforeHash(ts)
	b.PrevHash = getBlockHash(nowheight)
	b.Ts = ts
	b.Height = nowheight+1
	putBlock(*b)
}
