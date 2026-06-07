//go:build (lang_ja || lang_all) && (country_africa || country_all || country_eastern_africa || country_ke)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKenya.RegisterName(xlanguage.Japanese, "ケニア")
	dataKenya.RegisterOfficialName(xlanguage.Japanese, "ケニア共和国")
	dataKenya.RegisterCapital(xlanguage.Japanese, "ナイロビ")
}
