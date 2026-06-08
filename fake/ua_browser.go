package fake

import (
	"fmt"
	"sort"
)

// samsungBrowserVersions is the sorted slice of SamsungBrowser major
// versions present in samsungChromiumBase. Sorted at init so that seeded
// Fakers reproduce identical samples; map iteration order is randomised.
var samsungBrowserVersions []int

func init() {
	samsungBrowserVersions = make([]int, 0, len(samsungChromiumBase))
	for k := range samsungChromiumBase {
		samsungBrowserVersions = append(samsungBrowserVersions, k)
	}
	sort.Ints(samsungBrowserVersions)

	browserTemplates = append(browserTemplates,
		// Chrome — six platforms.
		uaTemplate{kind: UAKindBrowser, os: OSWindows, browser: BrowserChrome, weight: browserShare[BrowserChrome] * platformShare[OSWindows] / 100, build: chromeWin},
		uaTemplate{kind: UAKindBrowser, os: OSMacOS, browser: BrowserChrome, weight: browserShare[BrowserChrome] * platformShare[OSMacOS] / 100, build: chromeMac},
		uaTemplate{kind: UAKindBrowser, os: OSLinux, browser: BrowserChrome, weight: browserShare[BrowserChrome] * platformShare[OSLinux] / 100, build: chromeLinux},
		uaTemplate{kind: UAKindBrowser, os: OSAndroid, browser: BrowserChrome, weight: browserShare[BrowserChrome] * platformShare[OSAndroid] / 100, build: chromeAndroid},
		uaTemplate{kind: UAKindBrowser, os: OSIOSPhone, browser: BrowserChrome, weight: browserShare[BrowserChrome] * platformShare[OSIOSPhone] / 100, build: chromeIOSPhone},
		uaTemplate{kind: UAKindBrowser, os: OSIOSPad, browser: BrowserChrome, weight: browserShare[BrowserChrome] * platformShare[OSIOSPad] / 100, build: chromeIOSPad},

		// Edge — five platforms (no standalone iPad template; iPhone covers iOS).
		uaTemplate{kind: UAKindBrowser, os: OSWindows, browser: BrowserEdge, weight: browserShare[BrowserEdge] * platformShare[OSWindows] / 100, build: edgeWin},
		uaTemplate{kind: UAKindBrowser, os: OSMacOS, browser: BrowserEdge, weight: browserShare[BrowserEdge] * platformShare[OSMacOS] / 100, build: edgeMac},
		uaTemplate{kind: UAKindBrowser, os: OSLinux, browser: BrowserEdge, weight: browserShare[BrowserEdge] * platformShare[OSLinux] / 100, build: edgeLinux},
		uaTemplate{kind: UAKindBrowser, os: OSAndroid, browser: BrowserEdge, weight: browserShare[BrowserEdge] * platformShare[OSAndroid] / 100, build: edgeAndroid},
		uaTemplate{kind: UAKindBrowser, os: OSIOSPhone, browser: BrowserEdge, weight: browserShare[BrowserEdge] * platformShare[OSIOSPhone] / 100, build: edgeIOSPhone},

		// Firefox — six platforms.
		uaTemplate{kind: UAKindBrowser, os: OSWindows, browser: BrowserFirefox, weight: browserShare[BrowserFirefox] * platformShare[OSWindows] / 100, build: firefoxWin},
		uaTemplate{kind: UAKindBrowser, os: OSMacOS, browser: BrowserFirefox, weight: browserShare[BrowserFirefox] * platformShare[OSMacOS] / 100, build: firefoxMac},
		uaTemplate{kind: UAKindBrowser, os: OSLinux, browser: BrowserFirefox, weight: browserShare[BrowserFirefox] * platformShare[OSLinux] / 100, build: firefoxLinux},
		uaTemplate{kind: UAKindBrowser, os: OSAndroid, browser: BrowserFirefox, weight: browserShare[BrowserFirefox] * platformShare[OSAndroid] / 100, build: firefoxAndroid},
		uaTemplate{kind: UAKindBrowser, os: OSIOSPhone, browser: BrowserFirefox, weight: browserShare[BrowserFirefox] * platformShare[OSIOSPhone] / 100, build: firefoxIOSPhone},
		uaTemplate{kind: UAKindBrowser, os: OSIOSPad, browser: BrowserFirefox, weight: browserShare[BrowserFirefox] * platformShare[OSIOSPad] / 100, build: firefoxIOSPad},

		// Safari — macOS / iPhone / iPad only.
		uaTemplate{kind: UAKindBrowser, os: OSMacOS, browser: BrowserSafari, weight: browserShare[BrowserSafari] * platformShare[OSMacOS] / 100, build: safariMac},
		uaTemplate{kind: UAKindBrowser, os: OSIOSPhone, browser: BrowserSafari, weight: browserShare[BrowserSafari] * platformShare[OSIOSPhone] / 100, build: safariIOSPhone},
		uaTemplate{kind: UAKindBrowser, os: OSIOSPad, browser: BrowserSafari, weight: browserShare[BrowserSafari] * platformShare[OSIOSPad] / 100, build: safariIOSPad},

		// Opera — five platforms.
		uaTemplate{kind: UAKindBrowser, os: OSWindows, browser: BrowserOpera, weight: browserShare[BrowserOpera] * platformShare[OSWindows] / 100, build: operaWin},
		uaTemplate{kind: UAKindBrowser, os: OSMacOS, browser: BrowserOpera, weight: browserShare[BrowserOpera] * platformShare[OSMacOS] / 100, build: operaMac},
		uaTemplate{kind: UAKindBrowser, os: OSLinux, browser: BrowserOpera, weight: browserShare[BrowserOpera] * platformShare[OSLinux] / 100, build: operaLinux},
		uaTemplate{kind: UAKindBrowser, os: OSAndroid, browser: BrowserOpera, weight: browserShare[BrowserOpera] * platformShare[OSAndroid] / 100, build: operaAndroid},
		uaTemplate{kind: UAKindBrowser, os: OSIOSPhone, browser: BrowserOpera, weight: browserShare[BrowserOpera] * platformShare[OSIOSPhone] / 100, build: operaIOSPhone},

		// Samsung Internet — Android only.
		uaTemplate{kind: UAKindBrowser, os: OSAndroid, browser: BrowserSamsung, weight: browserShare[BrowserSamsung] * platformShare[OSAndroid] / 100, build: samsungAndroid},
	)

	// Floor every weight at 1 so that the rare combos (Firefox Linux,
	// Opera macOS, ...) remain reachable.
	for i := range browserTemplates {
		if browserTemplates[i].weight <= 0 {
			browserTemplates[i].weight = 1
		}
	}
}

