//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataThailand.RegisterName(xlanguage.Arabic, "تايلاند")
	dataThailand.RegisterOfficialName(xlanguage.Arabic, "مملكة تايلاند")
	dataThailand.RegisterCapital(xlanguage.Arabic, "بانكوك")
}
