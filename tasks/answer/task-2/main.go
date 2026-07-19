package main

import "fmt"

const subjectCount = 5
const passingMark = 40

type Student struct {
	Name  string
	Marks [subjectCount]int
}

func getStudentDetails() Student {
	var student Student

	fmt.Print("Enter Name: ")
	fmt.Scan(&student.Name)

	fmt.Println("Enter marks for 5 subjects:")

	for i := 0; i < len(student.Marks); i++ {
		fmt.Printf("Subject %d: ", i+1)
		fmt.Scan(&student.Marks[i])
	}

	return student
}

func getTotalMarks(student Student) int {
	total := 0

	for _, mark := range student.Marks {
		total += mark
	}

	return total
}

func getAverage(total int) float64 {
	return float64(total) / subjectCount
}

func getGrade(avg float64) string {
	switch {
	case avg >= 90:
		return "A"
	case avg >= 80:
		return "B"
	case avg >= 70:
		return "C"
	case avg >= 60:
		return "D"
	case avg >= 40:
		return "E"
	default:
		return "F"
	}
}

func getResult(user Student) string {
	for _, mark := range user.Marks {
		if mark < passingMark {
			return "FAIL"
		}
	}
	return "PASS"
}

func main() {
	student := getStudentDetails()
	total := getTotalMarks(student)
	average := getAverage(total)
	grade := getGrade(average)
	result := getResult(student)

	fmt.Println("\nStudent Details")
	fmt.Println("Name:", student.Name)

	fmt.Println("Marks:")
	for i, mark := range student.Marks {
		fmt.Printf("Subject %d: %d\n", i+1, mark)
	}
	fmt.Println("Total   :", total, "/", subjectCount*100)
	fmt.Printf("Average : %.2f%%\n", average)
	fmt.Println("Grade   :", grade)
	fmt.Println("Result  :", result)
}
