package main

var goRoutineNum int = 8

func findPrimeNum(intChan chan int, primeChan chan int, exitChan chan<- bool) {





}

func main() {

	intChan := make(chan int, 1000)
	primeChan := make(chan int, goRoutineNum*1100)
	exitChan := make(chan<- bool, goRoutineNum)

	for i := 1; i <= goRoutineNum; i++ {

	}
	for {
		if len(exitChan) == goRoutineNum {
			close(exitChan)
			close(primeChan)
			break
		}
	}

}
