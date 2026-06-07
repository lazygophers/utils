//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUruguay.RegisterName(xlanguage.Japanese, "ウルグアイ")
	dataUruguay.RegisterOfficialName(xlanguage.Japanese, "ウルグアイ東方共和国")
	dataUruguay.RegisterCapital(xlanguage.Japanese, "モンテビデオ")
}
