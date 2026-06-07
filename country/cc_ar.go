//go:build (lang_ar || lang_all) && (country_all || country_australia_and_new_zealand || country_cc || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCocosKeelingIslands.RegisterName(xlanguage.Arabic, "جزر كوكوس")
	dataCocosKeelingIslands.RegisterOfficialName(xlanguage.Arabic, "إقليم جزر كوكوس (كيلينغ)")
	dataCocosKeelingIslands.RegisterCapital(xlanguage.Arabic, "ويست آيلاند")
}
