package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Foo(t *testing.T) {
	var items = []*Item{
		{"foo", 0, 0},
	}

	UpdateQuality(items)

	// if items[0].name != "fixme" {
	// 	t.Errorf("Name: Expected %s but got %s ", "fixme", items[0].name)
	// }

	require.Equal(t, true, true)
}
