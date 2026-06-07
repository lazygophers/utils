//go:build (lang_ar || lang_all) && (country_africa || country_all || country_eastern_africa || country_yt)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMayotte.RegisterName(xlanguage.Arabic, "مايوت")
	dataMayotte.RegisterOfficialName(xlanguage.Arabic, "مايوت")
	dataMayotte.RegisterCapital(xlanguage.Arabic, "مامودزو")
}
