package trie

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSfxWildCardExactMatch(t *testing.T) {
	rTrie := NewRuneTrie()
	exactMatchList := []string{
		"first.second",
		"first.third",
		"today_is_sunday",
	}
	prefixWildCardList := []string{
		"te.wi.*",
		"eve_*",
		"first.second.third.*",
	}

	validate := func(strList []string) {
		for _, str := range strList {
			if rTrie.Get(str) == nil {
				t.Error("failed to find", str, "in rune trie")
			}
		}
	}

	Convey("testing exact match", t, func() {
		for _, str := range exactMatchList {
			rTrie.Put(str, len(str))
		}
		validate(exactMatchList)

		Convey("testing prefix wild card test", func() {
			for _, str := range prefixWildCardList {
				rTrie.Put(str[:len(str)-1], "*")
			}

			validate(exactMatchList)
			validate(prefixWildCardList)

			validate([]string{
				"te.wi.p99",
				"te.wi.count",
				"eve_is_good",
				"eve_is_working",
				"first.second.third.p99",
			})
		})
	})
}
