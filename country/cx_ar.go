//go:build (lang_ar || lang_all) && (country_all || country_australia_and_new_zealand || country_cx || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChristmasIsland.RegisterName(xlanguage.Arabic, "جزيرة عيد الميلاد")
	dataChristmasIsland.RegisterOfficialName(xlanguage.Arabic, "إقليم جزيرة عيد الميلاد")
	dataChristmasIsland.RegisterCapital(xlanguage.Arabic, "فلاينغ فيش كوف")
}
