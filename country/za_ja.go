//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthAfrica.RegisterName(xlanguage.Japanese, "南アフリカ")
	dataSouthAfrica.RegisterOfficialName(xlanguage.Japanese, "南アフリカ共和国")
	dataSouthAfrica.RegisterCapital(xlanguage.Japanese, "プレトリア")
}
