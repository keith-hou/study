/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package syntax

import (
	"fmt"
	"time"
)

func AddNumber(c chan<- int, out int) {
	for {
	c <- out
	time.Sleep(1 * time.Second)
	}
}

func Boot() {
	var chan1 = make(chan int, 10)
	var chan2 = make(chan int, 10)
	go AddNumber(chan1, 1)
	go AddNumber(chan2, 2)
	for {
		select {
		case e := <-chan1:
			{
				fmt.Print(fmt.Sprintf("chan1: %d\n", e))
			}
		case e := <-chan2:
			{
				fmt.Print(fmt.Sprintf("chan1: %d\n", e))
			}
		default:
			fmt.Print(fmt.Sprintf("no element \n"))
			time.Sleep(1 * time.Second)
		}
	}
}
