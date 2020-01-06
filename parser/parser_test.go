package parser

import "testing"

func TestParser(t *testing.T) {

	t.Run("Parser - simple object", func(t *testing.T) {
		input := []interface{}{
			"{", "A", ":", 1, "}",
		}

		resultObject := new(SimpleObject)

		err := Parse(input, resultObject)
		if err != nil {
			t.Errorf("Cannot parse object %t", err)
		}

		if resultObject.A != 1 {
			t.Errorf("Object was not parsed correctly!")
		}
	})

	t.Run("Parser - simple array", func(t *testing.T) {
		input := []interface{}{
			"[", 10, 20, 30, "]",
		}

		resultObject := make([]int, 0)

		err := Parse(input, &resultObject)
		if err != nil {
			t.Errorf("Cannot parse object %t", err)
		}

		if resultObject[0] != 10 || resultObject[1] != 20 || resultObject[2] != 30 {
			t.Errorf("Object was not parsed correctly!")
		}
	})

	t.Run("Parser - array of simple arrays", func(t *testing.T) {
		input := []interface{}{
			"[",
			"[", 10, 20, 30, "]",
			"[", 40, 50, 60, "]",
			"[", 70, 80, 90, "]",
			"]",
		}

		resultObject := make([][]int, 0)

		err := Parse(input, &resultObject)
		if err != nil {
			t.Errorf("Cannot parse object %t", err)
		}

		if resultObject[0][0] != 10 || resultObject[0][1] != 20 || resultObject[0][2] != 30 {
			t.Errorf("Object was not parsed correctly!")
		}
	})

	t.Run("Parser - array of objects", func(t *testing.T) {
		input := []interface{}{
			"[",
			"{", "A", ":", 11, "}",
			"{", "A", ":", 22, "}",
			"{", "A", ":", 33, "}",
			"]",
		}

		resultObject := make([]SimpleObject, 0)

		err := Parse(input, &resultObject)
		if err != nil {
			t.Errorf("Cannot parse object %t", err)
		}

		if resultObject[0].A != 11 || resultObject[1].A != 22 || resultObject[2].A != 33 {
			t.Errorf("Object was not parsed correctly!")
		}
	})
}

type SimpleObject struct {
	A int
}
