package ProtoType

import "testing"

func TestClone(t *testing.T) {
	ShirtsCache := GetShirtsCloner()
	if ShirtsCache == nil {
		t.Fatal("Received cache was nil")
	}
	item1, err := ShirtsCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}
	if item1 == whitePrototype {
		t.Error("item1 cannot be equal to the white prototype")
	}
	shirt1, ok := item1.(*Shirt)
	if !ok {
		t.Fatal("Type assertion for shirt1 couldn't be done successfully")
	}
	shirt1.SKU = "abbcc"
	item2, err := ShirtsCache.GetClone(White)
	if err != nil {
		t.Fatal(err)
	}
	shirt2, ok := item2.(*Shirt)
	if !ok {
		t.Fatal("Type assertion for shirt1 couldn't be done successfully")
	}
	if shirt1.SKU == shirt2.SKU {
		t.Error("SKU's of shirt1 and shirt2 must be different")
	}
	if shirt1 == shirt2 {
		t.Error("Shirt 1 cannot be equal to shirt 2")
	}
	t.Logf("Log: %s", shirt1.GetInfo())
	t.Logf("Log: %s", shirt2.GetInfo())
	t.Logf("Log: the memory positions of the shirts are different %p != %p \n", &shirt1, &shirt2)
}
