package services_test

import (
	"go-test/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckGrade(t *testing.T) {

	type testCase struct {
		name     string
		score    int
		expected string
	}

	cases := []testCase{
		{name: "a", score: 80, expected: "A"},
		{name: "b", score: 70, expected: "B"},
		{name: "c", score: 60, expected: "C"},
		{name: "d", score: 50, expected: "D"},
		{name: "f", score: 40, expected: "F"},
	}

	for _, c := range cases {

		t.Run(c.name, func(t *testing.T) {
			grade := services.CheckGrade(c.score)

			assert.Equal(t, c.expected, grade)
		})
	}
}

func BenchmarkCheckGrade(b *testing.B) {
	for i := 0; i < b.N; i++ {
		services.CheckGrade(80)
	}
}
