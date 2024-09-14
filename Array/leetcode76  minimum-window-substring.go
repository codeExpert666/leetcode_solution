package Array

// 滑动窗口（自己实现的滑动窗口，该窗口有些特别，不仅是传统的两个边界指针，还额外开辟空间存储了窗口中有价值的信息）
func minWindow(s string, t string) string {
	// 用哈希表替换字符串t，方便匹配
	// 哈希表的键是字符，值对应该字符出现的次数
	t_map := make(map[byte]int)
	for _, c := range t {
		t_map[byte(c)]++
	}
	// 滑动窗口遍历字符串s
	minLen, minLeft, minRight := len(s)+1, 0, 0 // 记录当前最小覆盖子串的信息
	window := make([]int, 0, len(s))            // 定义窗口，注意窗口中只包含t中元素出现的位置
	count := 0                                  // 记录窗口中已经完成寻找的元素个数
	for i, c := range s {
		if value, ok := t_map[byte(c)]; ok {
			window = append(window, i) // 移动窗口右边界
			t_map[byte(c)]--           // 更新每个元素的剩余出现次数
			if value <= 1 {            // 元素c已经寻找完毕
				if value == 1 {
					count++ // 完成寻找的元素个数增加
				} else {
					// 这种情况value为0，表明窗口中新加入元素的个数大于实际需求个数
					// 试图移动窗口左边界以移除多余的元素
					for t_map[byte(s[window[0]])] < 0 {
						t_map[byte(s[window[0]])]++
						window = window[1:]
					}
				}
				if count == len(t_map) && window[len(window)-1]-window[0]+1 < minLen {
					// 所有元素完成寻找
					minLeft, minRight = window[0], window[len(window)-1]
					minLen = minRight - minLeft + 1
				}
			}
		}
	}
	if minLen == len(s)+1 { // 未发生更新，说明覆盖子串不存在
		return ""
	}
	return s[minLeft : minRight+1]
}
