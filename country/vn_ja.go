//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVietnam.RegisterName(xlanguage.Japanese, "ベトナム")
	dataVietnam.RegisterOfficialName(xlanguage.Japanese, "ベトナム社会主義共和国")
	dataVietnam.RegisterCapital(xlanguage.Japanese, "ハノイ")
}
