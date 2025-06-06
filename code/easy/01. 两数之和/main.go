package main

import "fmt"

// func twoSum(nums []int, target int) []int {
//     for i:=0; i<len(nums)-1;i++{
//         for j:=i+1; j<len(nums);j++{
//             if nums[i] + nums[j] == target{
//                 return []int{i,j}
//             }
//         }
//     }
//     return nil
// }

// 当我们使用遍历整个数组的方式寻找 target - x 时，需要注意到每一个位于 x 之前的元素都已经和 x 匹配过，因此不需要再进行匹配。而每一个元素不能被使用两次，所以我们只需要在 x 后面的元素中寻找 target - x。

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
		// 目标是6
		// 第一次 hashTable[3] 0
		// 第二次 hashTable[2] 1
		// 第三次 p := hashTable[6-4] 找到1位置的hashTable[2]
	}
	return nil
}

func main() {
	fmt.Print(twoSum([]int{3, 2, 4}, 6))
}
