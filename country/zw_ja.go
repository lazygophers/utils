//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataZimbabwe.RegisterName(xlanguage.Japanese, "ジンバブエ")
	dataZimbabwe.RegisterOfficialName(xlanguage.Japanese, "ジンバブエ共和国")
	dataZimbabwe.RegisterCapital(xlanguage.Japanese, "ハラレ")
}
