package main

import (
	"Go-by-myself/ch4/github"
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

const (
	YEAR     = 31536000
	MONTH    = 2592000
	NOLIMITE = math.MaxInt64
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		if time.Now().Unix()-item.CreatedAt.Unix() <= YEAR {
			fmt.Printf("%s\t#%-5d\t%9.9s\t%.55s\n",
				item.CreatedAt.Format(time.ANSIC), item.Number, item.User.Login, item.Title)
		}

	}
}
