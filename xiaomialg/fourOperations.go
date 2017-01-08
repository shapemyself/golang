package main

import (
	"strings"
	"strconv"
	"fmt"
)

/**
描述

实现一个算法，可以进行任意非负整数的加减乘除组合四则运算。
请注意运算符的优先级。
输入

请输入一行算式，使用空格分隔数字与运算符。
数字为任意非负整数，运算符为+ - * /，不考虑括号。
输出

输出算式的运算结果。如果是小数，请向下取整（包含中间步骤结果）。
如果出现“除0异常”，输出err。
 */

func main() {
	fmt.Println(fourOperations("3 + 5"))
	fmt.Println(fourOperations("12 + 45 / 9"))
	fmt.Println(fourOperations("1 / 2"))
	fmt.Println(fourOperations("1 / 0"))
	fmt.Println(fourOperations("12 + 34 * 56 - 78"))
}

func fourOperations(line string) string {
	var calNumArr = strings.Split(line, " ")
	var queue []string
	var strLeft string

	var index int
	var arrLen = len(calNumArr)
	for index=0; index<arrLen; index++ {
		switch calNumArr[index] {
		case "*":
			strLeft,queue = pop(queue)
			var intLeft,_ = strconv.Atoi(strLeft)
			index++;
			var intRight,_ = strconv.Atoi(calNumArr[index])
			var res = intLeft * intRight
			queue = append(queue, strconv.Itoa(res))
		case "/":
			strLeft,queue = pop(queue)
			var intLeft,_ = strconv.Atoi(strLeft)
			index++
			var intRight,_ = strconv.Atoi(calNumArr[index])
			if intRight == 0 {
				return "err"
			}else{
				var res = intLeft/intRight
				queue = append(queue, strconv.Itoa(res))
			}
		case "+":
			fallthrough
		case "-":
			fallthrough
		default:
			queue=append(queue,calNumArr[index])
		}
	}
	var res,_ = strconv.Atoi(queue[0])
	for index=1; index<len(queue); index++ {
		if queue[index] == "+" {
			index++;
			var intRight,_ = strconv.Atoi(queue[index])
			res += intRight
		}
		if queue[index] == "-" {
			index++;
			var intRight,_ = strconv.Atoi(queue[index])
			res -= intRight
		}
	}

	return  strconv.Itoa(res)
}

func pop(quen []string) (string, []string){
	if len(quen) == 0 {
		return  "",quen
	}else{
		var res = quen[len(quen)-1]
		quen = quen[0:len(quen)-1]
		return  res,quen
	}
}
