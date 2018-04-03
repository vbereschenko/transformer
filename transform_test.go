package transformer

import "testing"

func TestTransform(t *testing.T) {
	source := struct {
		Name string
		Age uint
	}{
		Name: "RegularName",
		Age: 14,
	}
	target := struct {
		Name string
		Age uint
	}{
		Name: "Name to be replaced",
		Age: 18,
	}

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

func TestTransform_WithAnnotation(t *testing.T) {
	source := struct {
		AddInfo uint
		Name string
		Age uint
	}{
		Name: "RegularName",
		Age: 14,
	}
	target := struct {
		FirstName string `fromField:"Name"`
		Age uint
	}{
		FirstName: "Name to be replaced",
		Age: 14,
	}

	err := Transform(source, &target)
	if err != nil {
		t.Fatal("error: ", err)
	}
	t.Log("no error")

	if target.FirstName != source.Name {
		t.Fail()
	}
	t.Log("fields equal")

	if target.Age != source.Age {
		t.Fail()
	}

	t.Log("fields equal")
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