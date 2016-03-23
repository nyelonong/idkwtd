/*
$ go run 2010.africa.a.go < A-small-practice.in
Case #1: 2 3
Case #2: 1 4
Case #3: 4 5
Case #4: 29 46
Case #5: 11 56
Case #6: 4 5
Case #7: 40 46
Case #8: 16 35
Case #9: 55 74
Case #10: 7 9
*/

package main

import (
    "fmt"
    "sort"
)

type Product struct {
    A int
    B int
    Sum int
}

type BySum []Product

func (a BySum) Len() int           { return len(a) }
func (a BySum) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySum) Less(i, j int) bool { return a[i].Sum < a[j].Sum }

func main() {
    var c, credit, n int
    fmt.Scanf("%d", &c)

    for index := 1; index <= c; index++ {
        fmt.Scanf("%d", &credit)
        fmt.Scanf("%d", &n)

        prods := make([]int, n)
        result := make([]Product, 0)
        for i := 0; i < n; i++ {
            fmt.Scanf("%d", &prods[i])
        }

        for i := 0; i < n; i++ {
            for j := 0; j < n; j++ {
                if i == j {
                    continue
                }
                if credit == prods[i] + prods[j] {
                    p := Product{
                        A: i+1,
                        B: j+1,
                        Sum: (prods[i]+prods[j]),
                    }
                    if p.A > p.B {
                        p.A, p.B = p.B, p.A
                    }
                    result = append(result, p)
                }
            }
        }

        sort.Sort(BySum(result))
        fmt.Printf("Case #%d: %d %d\n", index, result[len(result)-1].A, result[len(result)-1].B)
    }
}
