package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

// - All items have a Quality value which denotes how valuable the item is
// - At the end of each day our system lowers both values for every item
// - Once the sell by date has passed, Quality degrades twice as fast
// - The Quality of an item is never negative
func TestOtherItems(t *testing.T) {
	testcases := []struct {
		items           []*Item
		expectedSellIn  int
		expectedQuality int
	}{
		{
			items: []*Item{
				{
					"Other",
					5,
					20,
				},
			},
			expectedSellIn:  4,
			expectedQuality: 19,
		},
		{
			items: []*Item{
				{
					"Other",
					0,
					20,
				},
			},
			expectedSellIn:  -1,
			expectedQuality: 18,
		},
		{
			items: []*Item{
				{
					"Other",
					-1,
					1,
				},
			},
			expectedSellIn:  -2,
			expectedQuality: 0,
		},
		{
			items: []*Item{
				{
					"Other",
					-1,
					10,
				},
			},
			expectedSellIn:  -2,
			expectedQuality: 8,
		},
	}

	for i, tc := range testcases {
		tc := tc
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			UpdateQuality(tc.items)
			require.Equal(t, tc.expectedQuality, tc.items[0].quality)
			require.Equal(t, tc.expectedSellIn, tc.items[0].sellIn)
		})
	}
}

func TestConjuredItems(t *testing.T) {
	testcases := []struct {
		items           []*Item
		expectedSellIn  int
		expectedQuality int
	}{
		{
			items: []*Item{
				{
					"Conjured",
					5,
					20,
				},
			},
			expectedSellIn:  4,
			expectedQuality: 18,
		},
		{
			items: []*Item{
				{
					"Conjured",
					0,
					20,
				},
			},
			expectedSellIn:  -1,
			expectedQuality: 16,
		},
		{
			items: []*Item{
				{
					"Conjured",
					-1,
					1,
				},
			},
			expectedSellIn:  -2,
			expectedQuality: 0,
		},
		{
			items: []*Item{
				{
					"Conjured",
					-1,
					10,
				},
			},
			expectedSellIn:  -2,
			expectedQuality: 6,
		},
	}

	for i, tc := range testcases {
		tc := tc
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			UpdateQuality(tc.items)
			require.Equal(t, tc.expectedQuality, tc.items[0].quality)
			require.Equal(t, tc.expectedSellIn, tc.items[0].sellIn)
		})
	}
}

// "Aged Brie" actually increases in Quality the older it gets
// The Quality of an item is never more than 50
func TestAgedBrie(t *testing.T) {
	testcases := []struct {
		items           []*Item
		expectedSellIn  int
		expectedQuality int
	}{
		{
			items: []*Item{
				{
					"Aged Brie",
					5,
					20,
				},
			},
			expectedSellIn:  4,
			expectedQuality: 21,
		},
		{
			items: []*Item{
				{
					"Aged Brie",
					10,
					20,
				},
			},
			expectedSellIn:  9,
			expectedQuality: 21,
		},
		{
			items: []*Item{
				{
					"Aged Brie",
					20,
					20,
				},
			},
			expectedSellIn:  19,
			expectedQuality: 21,
		},
		{
			items: []*Item{
				{
					"Aged Brie",
					20,
					50,
				},
			},
			expectedSellIn:  19,
			expectedQuality: 50,
		},
	}

	for i, tc := range testcases {
		tc := tc
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			UpdateQuality(tc.items)
			require.Equal(t, tc.expectedQuality, tc.items[0].quality)
			require.Equal(t, tc.expectedSellIn, tc.items[0].sellIn)
		})
	}
}

// "Sulfuras"
// is a legendary item and as such its Quality is 80 and it never alters.
func TestSulfuras(t *testing.T) {
	testcases := []struct {
		items           []*Item
		expectedSellIn  int
		expectedQuality int
	}{
		{
			items: []*Item{
				{
					"Sulfuras, Hand of Ragnaros",
					5,
					80,
				},
			},
			expectedSellIn:  5,
			expectedQuality: 80,
		},
		{
			items: []*Item{
				{
					"Sulfuras, Hand of Ragnaros",
					10,
					80,
				},
			},
			expectedSellIn:  10,
			expectedQuality: 80,
		},
		{
			items: []*Item{
				{
					"Sulfuras, Hand of Ragnaros",
					20,
					80,
				},
			},
			expectedSellIn:  20,
			expectedQuality: 80,
		},
	}

	for i, tc := range testcases {
		tc := tc
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			UpdateQuality(tc.items)
			require.Equal(t, tc.expectedQuality, tc.items[0].quality)
			require.Equal(t, tc.expectedSellIn, tc.items[0].sellIn)
		})
	}
}

// "Backstage passes", like aged brie, increases in Quality as its SellIn value approaches;
//	Quality increases by 2 when there are 10 days or less and by 3 when there are 5 days or less but
//	Quality drops to 0 after the concert
func TestBackStagePasses(t *testing.T) {
	testcases := []struct {
		items           []*Item
		expectedSellIn  int
		expectedQuality int
	}{
		{
			items: []*Item{
				{
					"Backstage passes to a TAFKAL80ETC concert",
					5,
					20,
				},
			},
			expectedSellIn:  4,
			expectedQuality: 23,
		},
		{
			items: []*Item{
				{
					"Backstage passes to a TAFKAL80ETC concert",
					10,
					20,
				},
			},
			expectedSellIn:  9,
			expectedQuality: 22,
		},
		{
			items: []*Item{
				{
					"Backstage passes to a TAFKAL80ETC concert",
					20,
					20,
				},
			},
			expectedSellIn:  19,
			expectedQuality: 21,
		},
		{
			items: []*Item{
				{
					"Backstage passes to a TAFKAL80ETC concert",
					0,
					20,
				},
			},
			expectedSellIn:  -1,
			expectedQuality: 0,
		},
	}

	for i, tc := range testcases {
		tc := tc
		t.Run(strconv.Itoa(i), func(t *testing.T) {

			UpdateQuality(tc.items)
			require.Equal(t, tc.expectedQuality, tc.items[0].quality)
			require.Equal(t, tc.expectedSellIn, tc.items[0].sellIn)
		})
	}
}
