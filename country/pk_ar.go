//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPakistan.RegisterName(xlanguage.Arabic, "باكستان")
	dataPakistan.RegisterOfficialName(xlanguage.Arabic, "جمهورية باكستان الإسلامية")
	dataPakistan.RegisterCapital(xlanguage.Arabic, "إسلام أباد")
}
