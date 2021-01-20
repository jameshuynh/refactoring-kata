package main

// Item presents data in Item
type Item struct {
	name            string
	sellIn, quality int
}

// UpdateQuality updates quality of input items
func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		updateQualityForItem(items[i])
	}
}

func updateQualityForItem(item *Item) {
	switch item.name {
	case "Sulfuras, Hand of Ragnaros":
		updateSufurasItem(item)
	case "Aged Brie":
		updateAgedBrieItem(item)
	case "Backstage passes to a TAFKAL80ETC concert":
		updateBackstageItem(item)
	case "Conjured":
		updateConjuredItem(item)
	default:
		updateOtherItem(item)
	}
}

// Do nothing with Sulfuras item
func updateSufurasItem(item *Item) {
	// nothing here
}

// Do nothing with Sulfuras item
func updateAgedBrieItem(item *Item) {
	item.sellIn = item.sellIn - 1
	item.quality = item.quality + 1
	if item.quality > 50 {
		item.quality = 50
	}
}

// Do nothing with Sulfuras item
func updateBackstageItem(item *Item) {
	if item.sellIn <= 0 {
		item.quality = 0
	} else if item.sellIn <= 5 {
		item.quality = item.quality + 3
	} else if item.sellIn <= 10 {
		item.quality = item.quality + 2
	} else {
		item.quality = item.quality + 1
	}

	item.sellIn = item.sellIn - 1
	if item.quality > 50 {
		item.quality = 50
	}
}

func updateOtherItem(item *Item) {
	if item.quality > 0 {
		item.quality = item.quality - 1
	}

	item.sellIn = item.sellIn - 1
	if item.sellIn < 0 && item.quality > 0 {
		item.quality = item.quality - 1
	}
}

func updateConjuredItem(item *Item) {
	if item.quality > 0 {
		item.quality = item.quality - 2
	}

	item.sellIn = item.sellIn - 1
	if item.sellIn < 0 && item.quality > 0 {
		item.quality = item.quality - 2
	}

	if item.quality < 0 {
		item.quality = 0
	}
}
