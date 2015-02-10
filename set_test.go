package Automation

import "testing"

func TestAdd(t *testing.T) {
	set := NewSet()

	set.Add("word1")
	set.Add("word2")
	set.Add("word3")

	if set.Cardinality() != 3 {
		t.Error("Problem")
	}

	set.Add("word4")

	if set.Cardinality() != 4 {
		t.Error("Problem")
	}

	if set.Contains("maslini") {
		t.Error("This is just wrong!!!")
	}

	if !set.Contains("word1") {
		t.Error("There is problem with contains: word not found!!!")
	}
}

func TestAdd2(t *testing.T) {
	slice := []string{"asd", "dsa", "qwerty", "bvcxz"}
	set := SetFromSlice(slice)

	if set.Cardinality() != 4 {
		t.Error("Problem")
	}

	if set.Contains("maslini") {
		t.Error("This is just wrong!!!")
	}

	if !set.Contains("qwerty") {
		t.Error("There is problem with contains: word not found!!!")
	}

	values := set.Values()
	for _, value := range values {
		if (value != slice[0]) && (value != slice[1]) && (value != slice[2]) && (value != slice[3]) {
			t.Error("Some strange value: " + value)
		}
	}

}

func TestAdd3(t *testing.T) {
	slice := []string{"asd", "dsa", "qwerty", "bvcxz"}
	set1 := SetFromSlice(slice)

	set := NewSet()
	set.Add("baklava")
	set.AddSet(set1)

	if set.Cardinality() != 5 {
		t.Error("Problem")
	}

	if set.Contains("maslini") {
		t.Error("This is just wrong!!!")
	}

	if !set.Contains("qwerty") {
		t.Error("There is problem with contains: word not found!!!")
	}

	values := set.Values()
	for _, value := range values {
		if (value != slice[0]) && (value != slice[1]) && (value != slice[2]) && (value != slice[3]) && (value != "baklava") {
			t.Error("Some strange value: " + value)
		}
	}

	interesction := set.Intersection(set1)
	if interesction.Cardinality() != 4 {
		t.Error("Problem with Intersection!")
		t.Error(interesction)
		t.Error(set)
		t.Error(set1)
	}

	vals := interesction.Values()
	for _, val := range vals {
		if (val != slice[0]) && (val != slice[1]) && (val != slice[2]) && (val != slice[3]) && (val != "baklava") {
			t.Error("Some strange value: " + val)
		}
	}

	if set1.Eq(set) {
		t.Error("Problem with equaliti of sets!")
	}
	if !set.Eq(set) {
		t.Error("Problem with equaliti of sets!")
	}
}
