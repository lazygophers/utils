//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBosniaAndHerzegovina.RegisterName(xlanguage.Arabic, "البوسنة والهرسك")
	dataBosniaAndHerzegovina.RegisterOfficialName(xlanguage.Arabic, "البوسنة والهرسك")
	dataBosniaAndHerzegovina.RegisterCapital(xlanguage.Arabic, "سراييفو")
}
