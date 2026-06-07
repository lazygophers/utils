//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPuertoRico.RegisterName(xlanguage.Japanese, "プエルトリコ")
	dataPuertoRico.RegisterOfficialName(xlanguage.Japanese, "プエルトリコ自治連邦区")
	dataPuertoRico.RegisterCapital(xlanguage.Japanese, "サンフアン")
}
