//go:build (lang_ko || lang_all) && (country_africa || country_all || country_ci || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIvoryCoast.RegisterName(xlanguage.Korean, "코트디부아르")
	dataIvoryCoast.RegisterOfficialName(xlanguage.Korean, "코트디부아르 공화국")
	dataIvoryCoast.RegisterCapital(xlanguage.Korean, "야무수크로")
}
