/*
https://code.google.com/codejam/contest/351101/dashboard
$ go run 2010.africa.a.go < A-large-practice.in
Case #1: 2 3
Case #2: 1 4
Case #3: 4 5
Case #4: 29 46
Case #5: 11 56
Case #6: 84 240
Case #7: 413 584
Case #8: 28 80
Case #9: 380 633
Case #10: 190 242
Case #11: 450 667
Case #12: 7 126
Case #13: 208 636
Case #14: 301 831
Case #15: 243 649
Case #16: 258 429
Case #17: 18 49
Case #18: 7 32
Case #19: 750 965
Case #20: 196 718
Case #21: 292 1236
Case #22: 539 935
Case #23: 20 263
Case #24: 30 107
Case #25: 3 94
Case #26: 425 484
Case #27: 623 923
Case #28: 219 656
Case #29: 133 242
Case #30: 691 800
Case #31: 317 338
Case #32: 207 485
Case #33: 63 64
Case #34: 13 206
Case #35: 407 421
Case #36: 654 972
Case #37: 205 275
Case #38: 499 883
Case #39: 471 633
Case #40: 371 393
Case #41: 25 170
Case #42: 64 521
Case #43: 28 78
Case #44: 949 1590
Case #45: 1 1043
Case #46: 636 773
Case #47: 240 908
Case #48: 6 47
Case #49: 81 288
Case #50: 274 983
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
