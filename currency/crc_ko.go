//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Crc.RegisterName(xlanguage.Korean, "코스타리카 콜론")
}
