package main

import "fmt"
import "strings"
func FindCmd(head *tDataNode,cmd string) *tDataNode{
	p := head
	if p==nil || cmd==""{
		return nil
	}
	for p!=nil{
		if strings.Compare(p.cmd,cmd)==0{
			return p
		}
		p = p.next
	}
	return nil
}
func ShowAllCmds(head *tDataNode) int{
	fmt.Println("Menu List:")
	p := head
	for p!=nil{
		fmt.Printf("%s - %s\n", p.cmd, p.desc)
        p = p.next
	}
	return 0
}
type tDataNode struct{
	cmd string
	desc string
	handler func() int
	next *tDataNode
}
var head [2]tDataNode= [2]tDataNode{
	tDataNode{"help", "this is help cmd!", nil,nil},
	tDataNode{"version", "menu program v1.0", nil, nil}}
func Help() int{
	ShowAllCmds(&head[0])
	return 0
}
func main(){
	head[0].next = &head[1]
	head[0].handler = Help
	for true{
		var cmd string;
        fmt.Printf("Input a cmd number > ")
        fmt.Scan(&cmd)
        p := FindCmd(&head[0], cmd)
        if p == nil{
            fmt.Printf("This is a wrong cmd!\n ")
            continue
        }
        fmt.Printf("%s - %s\n", p.cmd, p.desc)
        if p.handler != nil{ 
            p.handler()
        }
	}
}
