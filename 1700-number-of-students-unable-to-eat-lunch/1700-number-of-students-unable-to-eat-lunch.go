func countStudents(students []int, sandwiches []int) int {
	misses := 0
	for len(students) != misses {
		if students[0] == sandwiches[0] {
			students = students[1:]
			sandwiches = sandwiches[1:]
			misses = 0
		} else {
			students = append(students[1:], students[0])
			misses++
		}
	}
	return misses
}