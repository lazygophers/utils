//go:build (lang_fr || lang_all) && (country_all || country_asia || country_bn || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBrunei.RegisterName(xlanguage.French, "Brunei")
	dataBrunei.RegisterOfficialName(xlanguage.French, "Brunei Darussalam")
	dataBrunei.RegisterCapital(xlanguage.French, "Bandar Seri Begawan")
}
