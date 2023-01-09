package main

import (
	"fmt"
	"log"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},

	//"linear algebra": {"calculus"},
}

func main() {
	sorted, err := topoSort(prereqs)
	if err != nil {
		log.Println(err)
	}
	for i, course := range sorted {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
	//for i, course := range topoSortMap(prereqs) {
	//	fmt.Printf("%d:\t%s\n", i+1, course)
	//}
}

//ex 5.11
func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(item []string) error
	visitAll = func(items []string) error {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				err := visitAll(m[item])
				if err != nil {
					return err
				}
				order = append(order, item)
			} else {
				hasCycle := true
				for _, s := range order {
					if s == item {
						hasCycle = false
					}
				}
				if hasCycle {
					return fmt.Errorf("has cycle: %s", item)
				}
			}
		}
		return nil
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	err := visitAll(keys)
	if err != nil {
		return nil, err
	}
	return order, nil
}

//ex5.10
func topoSortMap(m map[string][]string) map[int]string {
	order := make(map[int]string)
	seen := make(map[string]bool)
	var visitAll func(item []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order[len(order)] = item
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
