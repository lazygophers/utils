//go:build (lang_ar || lang_all) && (country_all || country_americas || country_bo || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBolivia.RegisterName(xlanguage.Arabic, "بوليفيا")
	dataBolivia.RegisterOfficialName(xlanguage.Arabic, "دولة بوليفيا متعددة القوميات")
	dataBolivia.RegisterCapital(xlanguage.Arabic, "سوكري")
}
