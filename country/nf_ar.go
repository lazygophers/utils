//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorfolkIsland.RegisterName(xlanguage.Arabic, "جزيرة نورفولك")
	dataNorfolkIsland.RegisterOfficialName(xlanguage.Arabic, "إقليم جزيرة نورفولك")
	dataNorfolkIsland.RegisterCapital(xlanguage.Arabic, "كينغستون")
}
