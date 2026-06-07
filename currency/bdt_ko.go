//go:build (lang_ko || lang_all) && (country_all || country_asia || country_bd || country_southern_asia || currency_all || currency_bdt)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	BDT.RegisterName(xlanguage.Korean, "방글라데시 타카")
}
