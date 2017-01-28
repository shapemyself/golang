package main

import (
	"strings"
	"fmt"
)


/**
描述

给出三个队列 s1，s2，s3 ，判断 s3 是否是由 s1 和 s2 交叉得来。
如：s1 为 aabcc ， s2 为 dbbca。
当 s3 为 aadbbcbcac 时，返回 true（即将 s1 拆成三部分： aa，bc，c 分别插入 s2 对应位置）
否则返回 false。
输入

aabcc,dbbca,aadbbcbcac
输出

true
 */


type Index struct {
	lindex int
	rindex int
	mindex int
}

func (index Index)String() string{
	return fmt.Sprintf("leftIndex=%d; rightIndex=%d, mergeIndex=%d",index.lindex,index.rindex,index.mindex)
}

func main() {
	fmt.Println(judgeCharsMerge("aabcc,dbbca,aadbbcbcac"))
}


func judgeCharsMerge(line string) string{
	strArr := strings.Split(line, ",")
	var leftStr string = strArr[0]
	var rightStr string = strArr[1];
	var mergeStr string = strArr[2];

	//首先判断左，右字符串的长度之和与最终的字符串长度进行比较
	if ( len(leftStr)+len(rightStr) != len(mergeStr) ) {
		return  "false"
	}

	//三个字符串的遍历索引
	var index = Index{
		lindex:0,
		rindex:0,
		mindex:0,
	}

	res := judge(leftStr,rightStr,mergeStr,index)
	if res {
		return "true"
	}

	return "false"
}



func judge(lstr,rstr,mstr string, index Index) bool {
	if ( len(mstr) == index.mindex ) {
		return  true
	}

	var curCharValue = mstr[index.mindex]
	var lmatch bool = false
	var rmatch bool = false
	if len(lstr)!=index.lindex && lstr[index.lindex]==curCharValue {
		lmatch = true
	}
	if len(rstr)!=index.rindex && rstr[index.rindex]==curCharValue {
		rmatch = true
	}


	indexCopyOne := Index{
		index.lindex+1,
		index.rindex,
		index.mindex+1,
	}
	indexCopyTwo := Index{
		index.lindex,
		index.rindex+1,
		index.mindex+1,
	}

	if ( lmatch&&rmatch ) {
		return judge(lstr,rstr,mstr,indexCopyOne) || judge(lstr,rstr,mstr,indexCopyTwo)
	}

	if lmatch {
		return judge(lstr,rstr,mstr,indexCopyOne)
	}

	if rmatch {
		return judge(lstr,rstr,mstr,indexCopyTwo)
	}

	return false
}




























