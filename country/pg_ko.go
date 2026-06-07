//go:build (lang_ko || lang_all) && (country_all || country_melanesia || country_oceania || country_pg)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPapuaNewGuinea.RegisterName(xlanguage.Korean, "파푸아뉴기니")
	dataPapuaNewGuinea.RegisterOfficialName(xlanguage.Korean, "파푸아뉴기니 독립국")
	dataPapuaNewGuinea.RegisterCapital(xlanguage.Korean, "포트모르즈비")
}
