package main

import (
    "fmt"
    "sort"
)

var courses = map[string][]string {
    "algorithms":            {"data structures"},
    "calculus":              {"linear algebra"},
    "compilers":             {"data structures", "formal languages", "computer organization"},
    "data structures":       {"discrete math"},
    "databases":             {"data structures"},
    "discrete math":         {"intro to programming"},
    "formal languages":      {"discrete math"},
    "networks":              {"operating systems"},
    "operating systems":     {"data structures", "computer organization"},
    "programming languages": {"data structures", "computer organization"},
}

type DependencyList struct {
    Mapping map[string][]string
    sortedKeys []string
}

func (d DependencyList) Len() int {
    return len(d.Mapping)
}

func (d DependencyList) Less(i, j int) bool {
    keys := d.SortedKeys()
    return len(d.Mapping[keys[i]]) < len(d.Mapping[keys[j]])
}

func (d DependencyList) Swap(i, j int) {
    keys := d.SortedKeys()
    keys[i], keys[j] = keys[j], keys[i]
}

func (d DependencyList) SortedKeys() []string {
    var keys []string
    if d.sortedKeys == nil {
        for key := range d.Mapping {
            keys = append(keys, key)
        }
        d.sortedKeys = keys
    }
    keys = d.sortedKeys
    return keys
}

func main() {
    for i, course := range topoSort(courses) {
        fmt.Printf("%d:\t%s\n", i, course)
    }
}

func topoSort(m map[string][]string) []string {
    var order []string
    seen := make(map[string]bool)
    var visitAll func(items []string)

    visitAll = func(items []string) {
        for _, item := range items {
            if !seen[item] {
                seen[item] = true
                visitAll(m[item])
                order = append(order, item)
            }
        }
    }

    deps := DependencyList{Mapping: m}
    sort.Sort(deps)
    visitAll(deps.SortedKeys())
    return order
}