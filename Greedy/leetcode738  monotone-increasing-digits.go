package Greedy

import "strconv"

/**
 * 贪心法
 * 对于一个输入数字，从低位向高位遍历，按照输出各位需要递增的要求，每遇到
 * 低位小于相邻高位的情况，就将所有低位赋为9，高位减一
 * 这里的贪心还体现在低位到高位遍历，每次减一都是先处理“较低”的高位，这样就能保证每次减少的幅度最小
 * 最终保证输出最大
 */
func monotoneIncreasingDigits(N int) int {
	s := strconv.Itoa(N) //将数字转为字符串，方便使用下标
	ss := []byte(s)      //将字符串转为byte数组，方便更改。
	n := len(ss)
	if n <= 1 {
		return N
	}
	for i := n - 1; i > 0; i-- {
		if ss[i-1] > ss[i] { //前一个大于后一位,前一位减1，后面的全部置为9
			ss[i-1] -= 1
			for j := i; j < n; j++ { //后面的全部置为9
				ss[j] = '9'
			}
		}
	}
	res, _ := strconv.Atoi(string(ss))
	return res
}

/**
 * 上一方法的改进版：不用每次发现非递增都遍历到最低位置9
 * 缺点是空间复杂度稍稍提高，多了一个int类型变量
 */
func monotoneIncreasingDigits1(N int) int {
	s := strconv.Itoa(N) //将数字转为字符串，方便使用下标
	ss := []byte(s)      //将字符串转为byte数组，方便更改。
	n := len(ss)
	if n <= 1 {
		return N
	}
	flag := n // 记录'9'从哪里开始将低位置9
	for i := n - 1; i > 0; i-- {
		if ss[i-1] > ss[i] { //前一个大于后一位,前一位减1，后面的全部置为9
			ss[i-1] -= 1
			//for j := i; j < n; j++ { //后面的全部置为9
			//	ss[j] = '9'
			//}
			flag = i
		}
	}
	for i := flag; i < n; i++ { // 低位置9
		ss[i] = '9'
	}
	res, _ := strconv.Atoi(string(ss))
	return res
}
