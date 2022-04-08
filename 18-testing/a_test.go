
package main

import "testing"

func TestAbc(t *testing.T) {
	a := 1

	if a != 2 {

		t.Error("Exepected 1+1 = 2, but getting ", a)

	}

}
