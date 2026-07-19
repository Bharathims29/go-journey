package main

import "fmt"

func main() {
	fmt.Println("Module 5: Arrays, Slices, and Maps")
	fmt.Println()

	numbers := [3]int{10, 20, 30}

	fmt.Println("Array")
	fmt.Println("First number:", numbers[0])
	fmt.Println("Second number:", numbers[1])
	fmt.Println("Third number:", numbers[2])
	fmt.Println()

	scores := []int{80, 90, 75}

	fmt.Println("Slice before append")
	fmt.Println("Scores:", scores)
	fmt.Println("Length:", len(scores))
	fmt.Println("Capacity:", cap(scores))
	fmt.Println()

	scores = append(scores, 88, 95, 105, 200)

	fmt.Println("Slice after append")
	fmt.Println("Scores:", scores)
	fmt.Println("Length:", len(scores))
	fmt.Println("Capacity:", cap(scores))
	fmt.Println()

	fmt.Println("Range over slice")
	for index, score := range scores {
		fmt.Println("Index:", index, "Score:", score)
	}
	fmt.Println()

	total := 0
	maxScore := scores[0]
	minScore := scores[0]

	for _, score := range scores {
		total += score

		if score > maxScore {
			maxScore = score
		}

		if score < minScore {
			minScore = score
		}
	}

	average := float64(total) / float64(len(scores))

	fmt.Println("Slice calculations")
	fmt.Println("Total:", total)
	fmt.Println("Average:", average)
	fmt.Println("Max:", maxScore)
	fmt.Println("Min:", minScore)
	fmt.Println()

	studentMarks := map[string]int{
		"Bharathi": 90,
		"Anu":      85,
		"Ravi":     78,
		"raja":     89,
	}

	fmt.Println("Map")
	fmt.Println("Student marks:", studentMarks)
	fmt.Println("Bharath marks:", studentMarks["Bharath"])
	fmt.Println()

	studentMarks["Ravi"] = 82
	studentMarks["Meera"] = 88

	fmt.Println("Map after update and add")
	fmt.Println("Student marks:", studentMarks)
	fmt.Println()

	name := "Anu"
	marks, exists := studentMarks[name]

	fmt.Println("Map lookup")
	if exists {
		fmt.Println(name, "marks:", marks)
	} else {
		fmt.Println(name, "not found")
	}
	fmt.Println()

	delete(studentMarks, "Anu")

	fmt.Println("Map after delete")
	for student, marks := range studentMarks {
		fmt.Println(student, ":", marks)
	}
}
