package sliceutil

import (
	"testing"

	"github.com/frozzare/go-assert"
)

type Person struct {
	Name string
}

func TestIndexOf(t *testing.T) {
	actual := IndexOf([]string{"hello"}, "world")
	assert.Equal(t, -1, actual)

	actual = IndexOf([]int{1, 2, 3}, 2)
	assert.Equal(t, 1, actual)

	actual = IndexOf([]Person{Person{Name: "Elli"}, Person{Name: "Fredrik"}}, Person{Name: "Fredrik"})
	assert.Equal(t, 1, actual)
}

func TestHas(t *testing.T) {
	actual := Has([]string{"hello"}, "world")
	assert.False(t, actual)

	actual = Has([]int{1, 2, 3}, 2)
	assert.True(t, actual)

	actual = Has([]Person{Person{Name: "Elli"}, Person{Name: "Fredrik"}}, Person{Name: "Fredrik"})
	assert.True(t, actual)
}
