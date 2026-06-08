package fake

import (
	"math/rand/v2"
	"strconv"
	"strings"
	"time"

	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/country"
)

// jpMobilePrefixes are the two-digit mobile prefixes assigned by the Japanese
// Ministry of Internal Affairs and Communications (the leading "0" of the
// canonical 070/080/090 ranges is stripped so the international form
// "+81 70-xxxx-xxxx" composes cleanly).
var jpMobilePrefixes = []string{"70", "80", "90"}

// jpLandlinePrefixes lists representative trunk area codes for major Japanese
// cities, again with the leading domestic "0" trimmed so the international
// "+81" form composes correctly.
var jpLandlinePrefixes = []string{
	"3",  // Tokyo 23 wards
	"6",  // Osaka
	"11", // Sapporo
	"22", // Sendai
	"45", // Yokohama
	"52", // Nagoya
	"75", // Kyoto
	"78", // Kobe
	"82", // Hiroshima
	"92", // Fukuoka
	"98", // Naha (Okinawa)
}

// jpMyNumberWeights1 is the per-digit weighting vector applied to digits
// n1..n6 of the Japanese Individual Number ("My Number") when computing the
// 12th check digit, as published by Japan's National Tax Agency.
var jpMyNumberWeights1 = [6]int{6, 5, 4, 3, 2, 7}

// jpMyNumberWeights2 is the weighting vector applied to digits n7..n11.
var jpMyNumberWeights2 = [5]int{6, 5, 4, 3, 2}

// localeJP registers the Japan (JP) locale skeleton. Localised pools such as
// first / last names, cities and streets are filled in by the companion
// locale data files (e.g. jp_ja.go) during their own init phase.
var localeJP = &Locale{
	Country:        country.Japan,
	OfficialLangs:  []xlanguage.Tag{xlanguage.Japanese},
	PhonePrefixes:  jpMobilePrefixes,
	LandlinePrefix: jpLandlinePrefixes,
	ZipFormat:      "###-####",
	IdCardGen:      genMyNumberJP,
	Streets:        map[xlanguage.Tag][]string{},
	Cities:         map[xlanguage.Tag][]CityEntry{},
	FirstNames:     map[xlanguage.Tag]map[Gender][]string{},
	LastNames:      map[xlanguage.Tag][]string{},
	Domain:         "jp",
}

func init() { register(localeJP) }

// genMyNumberJP generates a 12-digit Japanese Individual Number (マイナンバー)
// whose 12th digit is the official check digit. The algorithm follows the
// National Tax Agency specification:
//
//	sum = Σ(n_i * P_i) for i in 1..11, where
//	  P_1..P_6  = 6,5,4,3,2,7
//	  P_7..P_11 = 6,5,4,3,2
//	check = 11 - (sum mod 11)
//	if check >= 10 → check = 0
//
// gender and birth are unused — My Number carries no demographic semantics.
// When rng is nil the runtime-wide math/rand/v2 source is used.
func genMyNumberJP(rng *rand.Rand, _ Gender, _ time.Time) string {
	digits := [11]int{}
	for i := 0; i < 11; i++ {
		if rng != nil {
			digits[i] = rng.IntN(10)
		} else {
			digits[i] = rand.IntN(10)
		}
	}

	sum := 0
	for i := 0; i < 6; i++ {
		sum += digits[i] * jpMyNumberWeights1[i]
	}
	for i := 0; i < 5; i++ {
		sum += digits[6+i] * jpMyNumberWeights2[i]
	}
	check := 11 - (sum % 11)
	if check >= 10 {
		check = 0
	}

	var b strings.Builder
	b.Grow(12)
	for i := 0; i < 11; i++ {
		b.WriteString(strconv.Itoa(digits[i]))
	}
	b.WriteString(strconv.Itoa(check))
	return b.String()
}
