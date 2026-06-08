package fake

import "math/rand/v2"

// UAKind enumerates the high-level category of a User-Agent template.
// The four categories drive the top-level mix when [Faker.UserAgent] is
// invoked and gate which template pool is sampled.
type UAKind uint8

// User-Agent kind constants. See [UAKind] for the role of each bucket.
const (
	// UAKindBrowser covers standalone browser engines on desktop and mobile.
	UAKindBrowser UAKind = iota
	// UAKindApp covers app-embedded web views (WeChat, QQ, Alipay, ...).
	UAKindApp
	// UAKindCLI covers HTTP clients shipped as CLIs or SDKs (curl, requests).
	UAKindCLI
	// UAKindProxy covers proxy clients (Clash, sing-box, Surge, ...).
	UAKindProxy
)

// String returns the lowercase kebab-case identifier of the UA kind.
func (k UAKind) String() string {
	switch k {
	case UAKindBrowser:
		return "browser"
	case UAKindApp:
		return "app"
	case UAKindCLI:
		return "cli"
	case UAKindProxy:
		return "proxy"
	}
	return "unknown"
}

// OS enumerates the operating-system axis used by browser UA templates.
type OS uint8

// Operating system constants. See [OS].
const (
	// OSWindows is desktop Windows (NT 10.0 / Win64; x64 token).
	OSWindows OS = iota
	// OSMacOS is desktop Apple macOS (Intel Mac OS X 10_15_7 token).
	OSMacOS
	// OSLinux is desktop Linux (X11; Linux x86_64 token).
	OSLinux
	// OSAndroid is mobile Android (Linux; Android <ver>; <model> token).
	OSAndroid
	// OSIOSPhone is iPhone-flavoured iOS (iPhone; CPU iPhone OS token).
	OSIOSPhone
	// OSIOSPad is iPad-flavoured iPadOS (iPad; CPU OS token).
	OSIOSPad
)

// String returns the lowercase kebab-case identifier of the OS.
func (o OS) String() string {
	switch o {
	case OSWindows:
		return "windows"
	case OSMacOS:
		return "macos"
	case OSLinux:
		return "linux"
	case OSAndroid:
		return "android"
	case OSIOSPhone:
		return "ios-phone"
	case OSIOSPad:
		return "ios-pad"
	}
	return "unknown"
}

// Browser enumerates the major browser engines covered by templates.
type Browser uint8

// Browser engine constants. See [Browser].
const (
	// BrowserChrome is Google Chrome on Chromium.
	BrowserChrome Browser = iota
	// BrowserEdge is Microsoft Edge on Chromium (Edg/EdgA/EdgiOS token).
	BrowserEdge
	// BrowserFirefox is Mozilla Firefox / Gecko (or FxiOS on iOS).
	BrowserFirefox
	// BrowserSafari is Apple Safari on WebKit (macOS / iOS only).
	BrowserSafari
	// BrowserOpera is Opera on Chromium (OPR / OPT on iOS).
	BrowserOpera
	// BrowserSamsung is Samsung Internet on Chromium (Android only).
	BrowserSamsung
)

// String returns the lowercase identifier of the browser.
func (b Browser) String() string {
	switch b {
	case BrowserChrome:
		return "chrome"
	case BrowserEdge:
		return "edge"
	case BrowserFirefox:
		return "firefox"
	case BrowserSafari:
		return "safari"
	case BrowserOpera:
		return "opera"
	case BrowserSamsung:
		return "samsung"
	}
	return "unknown"
}

// uaTemplate is one entry in a UA template pool. kind selects the bucket;
// os and browser are only meaningful when kind == [UAKindBrowser] (browser
// templates) or carry app-platform information for [UAKindApp]. weight is
// the base weight used when no locale-level bias overrides it; weight == 0
// means the template is excluded from weighted sampling and only reachable
// through targeted lookups (e.g. [Faker.BrowserUAOf]). build receives the
// owning Faker and returns the fully rendered UA string; the function is
// responsible for sampling version / device / language fields.
type uaTemplate struct {
	kind    UAKind
	os      OS
	browser Browser
	weight  int
	build   func(*Faker) string
}

// androidDevice describes one Android device model together with the set
// of Android OS releases that legitimately shipped on it. Used by Android
// browser and WeChat templates to avoid impossible model / OS combinations.
type androidDevice struct {
	model    string
	androids []string
}

