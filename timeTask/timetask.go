package timeTask

import (
	"time"
)

type JobFunc = func()

func AddTimeTask(duration time.Duration, times int64, myfunc JobFunc) {
	ticker := time.NewTicker(duration * time.Duration(times))
	for {
		<-ticker.C
		myfunc()
	}
}

func AddTimeAsynTask(duration time.Duration, times int64, myfunc JobFunc) {
	go func() {
		AddTimeTask(duration, times, myfunc)
	}()
}
