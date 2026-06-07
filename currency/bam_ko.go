//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Bam.RegisterName(xlanguage.Korean, "보스니아 헤르체고비나 태환 마르카")
}
