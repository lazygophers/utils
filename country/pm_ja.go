//go:build (lang_ja || lang_all) && (country_all || country_americas || country_northern_america || country_pm)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintPierreAndMiquelon.RegisterName(xlanguage.Japanese, "サンピエール島・ミクロン島")
	dataSaintPierreAndMiquelon.RegisterOfficialName(xlanguage.Japanese, "サンピエール島・ミクロン島")
	dataSaintPierreAndMiquelon.RegisterCapital(xlanguage.Japanese, "サンピエール")
}
