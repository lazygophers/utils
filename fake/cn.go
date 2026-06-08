package fake

import (
	"math/rand/v2"
	"strconv"
	"strings"
	"time"

	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/country"
)

// cnAreaCodes is the pool of GB/T 2260 administrative division codes used as
// the 6-digit prefix of generated Chinese resident identity numbers. The list
// covers province-level capitals plus a handful of municipalities so that
// generated IDs always resolve to a real-world region.
var cnAreaCodes = []string{
	"110000", // Beijing
	"120000", // Tianjin
	"130100", // Shijiazhuang, Hebei
	"140100", // Taiyuan, Shanxi
	"150100", // Hohhot, Inner Mongolia
	"210100", // Shenyang, Liaoning
	"220100", // Changchun, Jilin
	"230100", // Harbin, Heilongjiang
	"310000", // Shanghai
	"320100", // Nanjing, Jiangsu
	"330100", // Hangzhou, Zhejiang
	"340100", // Hefei, Anhui
	"350100", // Fuzhou, Fujian
	"360100", // Nanchang, Jiangxi
	"370100", // Jinan, Shandong
	"410100", // Zhengzhou, Henan
	"420100", // Wuhan, Hubei
	"430100", // Changsha, Hunan
	"440100", // Guangzhou, Guangdong
	"450100", // Nanning, Guangxi
	"460100", // Haikou, Hainan
	"500000", // Chongqing
	"510100", // Chengdu, Sichuan
	"520100", // Guiyang, Guizhou
	"530100", // Kunming, Yunnan
	"540100", // Lhasa, Tibet
	"610100", // Xi'an, Shaanxi
	"620100", // Lanzhou, Gansu
	"630100", // Xining, Qinghai
	"640100", // Yinchuan, Ningxia
	"650100", // Urumqi, Xinjiang
}

// cnIdCardWeights is the ISO 7064:1983 (MOD 11-2) weighting vector applied to
// the first 17 digits of a Chinese resident identity number.
var cnIdCardWeights = [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}

// cnIdCardCheckCodes maps the weighted-sum modulo 11 to the published check
// character table for Chinese resident identity numbers.
var cnIdCardCheckCodes = [11]string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}

// cnMobilePrefixes holds the 3-digit mobile number prefixes assigned by MIIT
// to Chinese carriers (China Mobile / Unicom / Telecom / Broadcast).
var cnMobilePrefixes = []string{
	"130", "131", "132", "133", "134", "135", "136", "137", "138", "139",
	"145", "147", "149",
	"150", "151", "152", "153", "155", "156", "157", "158", "159",
	"166",
	"170", "171", "172", "173", "174", "175", "176", "177", "178", "179",
	"180", "181", "182", "183", "184", "185", "186", "187", "188", "189",
	"191", "192", "195", "196", "197", "198", "199",
}

// cnLandlinePrefixes are the trunk area codes of major Chinese cities.
var cnLandlinePrefixes = []string{
	"010", "020", "021", "022", "023", "024", "025", "027", "028", "029",
}

// localeCN registers the China (CN) locale skeleton. Localised pools such as
// first / last names, cities and streets are filled in by the companion
// locale data files (e.g. cn_zh.go) during their own init phase.
var localeCN = &Locale{
	Country:        country.China,
	OfficialLangs:  []xlanguage.Tag{xlanguage.Chinese},
	PhonePrefixes:  cnMobilePrefixes,
	LandlinePrefix: cnLandlinePrefixes,
	ZipFormat:      "######",
	IdCardGen:      genIdCardCN,
	Streets:        map[xlanguage.Tag][]string{},
	Cities:         map[xlanguage.Tag][]CityEntry{},
	FirstNames:     map[xlanguage.Tag]map[Gender][]string{},
	LastNames:      map[xlanguage.Tag][]string{},
	Domain:         "cn",
}

func init() { register(localeCN) }

// genIdCardCN generates an 18-character Chinese resident identity number
// conforming to GB 11643-1999. The output layout is:
//
//	[6 area code][8 birth YYYYMMDD][3 sequence][1 ISO 7064 check]
//
// The sequence digit's parity encodes gender (odd = male, even = female);
// [GenderRandom] is resolved against rng before encoding. When rng is nil the
// runtime-wide math/rand/v2 source is used.
func genIdCardCN(rng *rand.Rand, gender Gender, birth time.Time) string {
	area := pick(rng, cnAreaCodes)

	var b strings.Builder
	b.Grow(18)
	b.WriteString(area)
	b.WriteString(birth.Format("20060102"))

	resolved := gender.Resolve(rng)

	// Sequence: 3 digits in [0, 999] with parity matching the gender.
	var seq int
	if rng != nil {
		seq = rng.IntN(1000)
	} else {
		seq = rand.IntN(1000)
	}
	if resolved == GenderMale && seq%2 == 0 {
		seq = (seq + 1) % 1000
	}
	if resolved == GenderFemale && seq%2 == 1 {
		seq = (seq + 1) % 1000
	}
	// Avoid leading "000" producing an all-zero sequence which is reserved.
	if seq == 0 {
		if resolved == GenderMale {
			seq = 1
		} else {
			seq = 2
		}
	}
	seqStr := strconv.Itoa(seq)
	for i := len(seqStr); i < 3; i++ {
		b.WriteByte('0')
	}
	b.WriteString(seqStr)

	head := b.String()
	sum := 0
	for i := 0; i < 17; i++ {
		sum += int(head[i]-'0') * cnIdCardWeights[i]
	}
	b.WriteString(cnIdCardCheckCodes[sum%11])
	return b.String()
}
