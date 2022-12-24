package iteration
import "testing"

func TestRepeat(t *testing.T){
	got := Repeat("d", 4)
	want := "dddd"

	if got != want{
		t.Errorf("want %q, got %q", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i:=0; i<b.N; i++ {
		Repeat("a", 4)
	}
}