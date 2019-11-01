package main

import (
    "fmt"
    "time"
)

func say(s string){
    for i := 0;i < 5; i++{
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

/* 开启4个线程，同时输出tom1,tom2,tom3,tom4,而tom5，tom6是等线程运行完再逐个输出 */
func main(){
    go say("tom1") //开启第1个say线程
    go say("tom2") //开启第2个say线程
    go say("tom3") //开启第3个say线程
    say("tom4")
    say("tom5")
    say("tom6")
}