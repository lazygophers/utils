//go:build (lang_ko || lang_all) && (country_africa || country_all || country_eastern_africa || country_yt)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMayotte.RegisterName(xlanguage.Korean, "마요트")
	dataMayotte.RegisterOfficialName(xlanguage.Korean, "마요트")
	dataMayotte.RegisterCapital(xlanguage.Korean, "마무주")
}
