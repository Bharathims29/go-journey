package main

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5

	if got != want {
		t.Fatalf("Add(2, 3) = %d, want %d", got, want)
	}
}

func TestDivide(t *testing.T) {
	got, err := Divide(10, 2)
	requireNoError(t, err)

	want := 5
	if got != want {
		t.Fatalf("Divide(10, 2) = %d, want %d", got, want)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := Divide(10, 0)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestIsValidUsername(t *testing.T) {
	tests := []struct {
		name     string
		username string
		want     bool
	}{
		{name: "valid username", username: "bharath", want: true},
		{name: "too short", username: "go", want: false},
		{name: "contains space", username: "go learner", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidUsername(tt.username)
			if got != tt.want {
				t.Fatalf("IsValidUsername(%q) = %v, want %v", tt.username, got, tt.want)
			}
		})
	}
}

func TestNormalizeName(t *testing.T) {
	got := NormalizeName("  Bharath  ")
	want := "bharath"

	if got != want {
		t.Fatalf("NormalizeName() = %q, want %q", got, want)
	}
}

func BenchmarkNormalizeName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NormalizeName("  Bharath  ")
	}
}

func requireNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
