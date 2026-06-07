//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHeardAndMcDonaldIslands.RegisterName(xlanguage.Arabic, "جزيرة هيرد وجزر ماكدونالد")
	dataHeardAndMcDonaldIslands.RegisterOfficialName(xlanguage.Arabic, "إقليم جزيرة هيرد وجزر ماكدونالد")
	dataHeardAndMcDonaldIslands.RegisterCapital(xlanguage.Arabic, "—")
}
