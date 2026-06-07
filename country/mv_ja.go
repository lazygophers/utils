//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMaldives.RegisterName(xlanguage.Japanese, "モルディブ")
	dataMaldives.RegisterOfficialName(xlanguage.Japanese, "モルディブ共和国")
	dataMaldives.RegisterCapital(xlanguage.Japanese, "マレ")
}