// Top-level UA kind mix weights used by [Faker.UserAgent].
const (
	// weightBrowser is the default weight given to standalone browsers.
	weightBrowser = 80
	// weightApp is the default weight given to app web views.
	weightApp = 15
	// weightCLI is the default weight given to CLI / SDK clients.
	weightCLI = 3
	// weightProxy is the default weight given to proxy clients.
	weightProxy = 2
)

// fallbackUA is the hard-coded UA returned when every pool is empty. It
// matches a modern desktop Chrome on Windows so that downstream consumers
// never observe an empty User-Agent header.
const fallbackUA = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36"

// Browser version sampling windows. Inclusive ranges; bumped manually on
// release. See .trellis/spec/backend/fake-ua-templates.md for cadence.
var (
	chromeVersions  = makeRange(120, 149)
	firefoxVersions = makeRange(120, 151)
	edgeVersions    = makeRange(120, 145)
	operaVersions   = makeRange(106, 126)
)

// iOS version pools. Underscore form matches the `iPhone OS X_Y` token
// in Apple's WebKit UA grammar; dot form matches `Version/X.Y` style
// segments used by some Chromium-on-iOS templates.
var (
	iosVersionsU = []string{"16_6", "16_7", "17_0", "17_1", "17_4", "17_5", "18_0", "18_4", "18_7"}
)

// androidVersions is the pool of Android OS releases sampled for generic
// templates that do not pin to a specific device.
var androidVersions = []string{"12", "13", "14", "15"}

// androidDevices is the pool of (model, supported Android versions) pairs
// used to keep Android UA strings internally consistent.
var androidDevices = []androidDevice{
	{"Pixel 9 Pro", []string{"14", "15"}},
	{"Pixel 9", []string{"14", "15"}},
	{"Pixel 8 Pro", []string{"14", "15"}},
	{"Pixel 8", []string{"14", "15"}},
	{"Pixel 7", []string{"13", "14"}},
	{"Pixel 6", []string{"12", "13", "14"}},
	{"SM-S928B", []string{"14", "15"}},
	{"SM-S921B", []string{"14", "15"}},
	{"SM-S918B", []string{"13", "14"}},
	{"SM-S911B", []string{"13", "14"}},
	{"SM-A546B", []string{"13", "14"}},
	{"SM-G991B", []string{"12", "13", "14"}},
	{"SM-F946B", []string{"13", "14"}},
}

// samsungChromiumBase maps a SamsungBrowser major version to the Chromium
// major version it shipped on. Independently sampling the two majors would
// produce impossible combinations spotted by anti-bot fingerprinters.
var samsungChromiumBase = map[int]int{
	24: 117,
	25: 121,
	26: 125,
	27: 128,
	28: 130,
	29: 135,
	30: 140,
}

// browserShare is the 2026 global desktop+mobile market share used for
// weighted browser sampling when no locale bias is supplied. Sums to 100.
var browserShare = map[Browser]int{
	BrowserChrome:  65,
	BrowserSamsung: 7,
	BrowserEdge:    5,
	BrowserSafari:  18,
	BrowserFirefox: 3,
	BrowserOpera:   2,
}

// platformShare is the 2026 global platform share used for weighted OS
// sampling inside [Faker.BrowserUA]. Sums to 100.
var platformShare = map[OS]int{
	OSWindows:  30,
	OSMacOS:    12,
	OSLinux:    3,
	OSAndroid:  35,
	OSIOSPhone: 18,
	OSIOSPad:   2,
}

// Template pools. Populated by ua_browser.go / ua_app.go / ua_cli.go /
// ua_proxy.go init blocks; declared here so the entry-point methods can
// reference them without an import-order dance.
var (
	browserTemplates []uaTemplate
	appTemplates     []uaTemplate
	cliTemplates     []uaTemplate
	proxyTemplates   []uaTemplate
)

// browserLegalOS lists, per browser engine, the operating systems on which
// that engine legitimately ships. Used to reject impossible combinations
// (Safari on Windows, Samsung Internet on iOS, ...) at sample time.
var browserLegalOS = map[Browser]map[OS]bool{
	BrowserChrome: {
		OSWindows: true, OSMacOS: true, OSLinux: true,
		OSAndroid: true, OSIOSPhone: true, OSIOSPad: true,
	},
	BrowserEdge: {
		OSWindows: true, OSMacOS: true, OSLinux: true,
		OSAndroid: true, OSIOSPhone: true, OSIOSPad: true,
	},
	BrowserFirefox: {
		OSWindows: true, OSMacOS: true, OSLinux: true,
		OSAndroid: true, OSIOSPhone: true, OSIOSPad: true,
	},
	BrowserSafari: {
		OSMacOS: true, OSIOSPhone: true, OSIOSPad: true,
	},
	BrowserOpera: {
		OSWindows: true, OSMacOS: true, OSLinux: true,
		OSAndroid: true, OSIOSPhone: true, OSIOSPad: true,
	},
	BrowserSamsung: {
		OSAndroid: true,
	},
}

