//go:build (lang_ar || lang_all) && (country_all || country_americas || country_caribbean || country_ms)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMontserrat.RegisterName(xlanguage.Arabic, "مونتسرات")
	dataMontserrat.RegisterOfficialName(xlanguage.Arabic, "إقليم مونتسرات البريطاني فيما وراء البحار")
	dataMontserrat.RegisterCapital(xlanguage.Arabic, "بليموث")
}
