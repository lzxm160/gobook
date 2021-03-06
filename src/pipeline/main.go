package main

import "fmt"
import "links"
import (
	"log" 
	"crypto/sha256"	
	"io/ioutil" 
	"os" 
	"path/filepath"
	"flag"
)
//go get godoc.org/golang.org/x/net/html
var tokens=make(chan struct{},200)
func crawl(url string)[]string{
	fmt.Println(url)
	tokens<-struct{}{}
	list,err:=links.Extract(url)
	<-tokens
	if err!=nil{
		log.Print(err)
	}
	return list
}
func startCrawl() {
	worklist:=make(chan[]string)
	var n int
	n++
	ll:=[]string{"https://baidu.com","http://gopl.io"}
	go func(){worklist<-ll}()
	seen:=make(map[string]bool)
	for ;n>0;n--{
		list:=<-worklist
		for _,link:=range list{
			if!seen[link]{
				seen[link]=true
				n++
				go func(link string) {
					worklist<-crawl(link)
				}(link)
			}

		}
	}
}
func test() {
	type currency int
	const(
		USD currency=iota
		EUR
		GBP
		RMB
		)
	symbol:=[]string{USD:"$",EUR:"￡"}

	fmt.Println(EUR,symbol[EUR])
	a:=[2]int{1,2}
	b:=[...]int{1,2}
	c:=[2]int{1,3}
	d:=[]int{1,2}
	fmt.Println(a==b,a==c,b==c)
	c1:=sha256.Sum256([]byte("x"))
	c2:=sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n",c1,c2,c1==c2,c1)
	fmt.Printf("%T %T %T\n",a,b,d)
	s:=[]int{0,1,2,3,4,5}
	fmt.Println(s[:2])
	slice1:=make([]int,5)
	slice2:=make([]int,5,5)
	fmt.Printf("%T %d %d\n",slice1,len(slice1),cap(slice1))
	fmt.Printf("%T %d %d\n",slice2,len(slice2),cap(slice2))
}
func test2() {
	var runes []rune
	for _,r:=range "hello,世界"{
		runes=append(runes,r)
	}
	fmt.Printf("%q\n",runes)
}

func main() {
	// start_du()
	// test3()
	// test2()
	// startCrawl()
	// naturals:=make(chan int)
	// squares:=make(chan int)
	// go counter(naturals)
	// go squarer(squares,naturals)
	// printer(squares)
	// ret:=mirroredQuery()
	// fmt.Println(ret)
	// go func(){
	// 	for x:=0;x<5;x++{
	// 		naturals<-x
	// 	}
	// 	close(naturals)
	// }()
	// go func(){
	// 	for x:=range naturals{
	// 		squares<-x*x
	// 	}
	// 	close(squares)
	// }()
	// for x:=range squares{
	// 	fmt.Println(x)
	// }
	// fmt.Println("vim-go")
}
