//go:build (lang_ar || lang_all) && (country_all || country_americas || country_northern_america || country_pm)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintPierreAndMiquelon.RegisterName(xlanguage.Arabic, "سان بيير وميكلون")
	dataSaintPierreAndMiquelon.RegisterOfficialName(xlanguage.Arabic, "جماعة سان بيير وميكلون")
	dataSaintPierreAndMiquelon.RegisterCapital(xlanguage.Arabic, "سان بيير")
}
