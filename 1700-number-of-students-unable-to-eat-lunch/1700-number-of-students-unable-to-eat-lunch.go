func countStudents(students []int, sandwiches []int) int {
	counter := 0
	for len(students) != 0 {
		if students[0] == sandwiches[0] {
			students = students[1:]
			sandwiches = sandwiches[1:]
			counter = 0
		} else {
			students = append(students[1:], students[0])
			counter++
			if counter == len(students) {
				return len(students)
			}
		}
	}
	return 0
}