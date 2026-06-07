//go:build (lang_ko || lang_all) && (country_africa || country_all || country_sh || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintHelena.RegisterName(xlanguage.Korean, "세인트헬레나")
	dataSaintHelena.RegisterOfficialName(xlanguage.Korean, "세인트헬레나, 어센션, 트리스탄다쿠냐")
	dataSaintHelena.RegisterCapital(xlanguage.Korean, "제임스타운")
}
