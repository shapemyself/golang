package main

import (
	"strings"
	"strconv"
	"fmt"
)



/**
描述

给定两个数组，由数字 0-9 组成的，长度分别为 a 和 b，需要用 a、b 两个数组中的数字组合得到长度为 k （k <= a+b）的正整数，输出所有可能的组合中数值最大的一个（原同一数组中的数字顺序不能改变）
输入

a、b 数组中的数字间用 , 隔开，两个数组以及 k 之间用空格隔开，如：1,3,4,5,1,2 8,9,1 5，其中 a = [1,3,4,5,1,2]，b = [8,9,1]，k = 5
输出

长度为 k 的所有组合中数值最大的整数，如：95121

* 从 a 或 b 中取数的时候需保证 a，b 内部的顺序不变，如第一个数取到 b 中的 9，那么 b 中只有 1 可以后续取用；第二个数取到 a 中的 5，则 a 中还剩下 1,2 可以后续取用。
 */



func main() {
	fmt.Println(getMaxCombineNum("6,3,8,9,4,6,0 3,5,7 6"))
	fmt.Println(getMaxCombineNum("2,6,8,4,3 6,9,2,5 3"))
	fmt.Println(getMaxCombineNum("3,7,2 7,9,5,1 7"))
	fmt.Println(getMaxCombineNum("5,9,7,5,3,9,0,5,5,6,3,2,6,8 9,1,4,8 6"))
	fmt.Println(getMaxCombineNum("7,8,7,7,4,4,9,1,5,5,5 8,2,9,8,2,0,7,4 12"))
	fmt.Println(getMaxCombineNum("9,3,1,1,4,9,5,9,0,9,3,4,0,7,5,7,5 0,6,4,7,8,4,0,0,6,2,4 25"))
}


type Index struct {
	indexI int
	indexJ int
	leftStep int
}

type IntArray struct {
	intArr []int
	indexLeft int
	indexRight int
}


func getMaxCombineNum(line string) string{
	var splitArr = strings.Split(line, " ")
	var leftCharArr = strings.Split(splitArr[0], ",")
	var rightCharArr = strings.Split(splitArr[1], ",")
	kNum,_ := strconv.Atoi(splitArr[2])

	var leftIntArr, rightIntArr []int
	for _,value := range leftCharArr {
		temp,_ := strconv.Atoi(value)
		leftIntArr = append(leftIntArr, temp)
	}

	for _,value := range rightCharArr {
		temp,_ := strconv.Atoi(value)
		rightIntArr = append(rightIntArr, temp)
	}

	var index = Index{
		indexI:0,
		indexJ:0,
		leftStep:kNum,
	}
	var step,tempMax int = 0,0
	var resArr []int
	//
	//fmt.Println(leftIntArr)
	//fmt.Println(rightIntArr)

	for step < kNum {
		tempMax, index = findMostRationalData(leftIntArr,rightIntArr,index)
		resArr = append(resArr, tempMax)
		step ++
	}

	var res string = ""
	for _,value := range resArr {
		res += strconv.Itoa(value)
	}

	return res
}

