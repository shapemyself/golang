package main




import (
	"strings"
	"strconv"
	"fmt"
)


/**
描述

等差数列是常见数列的一种，如果一个数列从第二项起，每一项与它的前一项的差等于同一个常数，这个数列就叫做等差数列，而这个常数叫做等差数列的公差，公差常用字母d表示。即对于数列S，它满足了(S[i]-S[i-1]) = d (i>1)。
显然，一个数字无法构成等差数列，而任意两个数字可以形成一个等差数列。
这里给出了一个长度为N的数字序列，每个位置有一个整数，需要找到这个数字序列里包含多少个等差数列，序列顺序固定，无需排序。
输入数据格式：S[0] S[1] S[2] …S[N]（以半角空格符分隔，N > 1）
输出数据格式：等差数列数量M；
其中数列S的项为整数
 */

func main() {
	fmt.Println(findAllGradeCount("2 7 4 5 6"))
	fmt.Println(findAllGradeCount("3 3 3 3"))
}


func findAllGradeCount(line string) string{
	var strArr = strings.Split(line, " ")
	var intArr []int
	for _,strNum := range strArr {
		tempInt,_ := strconv.Atoi(strNum)
		intArr = append(intArr, tempInt)
	}


	var length = len(intArr)
	var indexI,indexJ,sum int
	for indexI=0; indexI<length-1; indexI++ {
		for indexJ=indexI+1; indexJ<length; indexJ++ {
			sum ++
			sum += countGrade(intArr, []int{indexI,indexJ})
		}
	}

	return strconv.Itoa(sum)
}

func countGrade(intArr []int,indexes []int) (res int){

	var indexLen = len(indexes)
	var sumLen = len(intArr)
	if indexLen < 2 {
		return
	}

	var maxIndexI = indexes[indexLen-1]

	if maxIndexI >= len(intArr)-1 {
		return
	}

	var dex int = intArr[indexes[1]] - intArr[indexes[0]]
	var index int
	for index=maxIndexI+1; index<sumLen; index++ {
		if intArr[index] == intArr[maxIndexI]+dex {
			indexes = append(indexes, index)
			res += 1+countGrade(intArr,indexes)
		}
	}

	return
}