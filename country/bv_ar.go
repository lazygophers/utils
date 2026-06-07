//go:build (lang_ar || lang_all) && (country_all || country_antarctic || country_bv)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBouvetIsland.RegisterName(xlanguage.Arabic, "جزيرة بوفيه")
	dataBouvetIsland.RegisterOfficialName(xlanguage.Arabic, "جزيرة بوفيه")
	dataBouvetIsland.RegisterCapital(xlanguage.Arabic, "—")
}
