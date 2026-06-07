//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHongKong.RegisterName(xlanguage.Korean, "홍콩")
	dataHongKong.RegisterOfficialName(xlanguage.Korean, "중화인민공화국 홍콩 특별행정구")
	dataHongKong.RegisterCapital(xlanguage.Korean, "홍콩")
}
