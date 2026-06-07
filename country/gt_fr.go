//go:build (lang_fr || lang_all) && (country_all || country_americas || country_central_america || country_gt)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuatemala.RegisterName(xlanguage.French, "Guatemala")
	dataGuatemala.RegisterOfficialName(xlanguage.French, "République du Guatemala")
	dataGuatemala.RegisterCapital(xlanguage.French, "Guatemala")
}
