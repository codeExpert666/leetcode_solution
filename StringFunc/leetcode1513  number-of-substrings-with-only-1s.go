package StringFunc

// NumSub 模拟
func NumSub(s string) int {
	var res int
	var mod = int(1e9 + 7)
	var oneCount int              // 统计当前阶段1的数目
	for i := 0; i < len(s); i++ { // 遍历字符串
		if s[i] == '1' { // 遇到 1 则累加
			oneCount++
		} else if i != 0 && s[i-1] == '1' { // 遇到 0 且前一个字符为 1
			res = (res + (oneCount*(oneCount+1))>>1) % mod // 统计现阶段 1 能形成的子串
			oneCount = 0                                   // 重置 1 的数量
		}
	}
	return (res + (oneCount*(oneCount+1))>>1) % mod // 最后一段 1 尚未统计
}
