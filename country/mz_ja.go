//go:build (lang_ja || lang_all) && (country_africa || country_all || country_eastern_africa || country_mz)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMozambique.RegisterName(xlanguage.Japanese, "モザンビーク")
	dataMozambique.RegisterOfficialName(xlanguage.Japanese, "モザンビーク共和国")
	dataMozambique.RegisterCapital(xlanguage.Japanese, "マプト")
}
