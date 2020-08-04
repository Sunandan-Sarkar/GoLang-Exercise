package greetings

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s:= Greet("Novi")
	if s!= "Welcome my dear Novi"{
		t.Error("Got",s,"Expected","Welcome my dear Novi")
	}
}
func ExampleGreet() {
	fmt.Println(Greet("Novi"))
	//Output:
	//Welcome my dear Novi
}
func BenchmarkGreet(b *testing.B) {
	for i:=0;i<b.N;i++{
		Greet("Novi")
	}
}
