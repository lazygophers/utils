//go:build (lang_ja || lang_all) && (country_all || country_americas || country_south_america || country_uy)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUruguay.RegisterName(xlanguage.Japanese, "ウルグアイ")
	dataUruguay.RegisterOfficialName(xlanguage.Japanese, "ウルグアイ東方共和国")
	dataUruguay.RegisterCapital(xlanguage.Japanese, "モンテビデオ")
}