// desktopOS is the set of OSes considered "desktop" by [Faker.DesktopUA].
var desktopOS = map[OS]bool{OSWindows: true, OSMacOS: true, OSLinux: true}

// mobileOS is the set of OSes considered "mobile" by [Faker.MobileUA].
var mobileOS = map[OS]bool{OSAndroid: true, OSIOSPhone: true, OSIOSPad: true}

// safariFallbackOS is the OS used when callers request Safari on an OS
// it never shipped on. macOS Safari is the most common form of the engine.
const safariFallbackOS = OSMacOS

// UserAgent returns a User-Agent string sampled across all four buckets
// (browser / app / cli / proxy) using the top-level kind mix. Locale-level
// bias is honoured when present; otherwise the package-global weights
// apply. Always returns a non-empty value: when every pool is empty the
// [fallbackUA] desktop-Chrome string is returned.
func (f *Faker) UserAgent() string {
	kindWeights := map[UAKind]int{
		UAKindBrowser: weightBrowser,
		UAKindApp:     weightApp,
		UAKindCLI:     weightCLI,
		UAKindProxy:   weightProxy,
	}
	kind := f.pickKind(kindWeights)
	switch kind {
	case UAKindBrowser:
		return f.BrowserUA()
	case UAKindApp:
		return f.AppUA()
	case UAKindCLI:
		return f.CLIUA()
	case UAKindProxy:
		return f.ProxyUA()
	}
	return fallbackUA
}

// BrowserUA returns a browser UA sampled by global market share. Illegal
// (browser, OS) combinations are resampled until a legal pair is drawn.
// When the browser template pool is empty the [fallbackUA] is returned.
func (f *Faker) BrowserUA() string {
	if len(browserTemplates) == 0 {
		return fallbackUA
	}
	for attempts := 0; attempts < 16; attempts++ {
		br := f.pickBrowser(browserShare)
		osx := f.pickOS(platformShare)
		legal, ok := browserLegalOS[br]
		if !ok || !legal[osx] {
			continue
		}
		ua := f.findBrowserTemplate(osx, br)
		if ua != "" {
			return ua
		}
	}
	return f.pickTemplate(browserTemplates)
}

// DesktopUA restricts [Faker.BrowserUA] to desktop OSes (Windows, macOS,
// Linux). Falls back to [fallbackUA] when the pool is empty.
func (f *Faker) DesktopUA() string {
	return f.browserUAFiltered(desktopOS)
}

// MobileUA restricts [Faker.BrowserUA] to mobile OSes (Android, iOS).
// Falls back to [fallbackUA] when the pool is empty.
func (f *Faker) MobileUA() string {
	return f.browserUAFiltered(mobileOS)
}

// BrowserUAOf returns a UA for the explicit (os, br) pair. When the pair
// is illegal (e.g. Safari on Windows), the browser's most common legal OS
// is substituted; for Safari that is macOS, for all others Chrome falls
// back to the first available template of the browser. Returns
// [fallbackUA] when no template matches.
func (f *Faker) BrowserUAOf(osx OS, br Browser) string {
	legal, ok := browserLegalOS[br]
	if !ok || !legal[osx] {
		if br == BrowserSafari {
			osx = safariFallbackOS
		} else if br == BrowserSamsung {
			osx = OSAndroid
		} else {
			osx = OSWindows
		}
	}
	ua := f.findBrowserTemplate(osx, br)
	if ua != "" {
		return ua
	}
	for _, t := range browserTemplates {
		if t.browser == br {
			return t.build(f)
		}
	}
	return fallbackUA
}

// AppUA returns a UA sampled from the app-embedded WebView pool. When the
// app pool is empty the call degrades to [Faker.BrowserUA].
func (f *Faker) AppUA() string {
	if len(appTemplates) == 0 {
		return f.BrowserUA()
	}
	return f.pickTemplate(appTemplates)
}

