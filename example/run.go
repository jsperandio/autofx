package example

import (
	"fmt"
	"time"
)

func Run(s Something) {
	ticker := time.NewTicker(time.Second * 2)
	exit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println(s.DoSomething("lala"))
				fmt.Println(s.DoSomethingElse(2))
			case <-exit:
				ticker.Stop()
				return
			}
		}
	}()
}
