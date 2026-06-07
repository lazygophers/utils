//go:build (lang_ko || lang_all) && (country_all || country_europe || country_im || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIsleOfMan.RegisterName(xlanguage.Korean, "맨섬")
	dataIsleOfMan.RegisterOfficialName(xlanguage.Korean, "맨섬")
	dataIsleOfMan.RegisterCapital(xlanguage.Korean, "더글러스")
}
