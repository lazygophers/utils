//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurkmenistan.RegisterName(xlanguage.Arabic, "تركمانستان")
	dataTurkmenistan.RegisterOfficialName(xlanguage.Arabic, "تركمانستان")
	dataTurkmenistan.RegisterCapital(xlanguage.Arabic, "عشق آباد")
}
