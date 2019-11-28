package main

import "testing"

func TestCountFlips_SmallInput1(t *testing.T) {

	pancakeStack := "-"
	numFlips := countFlips(pancakeStack)

	if numFlips != 1 {
		t.Errorf(
			"Number of flips was incorrect.  Expected: %d, but got %d\n",
			1,
			numFlips)
	}

}

func TestCountFlips_SmallInput2(t *testing.T) {

	pancakeStack := "-+"
	numFlips := countFlips(pancakeStack)

	if numFlips != 1 {
		t.Errorf(
			"Number of flips was incorrect.  Expected: %d, but got %d\n",
			1,
			numFlips)
	}

}

func TestCountFlips_SmallInput3(t *testing.T) {

	pancakeStack := "+-"
	numFlips := countFlips(pancakeStack)

	if numFlips != 2 {
		t.Errorf(
			"Number of flips was incorrect.  Expected: %d, but got %d\n",
			2,
			numFlips)
	}

}

func TestCountFlips_SmallInput4(t *testing.T) {

	pancakeStack := "+++"
	numFlips := countFlips(pancakeStack)

	if numFlips != 0 {
		t.Errorf(
			"Number of flips was incorrect.  Expected: %d, but got %d\n",
			0,
			numFlips)
	}

}

func TestCountFlips_SmallInput5(t *testing.T) {

	pancakeStack := "--+-"
	numFlips := countFlips(pancakeStack)

	if numFlips != 3 {
		t.Errorf(
			"Number of flips was incorrect.  Expected: %d, but got %d\n",
			3,
			numFlips)
	}

}

func TestCountFlips_LargeInput(t *testing.T) {

	pancakeStack := "----+++-++++---+----+-++++---+++++-+-+-+++-----+-+--+++++---+-++++-++----++++---+-++++-++----++++-+-"
	numFlips := countFlips(pancakeStack)

	if numFlips != 43 {
		t.Errorf(
			"Number of flips was incorrect.  Expected: %d, but got %d\n",
			3,
			numFlips)
	}

}