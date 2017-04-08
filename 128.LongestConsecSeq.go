func longestConsecutive(nums []int) int {
    longest := 0
    seen := make(map[int]bool)

    for _, n := range nums {
        seen[n] = true
    }
    for k, _ := range seen {
        seq := 1
        seen[k] = false
        num := k
        for {
            num = num - 1
            if seen[num] {
                seq = seq + 1
                seen[num] = false
            } else {
                break
            }
        }
        num = k
        for {
            num = num + 1
            if seen[num] {
                seq = seq + 1
                seen[num] = false
            } else {
                break
            }
        }
        if seq > longest {
            longest = seq
        }
    }
    
    return longest
}
