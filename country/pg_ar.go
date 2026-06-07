//go:build (lang_ar || lang_all) && (country_all || country_melanesia || country_oceania || country_pg)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPapuaNewGuinea.RegisterName(xlanguage.Arabic, "بابوا غينيا الجديدة")
	dataPapuaNewGuinea.RegisterOfficialName(xlanguage.Arabic, "دولة بابوا غينيا الجديدة المستقلة")
	dataPapuaNewGuinea.RegisterCapital(xlanguage.Arabic, "بورت مورسبي")
}
