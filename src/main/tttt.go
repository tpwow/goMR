package main

import "fmt"
import "container/list"

func reduce() string {
    var l *list.List = list.New()
    l.PushBack("fuck")
    l.PushBack("you")
    for e := l.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
    }
    return "haha"
}

func main(){
    reduce()
}
