package main

import (
	"strconv"
	"fmt"
)

func range_lst(lo int, up int, res string) string {
	if lo > up {
		return res
	}
	res += "," + strconv.FormatInt(int64(lo), 10)
	return range_lst(lo+1, up, res)
}

func visit(lst []int, idx int, size int, pre int, res string, lo int, up int) string {
	if idx > size {
		return res
	}
	cur := pre + 5
	if idx < size {
		cur = lst[idx]
	}
	// fmt.Println(lo, up, pre, cur)
	if lo == up {
		if cur == pre + 1 {
			lo = pre
			up = cur
		} else {
			res += "," + strconv.FormatInt(int64(pre), 10)
		}
	} else {
		if cur == pre + 1 {
			up = cur
		} else {
			if (up - lo + 1) < 3 {
				res += range_lst(lo, up, "")
			} else {
				res += "," + strconv.FormatInt(int64(lo), 10) + "-" + strconv.FormatInt(int64(up), 10)
			}
			lo = cur
			up = cur
		}
	}
	pre = cur
	return visit(lst, idx+1, size, pre, res, lo, up)
}


func extract(lst []int) string {
	fmt.Println(lst)
	if lst == nil{
		return ""
	}
	if len(lst) == 1 {
		return strconv.FormatInt(int64(lst[0]), 10)
	}
	res := ""
	idx := 1
	size := len(lst)
	pre := lst[0]
	lo := pre
	up := pre
	return visit(lst, idx, size, pre, res, lo, up)[1:]
}


func main() {

	lst := []int{-6,-3,-2,-1,0,1,3,4,5,7,8,9,10,11,14,15,17,18,19,20,22}
	// lst := []int{}
	// lst := []int{-3}
	// fmt.Println(lst)
	res := extract(lst)
    fmt.Println(res)
}


