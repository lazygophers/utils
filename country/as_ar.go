//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAmericanSamoa.RegisterName(xlanguage.Arabic, "ساموا الأمريكية")
	dataAmericanSamoa.RegisterOfficialName(xlanguage.Arabic, "إقليم ساموا الأمريكية")
	dataAmericanSamoa.RegisterCapital(xlanguage.Arabic, "باغو باغو")
}
