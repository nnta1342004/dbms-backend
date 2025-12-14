package appCommon

func HasDuplicates(nums []any) bool {
	seen := make(map[any]bool)

	for _, num := range nums {
		if seen[num] {
			return true // Found a duplicate
		}
		seen[num] = true
	}

	return false // No duplicates found
}
