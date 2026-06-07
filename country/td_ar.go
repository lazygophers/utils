//go:build country_africa || country_all || country_middle_africa || country_td

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChad.RegisterName(xlanguage.Arabic, "تشاد")
	dataChad.RegisterOfficialName(xlanguage.Arabic, "جمهورية تشاد")
	dataChad.RegisterCapital(xlanguage.Arabic, "نجامينا")
}
