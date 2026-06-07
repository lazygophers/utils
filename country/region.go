package country

import (
	"sync"

	xlanguage "golang.org/x/text/language"
)

// Region is a UN M.49 geographic classification — continent + region + sub-region.
// Countries reference one of the singleton instances below; equality compares by
// pointer. Subregion labels support multi-language via [Region.RegisterName].
type Region struct {
	continent string
	region    string
	subregion string

	namesMu sync.RWMutex
	names   map[xlanguage.Tag]string
}

// Continent returns the two-letter continent code ("AS", "EU", "AF", "NA",
// "SA", "OC", "AN").
func (r *Region) Continent() string {
	if r == nil {
		return ""
	}
	return r.continent
}

// RegionName returns the UN M.49 region (e.g. "Asia", "Europe", "Americas").
func (r *Region) RegionName() string {
	if r == nil {
		return ""
	}
	return r.region
}

// Subregion returns the English UN M.49 sub-region label (e.g. "Eastern Asia").
func (r *Region) Subregion() string {
	if r == nil {
		return ""
	}
	return r.subregion
}

// Name returns the localized sub-region name in the current goroutine's language.
func (r *Region) Name() string {
	if r == nil {
		return ""
	}
	return r.NameIn(currentTag())
}

// NameIn returns the localized sub-region name in the given language, falling
// back to language base, then English (the canonical [Region.Subregion] label).
func (r *Region) NameIn(tag xlanguage.Tag) string {
	if r == nil {
		return ""
	}
	r.namesMu.RLock()
	defer r.namesMu.RUnlock()
	if v, ok := r.names[tag]; ok {
		return v
	}
	base, _ := tag.Base()
	baseTag := xlanguage.Make(base.String())
	if v, ok := r.names[baseTag]; ok {
		return v
	}
	if v, ok := r.names[xlanguage.English]; ok {
		return v
	}
	return r.subregion
}

// RegisterName registers a localized sub-region name. Intended for use in
// region_<lang>.go init() functions.
func (r *Region) RegisterName(tag xlanguage.Tag, name string) {
	r.namesMu.Lock()
	if r.names == nil {
		r.names = make(map[xlanguage.Tag]string)
	}
	r.names[tag] = name
	r.namesMu.Unlock()
}

// String returns the English sub-region label, satisfying fmt.Stringer.
func (r *Region) String() string {
	if r == nil {
		return ""
	}
	return r.subregion
}

func newRegion(continent, region, subregion string) *Region {
	return &Region{
		continent: continent,
		region:    region,
		subregion: subregion,
		names:     make(map[xlanguage.Tag]string),
	}
}

// UN M.49 region singletons. Countries reference these by pointer.
var (
	RegionEasternAsia      = newRegion("AS", "Asia", "Eastern Asia")
	RegionSouthEasternAsia = newRegion("AS", "Asia", "South-eastern Asia")
	RegionSouthernAsia     = newRegion("AS", "Asia", "Southern Asia")
	RegionWesternAsia      = newRegion("AS", "Asia", "Western Asia")
	RegionCentralAsia      = newRegion("AS", "Asia", "Central Asia")

	RegionEasternEurope  = newRegion("EU", "Europe", "Eastern Europe")
	RegionNorthernEurope = newRegion("EU", "Europe", "Northern Europe")
	RegionSouthernEurope = newRegion("EU", "Europe", "Southern Europe")
	RegionWesternEurope  = newRegion("EU", "Europe", "Western Europe")

	RegionNorthernAfrica = newRegion("AF", "Africa", "Northern Africa")
	RegionEasternAfrica  = newRegion("AF", "Africa", "Eastern Africa")
	RegionMiddleAfrica   = newRegion("AF", "Africa", "Middle Africa")
	RegionSouthernAfrica = newRegion("AF", "Africa", "Southern Africa")
	RegionWesternAfrica  = newRegion("AF", "Africa", "Western Africa")

	RegionNorthernAmerica = newRegion("NA", "Americas", "Northern America")
	RegionCentralAmerica  = newRegion("NA", "Americas", "Central America")
	RegionSouthAmerica    = newRegion("SA", "Americas", "South America")
	RegionCaribbean       = newRegion("NA", "Americas", "Caribbean")

	RegionAustraliaAndNewZealand = newRegion("OC", "Oceania", "Australia and New Zealand")
	RegionMelanesia              = newRegion("OC", "Oceania", "Melanesia")
	RegionMicronesia             = newRegion("OC", "Oceania", "Micronesia")
	RegionPolynesia              = newRegion("OC", "Oceania", "Polynesia")

	RegionAntarctic = newRegion("AN", "Antarctic", "Antarctic")
)
