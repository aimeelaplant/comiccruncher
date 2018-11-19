package comic_test

import (
	"github.com/aimeelaplant/comiccruncher/comic"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCharacterSlugs(t *testing.T) {
	slugs := []string{"a", "b", "c"}
	assert.Len(t, comic.NewCharacterSlugs(slugs...), 3)

	slugs1 := []string{}
	assert.Len(t, comic.NewCharacterSlugs(slugs1...), 0)
}

func TestAppearanceTypeHasAll(t *testing.T) {
	c := comic.AppearanceType(comic.Main)
	assert.True(t, c.HasAll(comic.Main))
	assert.False(t, c.HasAll(comic.Alternate))
	assert.False(t, c.HasAll(comic.Main|comic.Alternate))

	c = comic.AppearanceType(comic.Main | comic.Alternate)
	assert.True(t, c.HasAll(comic.Main|comic.Alternate))
	assert.True(t, c.HasAll(comic.Alternate))
	assert.True(t, c.HasAll(comic.Main))
}

func TestAppearanceTypeHasAny(t *testing.T) {
	c := comic.AppearanceType(comic.Main)
	assert.True(t, c.HasAny(comic.Main|comic.Alternate))
	assert.True(t, c.HasAny(comic.Main))
	assert.False(t, c.HasAny(comic.Alternate))

	c = comic.AppearanceType(comic.Main | comic.Alternate)
	assert.True(t, c.HasAny(comic.Main|comic.Alternate))
	assert.True(t, c.HasAny(comic.Main))
	assert.True(t, c.HasAny(comic.Alternate))
}

func TestAppearanceTypeMarshalJSON(t *testing.T) {
	main := comic.AppearanceType(comic.Main)
	bytes, err := main.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, "\"main\"", string(bytes))

	alternate := comic.AppearanceType(comic.Alternate)
	bytes, err = alternate.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, "\"alternate\"", string(bytes))

	all := comic.AppearanceType(comic.Main | comic.Alternate)
	bytes, err = all.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, "\"all\"", string(bytes))
}

func TestExpandedCharacterMarshalJSON(t *testing.T) {
	c := comic.NewCharacter("emma frost", 1, comic.VendorTypeCb, "123")
	c.VendorURL = "https://example.com"
	c.Slug = "emma-frost"
	stats := []comic.CharacterStats{
		{Category: comic.AllTimeStats, IssueCount: 1, IssueCountRank: 1, AverageRank: 1, Average: 1},
	}
	aggs := []comic.YearlyAggregate{
		{Count: 10, Year: 1900},
	}
	apps := []comic.AppearancesByYears{
		{CharacterSlug: c.Slug, Category: comic.Main, Aggregates: aggs},
	}
	ec := comic.ExpandedCharacter{
		Character:   c,
		Stats:       stats,
		Appearances: apps,
	}
	b, err := ec.MarshalJSON()
	assert.Nil(t, err)
	s := `{"publisher":{"name":"","slug":""},"name":"emma frost","other_name":"","description":"","slug":"emma-frost","vendor_url":"https://example.com","vendor_description":"","image":"","vendor_image":"","last_syncs":null,"stats":[{"category":"all_time","issue_count_rank":1,"issue_count":1,"average_issues_per_year":1,"average_issues_per_year_rank":1}],"appearances":[{"slug":"emma-frost","category":"main","aggregates":[{"year":1900,"count":10}]}]}`
	assert.Equal(t, s, string(b))
}
