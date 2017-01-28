package main

import (
	"strings"
	"strconv"
	"fmt"
)


/**
描述

有一行由正数组成的数字字符串，移除其中的 K 个数，使剩下的数字是所有可能中最小的。

假设：
* 字符串的长度一定大于等于 K
* 字符串不会以 0 开头
输入

一行由正整数组成的数字字符串，和一个正整数 K，两个数据由英文逗号隔开，如：1432219,3。
输出

移除 K 位后可能的最小的数字字符串。
如 1432219 移除 4, 3, 2 这 3 个数字后得到 1219，为所有可能中的最小值。
 */


func main() {
	fmt.Println(keepSmallByRemoveKnum("1432219,3"))
	fmt.Println(keepSmallByRemoveKnum("10200,1"))
}


func keepSmallByRemoveKnum(line string) string {
	var strArr = strings.Split(line, ",")
	var sourceStr string = strArr[0]
	kNum ,_ := strconv.Atoi(strArr[1])

	if len(sourceStr) == kNum {
		return "0"
	}

	var num int
	for num=0; num < kNum; num++ {
		sourceStr = removeAnumKeepSmallest(sourceStr)
	}

	temp,_ := strconv.Atoi(sourceStr)

	return strconv.Itoa(temp)
}

func removeAnumKeepSmallest(dealStr string) (res string) {
	if len(dealStr) == 0 {
		return ""
	}

	var fIndex int
	for fIndex=0; fIndex<len(dealStr); fIndex++ {
		if fIndex+1 < len(dealStr) {
			if dealStr[fIndex] > dealStr[fIndex+1] {
				break
			}
		}else{
			break
		}
	}

	charSlice := []rune(dealStr)
	length := len(charSlice)

	if fIndex == 0 {
		res = string(charSlice[fIndex+1:length])
	}else if fIndex+1 == length {
		res = string(charSlice[0:fIndex])
	}else{
		res = string(charSlice[0:fIndex]) + string(charSlice[fIndex+1:length])
	}
	return
}


