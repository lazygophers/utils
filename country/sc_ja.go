//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSeychelles.RegisterName(xlanguage.Japanese, "セーシェル")
	dataSeychelles.RegisterOfficialName(xlanguage.Japanese, "セーシェル共和国")
	dataSeychelles.RegisterCapital(xlanguage.Japanese, "ヴィクトリア")
}
