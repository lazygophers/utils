//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuernsey.RegisterName(xlanguage.Arabic, "غيرنزي")
	dataGuernsey.RegisterOfficialName(xlanguage.Arabic, "إقطاعية غيرنزي")
	dataGuernsey.RegisterCapital(xlanguage.Arabic, "سانت بيتر بورت")
}
