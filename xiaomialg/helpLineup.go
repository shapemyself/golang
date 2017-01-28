package main

import (
	"encoding/json"
	"fmt"
	"sort"
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
	fmt.Println(helpLineUp("[[7,0],[4,4],[7,1],[5,0],[6,1],[5,2],[4,5],[6,2],[7,2]]"))
	//helpLineUp("[[7,0],[4,4],[7,1],[5,0],[6,1],[5,2],[4,5]]")
}

type Student struct {
	height int `身高`
	numbefore int `前面对自己高的人数`
}
func (s Student) String() string{
	return fmt.Sprintf("[%d,%d]", s.height, s.numbefore)
}

type StundetList []Student
func (list StundetList) Len() int{
	return len(list)
}
func (list StundetList) Less(i,j int) bool{
	ires := list[i].height + list[i].numbefore
	jres := list[j].height + list[j].numbefore
	if ires < jres {
		return  true
	}else{
		return false
	}
}
func (list StundetList) Swap(i,j int) {
	var temp = list[i]
	list[i] = list[j]
	list[j] = temp
}




func helpLineUp(line string) string{
	var numArr [][2]int
	json.Unmarshal([]byte(line), &numArr)
	var list StundetList
	for _,value := range numArr {
		list = append(list,Student{value[0],value[1]})
	}
	sort.Sort(list)

	var resQuen =[]int{0}
	resQuen = findResult(resQuen, list, 1)

	var sortSet [][2]int
	for _,value := range resQuen {
		sortSet = append(sortSet, [2]int{list[value].height,list[value].numbefore})
	}

	ret,_ := json.Marshal(sortSet)
	return string(ret)
}

func findResult(quen []int,list StundetList,startIndex int) []int{
	var maxIndex = len(list)
	var length = len(quen)
	if  length == 0 || length==len(list){
		return quen
	}
	var index int
	for index=startIndex; index<maxIndex; index++ {
		if res := isRightPos(quen, index, list); res {
			quen = append(quen, index)
			break
		}
	}

	if length != len(quen) {
		quen = findResult(quen, list, 0)
	}else{
		//回滚
		startIndex = quen[length-1]+1
		quen = quen[0:length-1]
		quen = findResult(quen, list, startIndex)
	}
	return quen
}

func isItemInSlice(arr []int, item int) bool{
	var ret = false
	if ( len(arr) == 0 ) {
		return ret
	}

	for _,value := range arr {
		if value == item {
			ret = true
			break
		}
	}

	return ret
}

func isRightPos(arr []int, curIndex int, list StundetList) bool{
	if res := isItemInSlice(arr, curIndex) ;res {
		return false
	}

	var count = 0
	for _,value := range arr {
		if list[value].height >= list[curIndex].height {
			count ++
		}
	}

	if count == list[curIndex].numbefore {
		return  true
	}

	return  false
}











