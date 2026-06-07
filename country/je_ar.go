//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJersey.RegisterName(xlanguage.Arabic, "جيرزي")
	dataJersey.RegisterOfficialName(xlanguage.Arabic, "إقطاعية جيرزي")
	dataJersey.RegisterCapital(xlanguage.Arabic, "سانت هيلير")
}
