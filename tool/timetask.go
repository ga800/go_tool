package tool

import (
	"time"
)

type JobFunc = func()

func AddTimeTask(duration time.Duration, myfunc JobFunc) {
	ticker := time.NewTicker(duration)
	for {
		<-ticker.C
		myfunc()
	}
}
