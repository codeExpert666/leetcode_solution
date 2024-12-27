package Interview

// 字符串仅包含小写字母，采用数组哈希
func IsUnique(astr string) bool {
	// 全是小写字母，所以长度超过 26 则一定重复
	if len(astr) > 26 {
		return false
	}
	// 遍历字符串，哈希存储已经遍历的字符
	// 字符不重复等价于新出现字符不能在哈希表中
	chars := [26]bool{} // 数组代替哈希表
	for i := 0; i < len(astr); i++ {
		if chars[astr[i]-'a'] { // 字符已出现
			return false
		}
		chars[astr[i]-'a'] = true
	}
	return true
}

// 字符串改为仅包含 ascii 字符集中的字符（0～127），数组哈希仍然可做
// 额外空间为 bool 数组占据的空间，大小为 128 字节
func IsUnique1(astr string) bool {
	// ASCII字符集为128个字符，所以长度超过128则一定重复
	if len(astr) > 128 {
		return false
	}
	// 使用128大小的数组来记录ASCII字符
	chars := [128]bool{}
	for i := 0; i < len(astr); i++ {
		if chars[astr[i]] { // 直接使用ASCII码值作为索引
			return false
		}
		chars[astr[i]] = true
	}
	return true
}

// 在 ascii 字符集的情况下，位运算在空间消耗上更为优秀，仅占用 16 字节
// 因为ASCII字符集只需要128位（16字节）就可以表示
func IsUnique2(astr string) bool {
	// 使用两个uint64表示128位
	var mark1, mark2 uint64
	for i := 0; i < len(astr); i++ {
		pos := uint(astr[i])
		if pos < 64 {
			if mark1&(1<<pos) != 0 {
				return false
			}
			mark1 |= 1 << pos
		} else {
			pos -= 64
			if mark2&(1<<pos) != 0 {
				return false
			}
			mark2 |= 1 << pos
		}
	}
	return true
}

// 字符串改为包含 unicode 字符集，则没有什么捷径，因为 unicode 字符集太大，只能老老实实用 map
// Unicode字符集的范围是 0x0000 ~ 0x10FFFF，共计 17 * 65536 = 1114112 个字符
func IsUniqueUnicode(astr string) bool {
	// 使用map存储已出现的rune
	seen := make(map[rune]bool)
	for _, ch := range astr { // 使用range遍历rune而不是byte
		if seen[ch] {
			return false
		}
		seen[ch] = true
	}
	return true
}

// 也可稍微优化，主要思路是针对常用字符区间进行优化
func IsUniqueUnicodeOptimized(astr string) bool {
	// 针对常用Unicode区间做优化
	var (
		basicLatin [128]bool             // ASCII字符 (0-127)
		latin1     [128]bool             // Latin-1补充字符 (128-255)
		others     = make(map[rune]bool) // 其他Unicode字符
	)

	for _, ch := range astr {
		switch {
		case ch < 128: // ASCII范围
			if basicLatin[ch] {
				return false
			}
			basicLatin[ch] = true
		case ch < 256: // Latin-1补充范围
			if latin1[ch-128] {
				return false
			}
			latin1[ch-128] = true
		default: // 其他Unicode字符
			if others[ch] {
				return false
			}
			others[ch] = true
		}
	}
	return true
}
