package main

import (
	"fmt"
	"time"
)

func Ticker() {
	tekrar := time.NewTicker(500 * time.Millisecond) // her yarım saniyede 1
	bitti := make(chan bool)
	go func() {
		for {
			select {
			case <-bitti:
				return
			case zaman := <-tekrar.C:
				fmt.Println("Tekrar zamanı:", zaman)
			}
		}
	}()
	time.Sleep(1600 * time.Millisecond) // 1,6 saniye programı uyut

	//tekrar.Stop()                       // Durdurduk
	//bitti <- true // for döngüsünü sonlandırdık.
	fmt.Println("Tekrarlayıcı durdu!")
}
