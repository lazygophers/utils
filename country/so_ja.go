//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSomalia.RegisterName(xlanguage.Japanese, "ソマリア")
	dataSomalia.RegisterOfficialName(xlanguage.Japanese, "ソマリア連邦共和国")
	dataSomalia.RegisterCapital(xlanguage.Japanese, "モガディシュ")
}
