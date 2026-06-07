//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Shp.RegisterName(xlanguage.Korean, "세인트헬레나 파운드")
}
