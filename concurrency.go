package main

import (
	"fmt"
	"log"
	"time"
)

func concurrecncyExample() {
	apporder := make(chan order, 3)
	inShopOrder := make(chan order, 3)
	queue := make(chan order, 3)

	go func() {
		for i := 0; i < 6; i++ {
			apporder <- order(100 + i)
			time.Sleep(10 * time.Second)

		}
		close(apporder)

	}()

	go func() {
		for i := 0; i < 4; i++ {
			inShopOrder <- order(200 + i)
			time.Sleep(15 * time.Second)
		}
		close(inShopOrder)
	}()

	go partier(apporder, inShopOrder, queue)

	for o := range queue {
		log.Printf("Served %s\n", o)
	}
	log.Println("Done!")
}

func partier(
	appOrders <-chan order,
	inShopOrders <-chan order,
	queue chan<- order,
) {
	shouldClose := false
	closeTimeCh := time.After(1 * time.Minute)
	for !shouldClose {
		select {
		case ord, ok := <-appOrders:
			if ok {
				log.Printf("There is %s coming from appOrders channel\n", ord)
				queue <- ord
			}
		case ord, ok := <-inShopOrders:
			if ok {
				log.Printf("There is %s coming from inShopOrders channel\n", ord)
				queue <- ord
			}
		case now := <-closeTimeCh:
			log.Printf("It is %v now, the shop is closing\n", now)
			shouldClose = true
		default:
			log.Println("There is no order on both channels, I will go cleaning instead")
			doCleaning()
		}
	}
	close(queue)
	log.Println("Shop is closed, I’m going home now. Bye!")
}

func doCleaning() {
	time.Sleep(5 * time.Second)
	log.Println("Partier: Cleaning done")
}

type order int

func (o order) String() string {
	return fmt.Sprintf("order-%02d", o)
}
