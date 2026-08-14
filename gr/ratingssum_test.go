package gr

import (
	"encoding/json"
	"testing"
)

// Goodreads computes ratingsSum server-side and returns fractional values for
// some popular works (e.g. 7651285.5 for Eragon, legacy id 113436). With the
// field typed int64 the whole book payload failed to unmarshal and the book
// became invisible to Readarr. See upstream issues #577/#579/#581.
func TestRatingsSumToleratesFractionalValues(t *testing.T) {
	for _, raw := range []string{
		`{"averageRating":4.25,"ratingsCount":1800302,"ratingsSum":7651285.5}`,
		`{"averageRating":4.25,"ratingsCount":1800302,"ratingsSum":7651285}`,
	} {
		var stats BookInfoStatsBookOrWorkStats
		if err := json.Unmarshal([]byte(raw), &stats); err != nil {
			t.Fatalf("unmarshal %s: %v", raw, err)
		}
		if stats.RatingsSum < 7651285 || stats.RatingsSum >= 7651286 {
			t.Fatalf("unexpected RatingsSum %v for %s", stats.RatingsSum, raw)
		}
	}
}
