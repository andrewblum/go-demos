package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://golang.org",
		"http://reddit.com",
		"http://pokeapi.co/",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//fmt.Println(<-c)

	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// we will wait forever if we do this
	// fmt.Println(<-c)

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// for {
	// 	go checkLink2(<-c, c)
	// }

	// this is the same as the above
	// watch this channel, as soon as a value comes out
	// assign the value to l and interate the loop one time
	for l := range c {
		// this will sleep the MAIN routine
		// do not do this
		// time.Sleep(5 * time.Second)

		// using a functiona literal (a lambda function)
		// never reference the same variable in 2+ routines!
		// go is pass by value for a reason
		// this function closes over l and l changes over time
		// resulting in wonky behavior
		// provide l as an argument to the function literal
		go func() {
			time.Sleep(5 * time.Second)
			checkLink2(l, c)
		}()
	}

}

func checkLink2(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}
	fmt.Println(link, "Is ok")
	c <- link
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		//fmt.Println(link, "might be down")
		c <- link + " is down"
		return
	}
	c <- link + " is up"
	//fmt.Println(link, "Is ok")
}
