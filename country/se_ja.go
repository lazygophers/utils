//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSweden.RegisterName(xlanguage.Japanese, "スウェーデン")
	dataSweden.RegisterOfficialName(xlanguage.Japanese, "スウェーデン王国")
	dataSweden.RegisterCapital(xlanguage.Japanese, "ストックホルム")
}
