package transformer

import "testing"

func TestTransform(t *testing.T) {
	source := struct {
		Name string
		Age uint
	}{}
	target := struct {
		Name string
		Age uint
	}{}

	err := Transform(source, &target)
	if err != nil {
		t.Fatal("error: ", err)
	}

	if target.Name != source.Name {
		t.Fail()
	}

	if target.Age != source.Age {
		t.Fail()
	}
}

func TestTransform_Pointer(t *testing.T) {
	source := struct {
		Name string
		Age uint
	}{}
	target := struct {
		Name string
		Age uint
	}{}

	err := Transform(&source, &target)
	if err != nil {
		t.Fatal("error: ", err)
	}

	if target.Name != source.Name {
		t.Fail()
	}

	if target.Age != source.Age {
		t.Fail()
	}
}

func TestTransform_Invalid(t *testing.T) {
	source := struct {
		Name string
		Age uint
	}{}
	target := struct {
		Name string
		Age uint
	}{}

	err := Transform(source, target)
	if err == nil {
		t.Fail()
	}
}