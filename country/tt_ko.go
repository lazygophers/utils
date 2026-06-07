//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTrinidadAndTobago.RegisterName(xlanguage.Korean, "트리니다드 토바고")
	dataTrinidadAndTobago.RegisterOfficialName(xlanguage.Korean, "트리니다드 토바고 공화국")
	dataTrinidadAndTobago.RegisterCapital(xlanguage.Korean, "포트오브스페인")
}
