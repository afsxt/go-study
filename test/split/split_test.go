package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	testGroup := map[string]testCase{
		"case1": testCase{"a:b:c:d", ":", []string{"a", "b", "c", "d"}},
		"case2": testCase{":bc:ed:ff:", ":", []string{"bc", "ed", "ff"}},
		"case3": testCase{"abcdefg", "bc", []string{"a", "defg"}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got:%#v, but want:%#v\n", got, tc.want)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e", ":")
	}
}
