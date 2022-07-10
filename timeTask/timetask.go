package timeTask

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

func AddTimeAsynTask(duration time.Duration, myfunc JobFunc) {
	go func() {
		AddTimeTask(duration, myfunc)
	}()
}
