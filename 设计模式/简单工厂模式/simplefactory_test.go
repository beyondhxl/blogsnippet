package simplefactory

import "testing"

func TestHiAPI_Say(t *testing.T) {
	api := NewAPI(3)
	if api == nil {
		t.Fatal("TestHiAPI_Say test fail NewAPI invalid param")
	}

	str := api.Say("Mike")
	if str != "hi, Mike" {
		t.Fatal("TestHiAPI_Say test fail")
	}
}

func TestHelloAPI_Say(t *testing.T) {
	api := NewAPI(2)
	if api == nil {
		t.Fatal("TestHiAPI_Say test fail NewAPI invalid param")
	}

	str := api.Say("Tom")
	if str != "hello, Tom" {
		t.Fatal("TestHelloAPI_Say test fail")
	}
}