func findMostRationalData(leftArr,rightArr []int,index Index) (int,Index) {

	var leftLen, rigthLen int = len(leftArr), len(rightArr)

	//如果最合适的最大数在左数组中取得
	var leftIndex int = index.indexI
	var tempIndex = leftIndex
	for (leftLen - leftIndex) + (rigthLen - index.indexJ) >= index.leftStep && leftIndex<leftLen{
		if leftArr[leftIndex] > leftArr[tempIndex] {
			tempIndex = leftIndex
		}
		leftIndex ++
	}
	leftIndex = tempIndex

	//如果最合适的最大数在右数组中取得
	var rightIndex int = index.indexJ
	tempIndex = rightIndex
	for (rigthLen - rightIndex) + (leftLen - index.indexI) >= index.leftStep && rightIndex<rigthLen{
		if rightArr[rightIndex] > rightArr[tempIndex] {
			tempIndex = rightIndex
		}
		rightIndex ++
	}
	rightIndex = tempIndex

	if leftIndex >= leftLen {
		index = Index{
			indexI:index.indexI,
			indexJ:rightIndex+1,
			leftStep:index.leftStep-1,
		}
		return rightArr[rightIndex],index
	}

	if rightIndex >= rigthLen {
		index = Index{
			indexI:leftIndex + 1,
			indexJ:index.indexJ,
			leftStep:index.leftStep - 1,
		}
		return leftArr[leftIndex], index
	}

	if leftArr[leftIndex] > rightArr[rightIndex] {
		index = Index{
			indexI:leftIndex + 1,
			indexJ:index.indexJ,
			leftStep:index.leftStep - 1,
		}
		return leftArr[leftIndex], index
	}else if leftArr[leftIndex] < rightArr[rightIndex] {
		index = Index{
			indexI:index.indexI,
			indexJ:rightIndex+1,
			leftStep:index.leftStep-1,
		}
		return rightArr[rightIndex],index
	}else{
		var arrLeft = IntArray{
			intArr:leftArr,
			indexLeft:index.indexI,
			indexRight:leftIndex,
		}

		var arrRight = IntArray{
			intArr:rightArr,
			indexLeft:index.indexJ,
			indexRight:rightIndex,
		}

		res := selectFromSame(arrLeft,arrRight)

		if res == "right" {
			index = Index{
				indexI:index.indexI,
				indexJ:rightIndex+1,
				leftStep:index.leftStep-1,
			}
			return rightArr[rightIndex],index
		}else{
			index = Index{
				indexI:leftIndex + 1,
				indexJ:index.indexJ,
				leftStep:index.leftStep - 1,
			}
			return leftArr[leftIndex], index
		}
	}

}


func selectFromSame(arrLeft,arrRight IntArray) string{


	if arrLeft.indexLeft == arrLeft.indexRight && arrRight.indexLeft == arrRight.indexRight {
		if arrLeft.indexLeft+1<len(arrLeft.intArr) && arrRight.indexLeft+1<len(arrRight.intArr) {
			if  arrLeft.intArr[arrLeft.indexLeft+1] > arrRight.intArr[arrRight.indexLeft+1] {
				return "left"
			}
			if arrLeft.intArr[arrLeft.indexLeft+1] < arrRight.intArr[arrRight.indexLeft+1] {
				return "right"
			}
		}
	}

	if arrLeft.indexLeft >= arrLeft.indexRight {
		return "right"
	}

	if arrRight.indexLeft >= arrRight.indexRight {
		return "left"
	}

	var leftMax,rightMax int = arrLeft.indexLeft,arrRight.indexLeft
	var tempIndex = arrLeft.indexLeft
	for leftMax < arrLeft.indexRight {
		if arrLeft.intArr[leftMax] > arrLeft.intArr[tempIndex] {
			tempIndex = leftMax
		}
		leftMax ++
	}
	leftMax = tempIndex

	tempIndex = arrRight.indexLeft
	for rightMax < arrRight.indexRight {
		if arrRight.intArr[rightMax] > arrRight.intArr[tempIndex] {
			tempIndex = rightMax
		}
		rightMax ++
	}
	rightMax = tempIndex

	if arrLeft.intArr[leftMax] > arrRight.intArr[rightMax] {
		return "right"
	}else if arrLeft.intArr[leftMax] < arrRight.intArr[rightMax] {
		return  "left"
	}else{
		arrLeft = IntArray{
			intArr:arrLeft.intArr,
			indexLeft:leftMax+1,
			indexRight:arrLeft.indexRight,
		}
		arrRight = IntArray{
			intArr:arrRight.intArr,
			indexLeft:rightMax+1,
			indexRight:arrRight.indexRight,
		}

		return selectFromSame(arrLeft,arrRight)
	}

}