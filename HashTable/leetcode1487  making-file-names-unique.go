package HashTable

import "fmt"

func GetFolderNames(names []string) []string {
	nextSuffix := make(map[string]int) // 记录文件夹名称的下一后缀数字
	for i, name := range names {
		nameCopy := name
		if suffix, exist := nextSuffix[nameCopy]; exist { // 检查文件夹名称是否存在
			for { // 逐步试探添加后缀
				nameCopy += fmt.Sprintf("(%d)", suffix)
				if _, exist := nextSuffix[nameCopy]; !exist { // 文件名不存在
					break
				}
				nameCopy = name
				suffix++
			}
			nextSuffix[name] = suffix + 1 // 更新后缀
		}
		nextSuffix[nameCopy]++ // 更新新名称
		names[i] = nameCopy
	}
	return names
}
