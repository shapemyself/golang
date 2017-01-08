package main

import (
	"strings"
	"strconv"
	"sort"
	"fmt"
)
/**
输入

一个包含多个正整数的字符串，每个正整数之间用逗号（,）分隔，如：-1,0,1,2,-1,-4。
输出

输入满足加合结果正好等于 0 的三元组的个数，如对于 -1,0,1,2,-1,-4 有 [-1, 0, 1] 和 [-1, -1, 2]，所以输出 2
 */


func main() {
	fmt.Println(solution("-1,0,1,2,-1,-4"))
}

func solution(line string) (string ){
	var numArr = strings.Split(line, ",")

	var zeroArr []int
	var positiveArr []int
	var negativeArr []int

	for _,value := range numArr {
		var intValue,_ = strconv.Atoi(value)
		if intValue == 0 {
			zeroArr = append(zeroArr,0)
		}else if intValue > 0 {
			positiveArr = append(positiveArr,intValue)
		}else{
			negativeArr = append(negativeArr, intValue)
		}
	}

	var countMap = make(map[string]int,1)


	if ( len(zeroArr) > 0 ) {

		//首先是全0的情况
		if ( len(zeroArr) >= 3 ) {
			countMap["000"]++;
		}

		//其次是一正一负,还有个0的情况
		if ( len(positiveArr)>0 && len(negativeArr)>0 ) {
			for _, pValue := range positiveArr {
				for _,nValue := range negativeArr {
					if pValue + nValue == 0 {
						key := strconv.Itoa(nValue)+"0"+strconv.Itoa(pValue)
						countMap[key]++
						break;
					}
				}
			}
		}
	}

	//考虑一正两负数
	if ( len(positiveArr)>=1 && len(negativeArr)>=2 ) {
		for _,pValue := range positiveArr {
			var indexI,indexJ int
			for indexI=0; indexI<len(negativeArr);indexI++ {
				for indexJ=indexI+1; indexJ<len(negativeArr);indexJ++ {
					if negativeArr[indexI]+negativeArr[indexJ]+pValue == 0 {
						var tempIntArr = []int{pValue,negativeArr[indexJ],negativeArr[indexI]};
						sort.Ints(tempIntArr)
						key := strconv.Itoa(tempIntArr[0])+strconv.Itoa(tempIntArr[1])+strconv.Itoa(tempIntArr[2])
						countMap[key]++
					}
				}
			}
		}
	}

	//考虑一负两正数
	if ( len(positiveArr)>=2 && len(negativeArr)>=1 ) {
		for _,nValue := range negativeArr {
			var indexI,indexJ int
			for indexI=0; indexI<len(positiveArr);indexI++ {
				for indexJ=indexI+1; indexJ<len(positiveArr);indexJ++ {
					if positiveArr[indexI]+positiveArr[indexJ]+nValue == 0 {
						var tempIntArr = []int{nValue,positiveArr[indexJ],positiveArr[indexI]};
						sort.Ints(tempIntArr)
						key := strconv.Itoa(tempIntArr[0])+strconv.Itoa(tempIntArr[1])+strconv.Itoa(tempIntArr[2])
						countMap[key]++
					}
				}
			}
		}
	}

	return  strconv.Itoa(len(countMap))
}
