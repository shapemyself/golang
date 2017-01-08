package main

import (
	"encoding/json"
	"fmt"
)

/**
描述

用一个数组表示一群正在排队的小学生，每个小学生用一对整数来表示 [height, k], height 表示这个小学生的身高，k 表示这个小学生前面应该有 k 个人的身高 >= 他。

写一个算法，对给出的一组小学生计算出符合描述的正确排序。
输入

一行文本如：[[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]，格式为 JSON（算法中可使用 JSON decode 得到数组）。

可以发现该序列不满足描述的条件，比如 [4,4] ：4 前面只有一个 7 大于 4。
输出

根据输入，重新排列数组，然后将符合条件的数组编码为 JSON 输出（不要有空格），如：[[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]。
 */
func main() {
	fmt.Println(helpLineUp("[[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]"))
}


func helpLineUp(line string) string{
	var numArr [][2]int
	json.Unmarshal([]byte(line), &numArr)

	var tempMap = make(map[int][][2]int, 10)

	var minSum,maxSum int
	minSum = numArr[0][0]+numArr[0][1]
	maxSum = minSum
	for _,arrValue := range numArr {
		var sum = arrValue[0]+arrValue[1]
		if sum > maxSum {
			maxSum = sum
		}else if ( sum < minSum ) {
			minSum = sum
		}

		tempMap[sum] = append(tempMap[sum], arrValue)
	}

	var resultArr [][2]int
	var index int
	var studentBefore = make(map[int]int,10)
	var tempRes [][2]int
	for index=minSum; index<=maxSum; index++ {
		if value,ok := tempMap[index];ok{
			tempRes = sortSameSum(value,studentBefore)
			resultArr = append(resultArr,tempRes...)
		}
	}
	bytes,_ := json.Marshal(resultArr)
	return string(bytes)
}


func sortSameSum(inputArr [][2]int,studentBefore map[int]int) ([][2]int) {
	var result [][2]int
	var tempMap map[int][2]int = make(map[int][2]int,10)
	var minHeight,maxHeight,index int
	minHeight = inputArr[0][0]
	maxHeight = minHeight
	for _,value := range inputArr {
		if ( value[0] > maxHeight ) {
			maxHeight = value[0]
		}else if( value[0] < minHeight ) {
			minHeight = value[0]
		}

		tempMap[value[0]] = value
	}

	for{
		if len(tempMap) == 0 {
			break
		}
		for index=minHeight; index<=maxHeight;index++ {
			if value,ok := tempMap[index]; ok {
				numBefore := 0
				for mkey,mvalue := range studentBefore {
					if mkey >= value[0] {
						numBefore += mvalue
					}
				}
				if numBefore == value[1] {
					studentBefore[value[0]]++
					result = append(result, value)
					delete(tempMap,value[0])
					break
				}
			}
		}
	}

	return  result
}









