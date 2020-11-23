package main

var goRoutineNum int = 1










func main() {

	writeChanOk := make(chan<- bool, goRoutineNum)

	for i := 1; i <= goRoutineNum; i++ {
		go 
	}
}
