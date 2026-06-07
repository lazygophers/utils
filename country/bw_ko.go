//go:build (lang_ko || lang_all) && (country_africa || country_all || country_bw || country_southern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBotswana.RegisterName(xlanguage.Korean, "보츠와나")
	dataBotswana.RegisterOfficialName(xlanguage.Korean, "보츠와나 공화국")
	dataBotswana.RegisterCapital(xlanguage.Korean, "가보로네")
}
