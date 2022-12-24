

package main
import "testing"
// func TestHello(t *testing.T){
// 	got := Hello("Messi")
// 	want := "Hello, Messi"

// 	if got != want{
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

func TestHello(t *testing.T){
	//make a assertCorrectMsg function which would take three parameters: t, got, want, this function would then check 
	// got!=want and use t.Errorf() to print the error message. 
	 assertCorrectMsg := func(t testing.TB, got, want string){
		
		t.Helper()
		if got!=want{
			t.Errorf("got %q want %q", got, want)
		}
	} 
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Messi", "")                  //no language specified, defaults to English
		want := "Hello, Messi"
		assertCorrectMsg(t, got, want)
	})

	t.Run("saying hello in Spanish", func(t *testing.T) {
		got := Hello("Messi", "Spanish")
		want := "Hola, Messi"
		assertCorrectMsg(t, got, want)})

	t.Run("saying hello in French", func(t *testing.T) {
		got := Hello("Messi", "French")
		want := "Bonjour, Messi"
		assertCorrectMsg(t, got, want)
	})
	t.Run("saying hello world when there is no name", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMsg(t, got, want)
	})
} 