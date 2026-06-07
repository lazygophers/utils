//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Ang.RegisterName(xlanguage.Korean, "네덜란드령 안틸레스 길더")
}
