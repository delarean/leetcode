package main

func twoSum(nums []int, target int) []int {
    mp := make(map[int]int, len(nums))

    for i, num := range nums {
        mp[target-num] = i
    }

    for i, num := range nums {
        if ind, ok := mp[num]; ok && i != ind {
            return []int{i, ind}
        }
    }

    return []int{}
}