// CLIUA returns a UA sampled from the CLI / SDK client pool. When the
// pool is empty the call degrades to [Faker.BrowserUA].
func (f *Faker) CLIUA() string {
	if len(cliTemplates) == 0 {
		return f.BrowserUA()
	}
	return f.pickTemplate(cliTemplates)
}

// ProxyUA returns a UA sampled from the proxy-client pool. When the pool
// is empty the call degrades to [Faker.BrowserUA].
func (f *Faker) ProxyUA() string {
	if len(proxyTemplates) == 0 {
		return f.BrowserUA()
	}
	return f.pickTemplate(proxyTemplates)
}

// browserUAFiltered samples a browser UA constrained to the OSes in
// allowedOS. Used by [Faker.DesktopUA] / [Faker.MobileUA].
func (f *Faker) browserUAFiltered(allowedOS map[OS]bool) string {
	if len(browserTemplates) == 0 {
		return fallbackUA
	}
	candidates := make([]uaTemplate, 0, len(browserTemplates))
	for _, t := range browserTemplates {
		if allowedOS[t.os] {
			candidates = append(candidates, t)
		}
	}
	if len(candidates) == 0 {
		return f.pickTemplate(browserTemplates)
	}
	return f.pickTemplate(candidates)
}

// findBrowserTemplate looks up an exact (os, br) template in the global
// pool and runs its builder. Returns "" when no matching template exists.
func (f *Faker) findBrowserTemplate(osx OS, br Browser) string {
	matches := make([]uaTemplate, 0, 2)
	for _, t := range browserTemplates {
		if t.os == osx && t.browser == br {
			matches = append(matches, t)
		}
	}
	if len(matches) == 0 {
		return ""
	}
	if len(matches) == 1 {
		return matches[0].build(f)
	}
	return matches[f.intN(len(matches))].build(f)
}

// pickTemplate weight-samples one template from pool. Templates with
// weight == 0 are treated as having weight 1 so that explicit pools (e.g.
// CLI / proxy) always remain reachable.
func (f *Faker) pickTemplate(pool []uaTemplate) string {
	if len(pool) == 0 {
		return fallbackUA
	}
	total := 0
	for _, t := range pool {
		w := t.weight
		if w <= 0 {
			w = 1
		}
		total += w
	}
	if total <= 0 {
		return pool[f.intN(len(pool))].build(f)
	}
	r := f.intN(total)
	for _, t := range pool {
		w := t.weight
		if w <= 0 {
			w = 1
		}
		if r < w {
			return t.build(f)
		}
		r -= w
	}
	return pool[len(pool)-1].build(f)
}

// pickKind selects a UAKind from weights using the Faker's RNG.
func (f *Faker) pickKind(weights map[UAKind]int) UAKind {
	return pickWeighted(f, weights, UAKindBrowser)
}

// pickBrowser selects a Browser from weights using the Faker's RNG.
func (f *Faker) pickBrowser(weights map[Browser]int) Browser {
	return pickWeighted(f, weights, BrowserChrome)
}

// pickOS selects an OS from weights using the Faker's RNG.
func (f *Faker) pickOS(weights map[OS]int) OS {
	return pickWeighted(f, weights, OSWindows)
}

// pickWeighted draws one key from weights proportionally to the value.
// The fallback is returned when the map is empty or every weight is
// non-positive. Iteration order over a map is randomised by the runtime;
// callers that need stable ordering should sort externally before
// invoking this helper.
func pickWeighted[K comparable](f *Faker, weights map[K]int, fallback K) K {
	total := 0
	for _, w := range weights {
		if w > 0 {
			total += w
		}
	}
	if total <= 0 {
		return fallback
	}
	r := f.intN(total)
	for k, w := range weights {
		if w <= 0 {
			continue
		}
		if r < w {
			return k
		}
		r -= w
	}
	return fallback
}

// makeRange returns the closed integer interval [lo, hi] as a slice.
// Callers use the result with [pick] to sample a random integer in range.
func makeRange(lo, hi int) []int {
	if hi < lo {
		return nil
	}
	out := make([]int, 0, hi-lo+1)
	for v := lo; v <= hi; v++ {
		out = append(out, v)
	}
	return out
}

// _ keeps math/rand/v2 referenced; template build funcs in ua_*.go consume
// the rand package via pick / shuffle helpers.
var _ = rand.IntN
