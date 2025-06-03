package main

import (
	"testing"
)

func TestRenderGridAuto4x4(t *testing.T) {
	expectedGrid := `┌───┬───┬───┬───┐
│123│abc│456│wnf│
│456│def│dsf│pba│
├───┼───┼───┼───┤
│klm│tuv│nsl│9jf│
│nop│wxy│pql│0sk│
├───┼───┼───┼───┤
│hvm│zuf│1jf│5jf│
│anp│bvh│lam│29j│
├───┼───┼───┼───┤
│9fn│83j│afd│3jf│
│4og│kme│38j│59j│
└───┴───┴───┴───┘
`
	dataGrid := [][][]string{
		{{"123", "456"}, {"abc", "def"}, {"456", "dsf"}, {"wnf", "pba"}},
		{{"klm", "nop"}, {"tuv", "wxy"}, {"nsl", "pql"}, {"9jf", "0sk"}},
		{{"hvm", "anp"}, {"zuf", "bvh"}, {"1jf", "lam"}, {"5jf", "29j"}},
		{{"9fn", "4og"}, {"83j", "kme"}, {"afd", "38j"}, {"3jf", "59j"}},
	}
	grid := printGrid(RenderGridAuto(dataGrid))
	if expectedGrid != grid {
		t.Fatalf("\nexpected:\n%q\nresult:\n%q", expectedGrid, grid)
	}
}

func TestRenderGrid3x3(t *testing.T) {
	expectedGrid := `┌───┬───┬───┐
│123│abc│456│
│456│def│dsf│
├───┼───┼───┤
│tuv│nsl│9jf│
│wxy│pql│0sk│
├───┼───┼───┤
│zuf│1jf│5jf│
│bvh│lam│29j│
└───┴───┴───┘
`
	dataGrid := [][][]string{
		{{"123", "456"}, {"abc", "def"}, {"456", "dsf"}},
		{{"tuv", "wxy"}, {"nsl", "pql"}, {"9jf", "0sk"}},
		{{"zuf", "bvh"}, {"1jf", "lam"}, {"5jf", "29j"}},
	}
	grid := printGrid(RenderGridAuto(dataGrid))
	if expectedGrid != grid {
		t.Fatalf("\nexpected:\n%q\nresult:\n%q", expectedGrid, grid)
	}
}

func TestRenderGrid2x2(t *testing.T) {
	expectedGrid := `┌───┬───┐
│123│456│
│456│123│
├───┼───┤
│klm│fjn│
│fjn│klm│
└───┴───┘
`
	dataGrid := [][][]string{
		{{"123", "456"}, {"456", "123"}},
		{{"klm", "fjn"}, {"fjn", "klm"}},
	}

	grid := printGrid(RenderGridAuto(dataGrid))
	if expectedGrid != grid {
		t.Fatalf("\nexpected:\n%q\nresult:\n%q", expectedGrid, grid)
	}
}
