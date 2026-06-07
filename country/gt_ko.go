//go:build (lang_ko || lang_all) && (country_all || country_americas || country_central_america || country_gt)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuatemala.RegisterName(xlanguage.Korean, "과테말라")
	dataGuatemala.RegisterOfficialName(xlanguage.Korean, "과테말라 공화국")
	dataGuatemala.RegisterCapital(xlanguage.Korean, "과테말라시티")
}
