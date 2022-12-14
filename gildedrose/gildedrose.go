package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		updateItemQuality(item)
	}
}

const AGEDBRIE = "Aged Brie"
const BACKSTAGEPASSES = "Backstage passes to a TAFKAL80ETC concert"
const SULFURAS = "Sulfuras, Hand of Ragnaros"
const CONJURED = "Conjured Mana Cake"

func updateItemQuality(item *Item) {
	switch item.Name {
	case AGEDBRIE:
		updateAgedBrieQuality(item)
		break
	case BACKSTAGEPASSES:
		updatebackStageQuality(item)
		break
	case SULFURAS:
		break
	case CONJURED:
		updateConjuredQuality(item)
	default:
		updateRegularProductQuality(item)
	}
}

func updateConjuredQuality(item *Item) {
	qualityAdjustment(item, -2)
	item.SellIn = item.SellIn - 1

	if item.SellIn < 0 {
		qualityAdjustment(item, -2)
	}
}

func updateRegularProductQuality(item *Item) {
	qualityAdjustment(item, -1)
	item.SellIn = item.SellIn - 1

	if item.SellIn < 0 {
		qualityAdjustment(item, -1)
	}
}

func updatebackStageQuality(item *Item) {
	qualityAdjustment(item, 1)
	if item.SellIn < 11 {
		qualityAdjustment(item, 1)
	}
	if item.SellIn < 6 {
		qualityAdjustment(item, 1)
	}
	item.SellIn = item.SellIn - 1
	if item.SellIn < 0 {
		item.Quality = 0
	}
}

func updateAgedBrieQuality(item *Item) {
	qualityAdjustment(item, 1)
	item.SellIn = item.SellIn - 1
	if item.SellIn < 0 {
		qualityAdjustment(item, 1)
	}
}

func qualityAdjustment(item *Item, adjustment int) {
	newQuality := item.Quality + adjustment
	if item.Quality > 0 && item.Quality < 50 {
		item.Quality = newQuality
	}
}
