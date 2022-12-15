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
		item.updateAgedBrieQuality()
		break
	case BACKSTAGEPASSES:
		item.updatebackStageQuality()
		break
	case SULFURAS:
		break
	case CONJURED:
		item.updateConjuredQuality()
	default:
		item.updateRegularProductQuality()
	}
}

func (item *Item) updateConjuredQuality() {
	qualityAdjustment(item, -2)
	item.SellIn = item.SellIn - 1

	if item.SellIn < 0 {
		qualityAdjustment(item, -2)
	}
}

func (item *Item) updateRegularProductQuality() {
	qualityAdjustment(item, -1)
	item.SellIn = item.SellIn - 1

	if item.SellIn < 0 {
		qualityAdjustment(item, -1)
	}
}

func (item *Item) updatebackStageQuality() {
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

func (item *Item) updateAgedBrieQuality() {
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
