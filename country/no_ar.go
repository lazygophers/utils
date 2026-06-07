//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorway.RegisterName(xlanguage.Arabic, "النرويج")
	dataNorway.RegisterOfficialName(xlanguage.Arabic, "مملكة النرويج")
	dataNorway.RegisterCapital(xlanguage.Arabic, "أوسلو")
}
