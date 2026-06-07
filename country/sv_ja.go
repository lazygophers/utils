//go:build (lang_ja || lang_all) && (country_all || country_americas || country_central_america || country_sv)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataElSalvador.RegisterName(xlanguage.Japanese, "エルサルバドル")
	dataElSalvador.RegisterOfficialName(xlanguage.Japanese, "エルサルバドル共和国")
	dataElSalvador.RegisterCapital(xlanguage.Japanese, "サンサルバドル")
}
