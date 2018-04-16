package main

// todo

/*

- save data to csv/sql
- set start and stop from command line argument

*/

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"os"
)


func main() {
	for i := 1; i < 1000; i++{
	       resp, err := soup.Get(fmt.Sprintf("https://xkcd.com/%d", i))
	       fmt.Println(fmt.Sprintf("grabbing data from xkcd %d", i))
	          if err != nil {
		      os.Exit(1)
	   }
	doc := soup.HTMLParse(resp)
	links := doc.Find("div", "id", "comic").FindAll("img")
	for _, link := range links {
		fmt.Println(link.Attrs()["src"])
	}
	}
}
