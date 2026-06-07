//go:build (lang_ja || lang_all) && (country_all || country_oceania || country_polynesia || country_tv)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTuvalu.RegisterName(xlanguage.Japanese, "ツバル")
	dataTuvalu.RegisterOfficialName(xlanguage.Japanese, "ツバル")
	dataTuvalu.RegisterCapital(xlanguage.Japanese, "フナフティ")
}