// pickAndroidDevice picks a (model, androidVersion) pair from
// androidDevices so that the Android version is one the model actually
// shipped with.
func pickAndroidDevice(f *Faker) (string, string) {
	d := androidDevices[f.intN(len(androidDevices))]
	av := d.androids[f.intN(len(d.androids))]
	return d.model, av
}

// pickIntFrom returns a random element from a non-empty int slice.
func pickIntFrom(f *Faker, s []int) int {
	return s[f.intN(len(s))]
}

// pickStrFrom returns a random element from a non-empty string slice.
func pickStrFrom(f *Faker, s []string) string {
	return s[f.intN(len(s))]
}

// samsungVersionAndChromium draws a SamsungBrowser major version and
// returns the Chromium major bound to it by samsungChromiumBase.
func samsungVersionAndChromium(f *Faker) (int, int) {
	sb := samsungBrowserVersions[f.intN(len(samsungBrowserVersions))]
	return sb, samsungChromiumBase[sb]
}

// --- Chrome ---------------------------------------------------------------

func chromeWin(f *Faker) string {
	c := pickIntFrom(f, chromeVersions)
	return fmt.Sprintf("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%d.0.0.0 Safari/537.36", c)
}

func chromeMac(f *Faker) string {
	c := pickIntFrom(f, chromeVersions)
	return fmt.Sprintf("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%d.0.0.0 Safari/537.36", c)
}

