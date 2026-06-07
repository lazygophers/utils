//go:build (lang_ar || lang_all) && (country_all || country_europe || country_gr || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGreece.RegisterName(xlanguage.Arabic, "اليونان")
	dataGreece.RegisterOfficialName(xlanguage.Arabic, "الجمهورية الهيلينية")
	dataGreece.RegisterCapital(xlanguage.Arabic, "أثينا")
}
