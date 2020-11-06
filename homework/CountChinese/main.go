package main

import "fmt"

func main(){
	s:="Hello沙河小王子"
	slen:=len(s)
	flag:=0
	for i:=0, c:= range s{
		if((s[i]>='a'&&s[i]<='z')||(s[i]>='A'&&s[i])<='Z'){
			
		}
		else{
			flag++
		}
	}
}