func chromeLinux(f *Faker) string {
	c := pickIntFrom(f, chromeVersions)
	return fmt.Sprintf("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%d.0.0.0 Safari/537.36", c)
}

func chromeAndroid(f *Faker) string {
	model, av := pickAndroidDevice(f)
	c := pickIntFrom(f, chromeVersions)
	return fmt.Sprintf("Mozilla/5.0 (Linux; Android %s; %s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%d.0.0.0 Mobile Safari/537.36", av, model, c)
}

func chromeIOSPhone(f *Faker) string {
	ios := pickStrFrom(f, iosVersionsU)
	c := pickIntFrom(f, chromeVersions)
	return fmt.Sprintf("Mozilla/5.0 (iPhone; CPU iPhone OS %s like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/%d.0.0.0 Mobile/15E148 Safari/604.1", ios, c)
}

func chromeIOSPad(f *Faker) string {
	ios := pickStrFrom(f, iosVersionsU)
	c := pickIntFrom(f, chromeVersions)
	return fmt.Sprintf("Mozilla/5.0 (iPad; CPU OS %s like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/%d.0.0.0 Mobile/15E148 Safari/604.1", ios, c)
}

// --- Edge -----------------------------------------------------------------

func edgeWin(f *Faker) string {
	e := pickIntFrom(f, edgeVersions)
	return fmt.Sprintf("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%d.0.0.0 Safari/537.36 Edg/%d.0.0.0", e, e)
}

func edgeMac(f *Faker) string {
	e := pickIntFrom(f, edgeVersions)
	return fmt.Sprintf("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%d.0.0.0 Safari/537.36 Edg/%d.0.0.0", e, e)
}

func edgeLinux(f *Faker) string {
	e := pickIntFrom(f, edgeVersions)
	return fmt.Sprintf("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%d.0.0.0 Safari/537.36 Edg/%d.0.0.0", e, e)
}

func edgeAndroid(f *Faker) string {
	model, av := pickAndroidDevice(f)
	e := pickIntFrom(f, edgeVersions)
	return fmt.Sprintf("Mozilla/5.0 (Linux; Android %s; %s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%d.0.0.0 Mobile Safari/537.36 EdgA/%d.0.0.0", av, model, e, e)
}

func edgeIOSPhone(f *Faker) string {
	ios := pickStrFrom(f, iosVersionsU)
	e := pickIntFrom(f, edgeVersions)
	return fmt.Sprintf("Mozilla/5.0 (iPhone; CPU iPhone OS %s like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 EdgiOS/%d.0.0.0 Mobile/15E148 Safari/604.1", ios, e)
}

// --- Firefox --------------------------------------------------------------

func firefoxWin(f *Faker) string {
	v := pickIntFrom(f, firefoxVersions)
	return fmt.Sprintf("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:%d.0) Gecko/20100101 Firefox/%d.0", v, v)
}

func firefoxMac(f *Faker) string {
	v := pickIntFrom(f, firefoxVersions)
	// macOS Firefox token uses dotted "10.15" — not the underscore form
	// Safari emits. Anti-bot fingerprinters key on this delta.
	return fmt.Sprintf("Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:%d.0) Gecko/20100101 Firefox/%d.0", v, v)
}

func firefoxLinux(f *Faker) string {
	v := pickIntFrom(f, firefoxVersions)
	return fmt.Sprintf("Mozilla/5.0 (X11; Linux x86_64; rv:%d.0) Gecko/20100101 Firefox/%d.0", v, v)
}

