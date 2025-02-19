package main

/*
	美团2024年秋季笔试 01场 11题
	小美准备登录美团，需要输入密码，小美忘记了密码，只记得密码可能是 n 个字符串中的一个。
	小美会按照密码的长度从小到大依次尝试每个字符串，对于相同长度的字符串，小美随机尝试，并且相同的密码只会尝试一次。
	小美想知道，她最少需要尝试多少次才能登录成功，最多需要尝试多少次才能登录成功。
	小美不会重新尝试已经尝试过的字符串。成功登录后会立即停止尝试。
	题目链接：https://www.nowcoder.com/exam/test/85714251/detail?pid=58083953&examPageSource=Company&subTabName=written_page&testCallback=https%3A%2F%2Fwww.nowcoder.com%2Fexam%2Fcompany%3FquestionJobId%3D10%26subTabName%3Dwritten_page&testclass=%E8%BD%AF%E4%BB%B6%E5%BC%80%E5%8F%91
*/

import "fmt"

func MT11() {
	var n int
	_, _ = fmt.Scan(&n)

	var truePwd string
	_, _ = fmt.Scan(&truePwd)

	// less: 小于 truePwd 长度的不重复密码数量; equal: 等于 truePwd 长度的不重复密码数量
	var less, equal int
	pwdSet := make(map[string]struct{})
	var candidatePwd string
	for range n {
		_, _ = fmt.Scan(&candidatePwd)
		if _, ok := pwdSet[candidatePwd]; !ok {
			if len(candidatePwd) == len(truePwd) {
				equal++
				pwdSet[candidatePwd] = struct{}{}
			} else if len(candidatePwd) < len(truePwd) {
				less++
				pwdSet[candidatePwd] = struct{}{}
			}
		}
	}

	fmt.Print(less+1, less+equal)
}
