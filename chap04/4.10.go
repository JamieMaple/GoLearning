package main

import (
	"fmt"
	"github.com/jamie/goland/github"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	var lessOneHour, lessOneMonth, lessOneYear []*github.Issue
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	y0, m0, d0 := time.Now().Date()
	for _, item := range result.Items {
		y, m, d := item.CreatedAt.Date()
		switch {
		case y0-y > 0:
			lessOneYear = append(lessOneYear, item)
		case m != m0:
			lessOneMonth = append(lessOneMonth, item)
		case d0-d >= 0:
			lessOneHour = append(lessOneHour, item)
		}
	}
	if len(lessOneYear) > 0 {
		fmt.Printf("less one hour:\n")
		for _, item := range lessOneHour {
			fmt.Printf("\t#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
	if len(lessOneMonth) > 0 {
		fmt.Printf("less one month:\n")
		for _, item := range lessOneMonth {
			fmt.Printf("\t#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
	if len(lessOneHour) > 0 {
		fmt.Printf("less one hour:\n")
		for _, item := range lessOneHour {
			fmt.Printf("\t#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}
