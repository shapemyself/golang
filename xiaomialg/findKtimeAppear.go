package main;

/**
描述:
有一个不为空且仅包含正整数的数组，找出其中出现频率最高的前 K 个数，时间复杂度必须在 O(n log n) 以内。
输入:
一行数据包括两部分，一个正整数数组（数字间 ',' 分隔）和一个正整数 K （1 ≤ K ≤ 数组长度），数组和 K 之间有一个空格。
输出:
输出包含前 K 个出现频率最高的数，用 ', ' 分隔，保证升序排列。
 */

import (
	"strings"
	"strconv"
	"sort"
	"fmt"
)

func main() {
	fmt.Println(findKNum("1,2,2,3,4,5,3,3,3,2,2,2,1,1,6,6,6 3"))
}

type node struct{
	num string
	count int
}

func findKNum(line string)(string) {
	var tempArr = strings.Split(line, " ")
	var kValue,_ = strconv.Atoi(tempArr[1])
	var numList = strings.Split(tempArr[0], ",")

	var countMap = make(map[string]int,kValue);

	for _,value := range numList {
		if _,ok := countMap[value];ok {
			countMap[value]++;
		}else{
			countMap[value] = 1;
		}
	}

	var countSilce []node

	for key,value := range countMap {
		countSilce = append(countSilce, node{key,value})
	}
	sumLen := len(countSilce)



	var kNumSlice = countSilce[0:kValue]
	kNumSlice = blanceHeap(kNumSlice, kValue)

	var index int
	for index=kValue; index<sumLen; index++ {
		if kNumSlice[0].count < countSilce[index].count {
			kNumSlice[0] = countSilce[index]
			kNumSlice = blanceHeap(kNumSlice, kValue)
		}
	}

	var resArr []string
	for _,v := range kNumSlice {
		resArr = append(resArr,v.num)
	}
	sort.Strings(resArr)
	var res = resArr[0]
	var i int
	for i=1;i<len(resArr);i++ {
		res = res + "," + resArr[i]
	}

	return  res
}

func blanceHeap(kArr []node, kvalue int)([]node) {
	var curIndex = kvalue-1
	var parentIndex = (curIndex+1)/2 - 1
	for curIndex>=0 && parentIndex>=0 {
		if kArr[parentIndex].count > kArr[curIndex].count {
			kArr[parentIndex],kArr[curIndex] = kArr[curIndex],kArr[parentIndex]
		}
		curIndex--
		parentIndex = (curIndex+1)/2 - 1
	}

	return  kArr
}

