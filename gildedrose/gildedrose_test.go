package gildedrose_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_Regular_Product(t *testing.T) {
	t.Run("Regular item decrease by 1 in selling and quality", func(t *testing.T) {
		var items = []*gildedrose.Item{
			{"+5 Dexterity Vest", 5, 4},
		}

		gildedrose.UpdateQuality(items)

		assert.Equal(t, 4, items[0].SellIn)
		assert.Equal(t, 3, items[0].Quality)
	})

	t.Run("quality degrade twice once sell date has passed", func(t *testing.T) {
		var items = []*gildedrose.Item{
			{"+5 Dexterity Vest", -1, 4},
		}

		gildedrose.UpdateQuality(items)

		assert.Equal(t, -2, items[0].SellIn)
		assert.Equal(t, 2, items[0].Quality)
	})

	t.Run("quality of an item is never negative", func(t *testing.T) {
		var items = []*gildedrose.Item{
			{"+5 Dexterity Vest", -1, 0},
		}

		gildedrose.UpdateQuality(items)

		assert.Equal(t, -2, items[0].SellIn)
		assert.Equal(t, 0, items[0].Quality)
	})
}

func Test_Aged_Brie(t *testing.T) {
	t.Run("aged brie quality increases quality by 1", func(t *testing.T) {
		var items = []*gildedrose.Item{{"Aged Brie", 11, 30}}

		gildedrose.UpdateQuality(items)

		assert.Equal(t, "Aged Brie", items[0].Name)
		assert.Equal(t, 10, items[0].SellIn)
		assert.Equal(t, 31, items[0].Quality)
	})

	t.Run("Quality of an item is never more than 50", func(t *testing.T) {
		var items = []*gildedrose.Item{{"Aged Brie", 10, 50}}

		gildedrose.UpdateQuality(items)

		assert.Equal(t, "Aged Brie", items[0].Name)
		assert.Equal(t, 9, items[0].SellIn)
		assert.Equal(t, 50, items[0].Quality)
	})
}

func Test_Sulfuras(t *testing.T) {
	t.Run("Quality of Sulphuras Never Decreases with day and quality", func(t *testing.T) {
		var items = []*gildedrose.Item{{"Sulfuras, Hand of Ragnaros", 1, 80}}

		gildedrose.UpdateQuality(items)

		assert.Equal(t, "Sulfuras, Hand of Ragnaros", items[0].Name)
		assert.Equal(t, 1, items[0].SellIn)
		assert.Equal(t, 80, items[0].Quality)
	})
}

func Test_BackStagePassProducts(t *testing.T) {
	t.Run("Backstage passes to a TAFKAL80ETC concert quality increases quality ", func(t *testing.T) {
		var items = []*gildedrose.Item{{"Backstage passes to a TAFKAL80ETC concert", 11, 20}}

		gildedrose.UpdateQuality(items)

		assert.Equal(t, "Backstage passes to a TAFKAL80ETC concert", items[0].Name)
		assert.Equal(t, 10, items[0].SellIn)
		assert.Equal(t, 21, items[0].Quality)
	})

	t.Run("Backstage passes to a TAFKAL80ETC concert quality by 2 when 10 days", func(t *testing.T) {
		var items = []*gildedrose.Item{{"Backstage passes to a TAFKAL80ETC concert", 10, 30}}

		gildedrose.UpdateQuality(items)

		assert.Equal(t, "Backstage passes to a TAFKAL80ETC concert", items[0].Name)
		assert.Equal(t, 9, items[0].SellIn)
		assert.Equal(t, 32, items[0].Quality)
	})

	t.Run("Backstage passes to a TAFKAL80ETC concert quality by 2 when 6 days ", func(t *testing.T) {
		var items = []*gildedrose.Item{{"Backstage passes to a TAFKAL80ETC concert", 6, 20}}

		gildedrose.UpdateQuality(items)

		assert.Equal(t, "Backstage passes to a TAFKAL80ETC concert", items[0].Name)
		assert.Equal(t, 5, items[0].SellIn)
		assert.Equal(t, 22, items[0].Quality)
	})

	t.Run("Backstage passes to a TAFKAL80ETC concert quality by 3 when 5 days ", func(t *testing.T) {
		var items = []*gildedrose.Item{{"Backstage passes to a TAFKAL80ETC concert", 5, 20}}

		gildedrose.UpdateQuality(items)

		assert.Equal(t, "Backstage passes to a TAFKAL80ETC concert", items[0].Name)
		assert.Equal(t, 4, items[0].SellIn)
		assert.Equal(t, 23, items[0].Quality)
	})

	t.Run("Backstage passes to a TAFKAL80ETC concert quality by 3 when 1 days ", func(t *testing.T) {
		var items = []*gildedrose.Item{{"Backstage passes to a TAFKAL80ETC concert", 1, 20}}

		gildedrose.UpdateQuality(items)

		assert.Equal(t, "Backstage passes to a TAFKAL80ETC concert", items[0].Name)
		assert.Equal(t, 0, items[0].SellIn)
		assert.Equal(t, 23, items[0].Quality)
	})

	t.Run("Backstage passes to a TAFKAL80ETC concert quality can never be more than 50 ", func(t *testing.T) {
		var items = []*gildedrose.Item{{"Backstage passes to a TAFKAL80ETC concert", 1, 50}}

		gildedrose.UpdateQuality(items)

		assert.Equal(t, "Backstage passes to a TAFKAL80ETC concert", items[0].Name)
		assert.Equal(t, 0, items[0].SellIn)
		assert.Equal(t, 50, items[0].Quality)
	})

	t.Run("Backstage passes to a TAFKAL80ETC concert quality drops to 0 after the concert ", func(t *testing.T) {
		var items = []*gildedrose.Item{{"Backstage passes to a TAFKAL80ETC concert", 0, 20}}

		gildedrose.UpdateQuality(items)

		assert.Equal(t, "Backstage passes to a TAFKAL80ETC concert", items[0].Name)
		assert.Equal(t, 0, items[0].Quality)
		assert.Equal(t, -1, items[0].SellIn)
	})
}

func TestConjuredProducts(t *testing.T) {
	t.Run("conjured products quality is reduced twice ", func(t *testing.T) {
		items := []*gildedrose.Item{{"Conjured Mana Cake", 3, 6}}
		gildedrose.UpdateQuality(items)

		assert.Equal(t, "Conjured Mana Cake", items[0].Name)
		assert.Equal(t, 2, items[0].SellIn)
		assert.Equal(t, 4, items[0].Quality)
	})

	t.Run("conjured products quality is never negative ", func(t *testing.T) {
		items := []*gildedrose.Item{{"Conjured Mana Cake", 0, 2}}
		gildedrose.UpdateQuality(items)

		assert.Equal(t, "Conjured Mana Cake", items[0].Name)
		assert.Equal(t, -1, items[0].SellIn)
		assert.Equal(t, 0, items[0].Quality)
	})
}
