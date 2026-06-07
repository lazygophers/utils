//go:build (lang_ar || lang_all) && (country_all || country_oceania || country_polynesia || country_tv)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTuvalu.RegisterName(xlanguage.Arabic, "توفالو")
	dataTuvalu.RegisterOfficialName(xlanguage.Arabic, "توفالو")
	dataTuvalu.RegisterCapital(xlanguage.Arabic, "فونافوتي")
}
