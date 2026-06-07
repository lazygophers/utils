//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAfghanistan.RegisterName(xlanguage.Korean, "아프가니스탄")
	dataAfghanistan.RegisterOfficialName(xlanguage.Korean, "아프가니스탄 이슬람 토후국")
	dataAfghanistan.RegisterCapital(xlanguage.Korean, "카불")
}
