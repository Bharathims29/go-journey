package main

import "testing"

func TestCalculateGrade(t *testing.T) {
	tests := []struct {
		name     string
		marks    int
		maxMarks int
		want     string
	}{
		{name: "a plus", marks: 95, maxMarks: 100, want: "A+"},
		{name: "a", marks: 80, maxMarks: 100, want: "A"},
		{name: "b", marks: 70, maxMarks: 100, want: "B"},
		{name: "c", marks: 60, maxMarks: 100, want: "C"},
		{name: "d", marks: 50, maxMarks: 100, want: "D"},
		{name: "fail", marks: 49, maxMarks: 100, want: "F"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateGrade(tt.marks, tt.maxMarks)
			if got != tt.want {
				t.Fatalf("calculateGrade() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestValidateInput(t *testing.T) {
	input := StudentMarkInput{
		StudentName: "Bharath",
		RollNumber:  "R001",
		Subject:     "Math",
		Marks:       88,
		MaxMarks:    100,
	}

	if err := validateInput(input); err != nil {
		t.Fatalf("validateInput() error = %v", err)
	}
}

func TestValidateInputRejectsMarksGreaterThanMaxMarks(t *testing.T) {
	input := StudentMarkInput{
		StudentName: "Bharath",
		RollNumber:  "R001",
		Subject:     "Math",
		Marks:       101,
		MaxMarks:    100,
	}

	if err := validateInput(input); err == nil {
		t.Fatal("expected validation error, got nil")
	}
}

func TestIDFromPath(t *testing.T) {
	got, err := idFromPath("/marks/42")
	if err != nil {
		t.Fatalf("idFromPath() error = %v", err)
	}

	if got != 42 {
		t.Fatalf("idFromPath() = %d, want 42", got)
	}
}