func firefoxAndroid(f *Faker) string {
	av := pickStrFrom(f, androidVersions)
	v := pickIntFrom(f, firefoxVersions)
	// Android Gecko trail equals the Firefox major; desktop uses 20100101.
	return fmt.Sprintf("Mozilla/5.0 (Android %s; Mobile; rv:%d.0) Gecko/%d.0 Firefox/%d.0", av, v, v, v)
}

func firefoxIOSPhone(f *Faker) string {
	ios := pickStrFrom(f, iosVersionsU)
	v := pickIntFrom(f, firefoxVersions)
	return fmt.Sprintf("Mozilla/5.0 (iPhone; CPU iPhone OS %s like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/%d.0 Mobile/15E148 Safari/605.1.15", ios, v)
}

func firefoxIOSPad(f *Faker) string {
	ios := pickStrFrom(f, iosVersionsU)
	v := pickIntFrom(f, firefoxVersions)
	return fmt.Sprintf("Mozilla/5.0 (iPad; CPU OS %s like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/%d.0 Mobile/15E148 Safari/605.1.15", ios, v)
}

// --- Safari ---------------------------------------------------------------

func safariMac(f *Faker) string {
	_ = f
	return "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 Safari/605.1.15"
}

func safariIOSPhone(f *Faker) string {
	ios := pickStrFrom(f, iosVersionsU)
	return fmt.Sprintf("Mozilla/5.0 (iPhone; CPU iPhone OS %s like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 Mobile/15E148 Safari/604.1", ios)
}

func safariIOSPad(f *Faker) string {
	ios := pickStrFrom(f, iosVersionsU)
	return fmt.Sprintf("Mozilla/5.0 (iPad; CPU OS %s like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 Mobile/15E148 Safari/604.1", ios)
}

// --- Opera ----------------------------------------------------------------

func operaWin(f *Faker) string {
	c := pickIntFrom(f, chromeVersions)
	o := pickIntFrom(f, operaVersions)
	return fmt.Sprintf("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%d.0.0.0 Safari/537.36 OPR/%d.0.0.0", c, o)
}

func operaMac(f *Faker) string {
	c := pickIntFrom(f, chromeVersions)
	o := pickIntFrom(f, operaVersions)
	return fmt.Sprintf("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%d.0.0.0 Safari/537.36 OPR/%d.0.0.0", c, o)
}

func operaLinux(f *Faker) string {
	c := pickIntFrom(f, chromeVersions)
	o := pickIntFrom(f, operaVersions)
	return fmt.Sprintf("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%d.0.0.0 Safari/537.36 OPR/%d.0.0.0", c, o)
}

func operaAndroid(f *Faker) string {
	model, av := pickAndroidDevice(f)
	c := pickIntFrom(f, chromeVersions)
	o := pickIntFrom(f, operaVersions)
	return fmt.Sprintf("Mozilla/5.0 (Linux; Android %s; %s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%d.0.0.0 Mobile Safari/537.36 OPR/%d.0.0.0", av, model, c, o)
}

func operaIOSPhone(f *Faker) string {
	ios := pickStrFrom(f, iosVersionsU)
	// OPT iOS uses a three-segment marketing version such as 5.0.4.
	major := 4 + f.intN(3)   // 4..6
	minor := f.intN(5)       // 0..4
	patch := f.intN(10)      // 0..9
	return fmt.Sprintf("Mozilla/5.0 (iPhone; CPU iPhone OS %s like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 OPT/%d.%d.%d Mobile/15E148 Safari/604.1", ios, major, minor, patch)
}

// --- Samsung Internet -----------------------------------------------------

func samsungAndroid(f *Faker) string {
	model, av := pickAndroidDevice(f)
	sb, c := samsungVersionAndChromium(f)
	return fmt.Sprintf("Mozilla/5.0 (Linux; Android %s; %s) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/%d.0 Chrome/%d.0.0.0 Mobile Safari/537.36", av, model, sb, c)
}
