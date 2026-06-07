//go:build (lang_ja || lang_all) && (country_all || country_melanesia || country_oceania || country_vu)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVanuatu.RegisterName(xlanguage.Japanese, "バヌアツ")
	dataVanuatu.RegisterOfficialName(xlanguage.Japanese, "バヌアツ共和国")
	dataVanuatu.RegisterCapital(xlanguage.Japanese, "ポートビラ")
}
