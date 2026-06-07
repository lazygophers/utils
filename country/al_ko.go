//go:build (lang_ko || lang_all) && (country_al || country_all || country_europe || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlbania.RegisterName(xlanguage.Korean, "알바니아")
	dataAlbania.RegisterOfficialName(xlanguage.Korean, "알바니아 공화국")
	dataAlbania.RegisterCapital(xlanguage.Korean, "티라나")
}
