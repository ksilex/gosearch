package main

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestInts(t *testing.T) {

	got := []int{5, 3, 7}
	Ints(got)
	want := []int{3, 5, 7}

	for i, v := range want {
		if v != got[i] {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}
}

func TestStrings(t *testing.T) {
	tests := []struct {
		name string
		got  []string
		want []string
	}{
		{
			name: "simple",
			got:  []string{"G", "B", "A"},
			want: []string{"A", "B", "G"},
		},
		{
			name: "sorts not only by first letter",
			got:  []string{"Go", "Golang", "Gas"},
			want: []string{"Gas", "Go", "Golang"},
		},
		{
			name: "sorts by number in string",
			got:  []string{"Go3", "Go1", "Go2", "Go"},
			want: []string{"Go", "Go1", "Go2", "Go3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Strings(tt.got)
			for i, v := range tt.want {
				if v != tt.got[i] {
					t.Errorf("got %v, wanted %v", tt.got, tt.want)
				}
			}
		})
	}
}

func intSlice(sorted bool) []int {
	rand.Seed(time.Now().Unix())
	res := rand.Perm(100)
	if sorted {
		sort.Ints(res)
	}
	return res
}

func float64Slice(sorted bool) []float64 {
	res := make([]float64, 100)
	for i := range res {
		rand.Seed(time.Now().Unix())
		res[i] = rand.Float64()
	}
	if sorted {
		sort.Float64s(res)
	}
	return res
}

func benchmarkInts(x []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		sort.Ints(x)
	}
}

func benchmarkFloat64s(x []float64, b *testing.B) {
	for n := 0; n < b.N; n++ {
		sort.Float64s(x)
	}
}
func BenchmarkInts(b *testing.B)           { benchmarkInts(intSlice(false), b) }
func BenchmarkIntsSorted(b *testing.B)     { benchmarkInts(intSlice(true), b) }
func BenchmarkFloat64s(b *testing.B)       { benchmarkFloat64s(float64Slice(false), b) }
func BenchmarkFloat64sSorted(b *testing.B) { benchmarkFloat64s(float64Slice(true), b) }
