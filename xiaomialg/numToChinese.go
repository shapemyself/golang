package main

import (
	"strconv"
	"fmt"
)

var consants = map[int]string{
	0:"零",
	1:"壹",
	2:"贰",
	3:"叁",
	4:"肆",
	5:"伍",
	6:"陆",
	7:"柒",
	8:"捌",
	9:"玖",
	10:"拾",
	100:"佰",
	1000:"仟",
	10000:"万",
	100000000:"亿",
}

func main() {
	//fmt.Println(convertNumToChinese("0"))
	//fmt.Println(convertNumToChinese("5"))
	//fmt.Println(convertNumToChinese("233"))
	//fmt.Println(convertNumToChinese("1001"))
	//fmt.Println(convertNumToChinese("40607"))
	fmt.Println(convertNumToChinese("4494800045"))
}

func convertNumToChinese(line string) string{


	var prefix = "元整"
	var intNum,_ = strconv.Atoi(line)
	var res = ""

	if intNum == 0 {
		return consants[0]+prefix
	}


	//计算超过亿以上
	var dealValue = intNum / 100000000
	intNum = intNum % 100000000
	if dealValue > 0 {
		var ret string
		if len(res) > 0 {
			ret = innerFunc(dealValue,true);
		}else{
			ret = innerFunc(dealValue,false);
		}
		res += ret+consants[100000000]
	}

	//计算亿到万之内
	dealValue = intNum / 10000
	intNum = intNum % 10000
	if dealValue > 0 {
		var ret string
		if len(res) > 0 {
			ret = innerFunc(dealValue,true)
		}else{
			ret = innerFunc(dealValue,false)
		}
		res += ret+consants[10000]
	}

	//计算万以内
	dealValue = intNum
	if len(res) > 0 {
		res += innerFunc(dealValue,true)+prefix
	}else{
		res += innerFunc(dealValue,false)+prefix
	}

	return  res
}

func innerFunc(num int, flag bool) string {
	var res = ""
	var curDealQian = num / 1000;
	num = num % 1000;
	if curDealQian > 0  {
		res += consants[curDealQian]+consants[1000]
	}

	var curDealBai = num / 100;
	num = num % 100
	if curDealBai > 0 {
		if curDealQian == 0 && flag {
			res += consants[0]+consants[curDealBai] + consants[100]
		}else {
			res += consants[curDealBai] + consants[100]
		}
	}

	var curDealShi = num / 10
	num = num % 10
	if curDealShi > 0 {
		if curDealBai == 0 && (len(res)>0 || flag){
			res += consants[0]+consants[curDealShi]+consants[10]
		}else{
			res += consants[curDealShi] + consants[10]
		}
	}

	if num > 0 {
		if curDealShi == 0 && (len(res)>0 || flag){
			res += consants[0]+consants[num]
		}else{
			res += consants[num]
		}
	}

	return res
}
