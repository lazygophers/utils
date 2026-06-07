//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataYemen.RegisterName(xlanguage.Japanese, "イエメン")
	dataYemen.RegisterOfficialName(xlanguage.Japanese, "イエメン共和国")
	dataYemen.RegisterCapital(xlanguage.Japanese, "サナア")
}
