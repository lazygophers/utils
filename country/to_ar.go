//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTonga.RegisterName(xlanguage.Arabic, "تونغا")
	dataTonga.RegisterOfficialName(xlanguage.Arabic, "مملكة تونغا")
	dataTonga.RegisterCapital(xlanguage.Arabic, "نوكوألوفا")
}
