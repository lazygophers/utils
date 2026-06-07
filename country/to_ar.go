//go:build (lang_ar || lang_all) && (country_all || country_oceania || country_polynesia || country_to)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTonga.RegisterName(xlanguage.Arabic, "تونغا")
	dataTonga.RegisterOfficialName(xlanguage.Arabic, "مملكة تونغا")
	dataTonga.RegisterCapital(xlanguage.Arabic, "نوكوألوفا")
}
