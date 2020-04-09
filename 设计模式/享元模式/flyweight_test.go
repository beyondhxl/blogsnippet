package flyweight

import (
	"testing"
)

func ExampleFlyweight() {
	odisp := NewObjDisplay("obj1")
	odisp.Display()
	// Output:
	// Display: data obj1
}

func TestFlyweight(t *testing.T) {
	odisp1 := NewObjDisplay("obj1")
	odisp2 := NewObjDisplay("obj2")
	odisp3 := NewObjDisplay("obj1")

	if odisp1.ObjFlyweight != odisp2.ObjFlyweight {
		t.Fail()
	}

	if odisp3.ObjFlyweight == odisp1.ObjFlyweight {
		t.Log("Pass")
	}
}
