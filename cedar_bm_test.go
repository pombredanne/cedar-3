package cedar

import (
	"testing"

	"github.com/vcaesar/tt"
)

func init() {
	InitCd()
}

func BenchmarkInsert(t *testing.B) {
	fn := func() {
		cd.Insert([]byte("a"), 1)
		cd.Insert([]byte("b"), 3)
		cd.Insert([]byte("d"), 6)
	}

	tt.BM(t, fn)
}

func BenchmarkJump(t *testing.B) {
	fn := func() {
		cd.Jump([]byte("a"), 1)
	}

	tt.BM(t, fn)
}

func BenchmarkFind(t *testing.B) {
	fn := func() {
		cd.Find([]byte("a"), 1)
	}

	tt.BM(t, fn)
}

func BenchmarkValue(t *testing.B) {
	fn := func() {
		cd.Value(1)
	}

	tt.BM(t, fn)
}

func BenchmarkUpdate(t *testing.B) {
	fn := func() {
		cd.Update([]byte("a"), 1)
	}

	tt.BM(t, fn)
}

func BenchmarkDelete(t *testing.B) {
	fn := func() {
		cd.Delete([]byte("b"))
	}

	tt.BM(t, fn)
}
