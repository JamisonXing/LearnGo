package raceProblem

import (
	"fmt"
	"sync"
)

type ad struct {
	Id int
}

// RaceProblem 解决for循环的race问题
func RaceProblem() {
	var sum int
	var ads []ad
	adMap := make(map[int]ad, 0)
	var wg sync.WaitGroup
	mutex := sync.Mutex{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(j int) {
			mutex.Lock()
			defer mutex.Unlock()
			ads = append(ads, ad{j})
			adMap[j] = ad{j}
			sum += j
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(sum)
	fmt.Println(ads)
	fmt.Println(len(ads))
	fmt.Println(adMap)
	fmt.Println(len(adMap))
}
