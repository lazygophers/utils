//go:build (lang_ko || lang_all) && (country_all || country_asia || country_bd || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBangladesh.RegisterName(xlanguage.Korean, "방글라데시")
	dataBangladesh.RegisterOfficialName(xlanguage.Korean, "방글라데시 인민 공화국")
	dataBangladesh.RegisterCapital(xlanguage.Korean, "다카")
}
