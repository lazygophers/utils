//go:build lang_ja || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Ang.RegisterName(xlanguage.Japanese, "オランダ領アンティル・ギルダー")
}
