package fake

import "fmt"

// App-embedded WebView UA template field-sampling pools. Values are kept
// here (not in useragents.go) because they are only consumed by this file's
// builder functions; bumping versions does not touch the shared browser
// pools. See .trellis/spec/backend/fake-ua-templates.md (planned) for the
// bump cadence and source of truth for each value.
var (
	// wechatIOSVersions is the iOS WeChat (MicroMessenger) version pool.
	// iOS WeChat reports a three-segment version (8.0.61), distinct from
	// the four-segment form used on Android.
	wechatIOSVersions = []string{"8.0.42", "8.0.55", "8.0.58", "8.0.61"}

	// wechatAndroidVersions is the Android WeChat version pool. Android
	// WeChat reports a four-segment version (8.0.55.2780) where the last
	// segment is a build number.
	wechatAndroidVersions = []string{"8.0.42.2580", "8.0.55.2780", "8.0.58.2900"}

	// qqVersions is the QQ messenger version pool, used by both iOS and
	// Android QQ WebView UAs.
	qqVersions = []string{"8.9.20", "8.9.50", "8.9.80", "9.0.0"}

	// alipayVersions is the Alipay version pool. The same token feeds both
	// AliApp(AP/...) and AlipayClient/... segments.
	alipayVersions = []string{"10.5.30.5300", "10.6.0.6000", "10.7.0.7000"}

	// douyinVersions is the Douyin (TikTok CN) app version pool.
	douyinVersions = []string{"23.4.0", "24.0.0", "25.5.0", "26.2.0"}

	// weiboVersions is the Sina Weibo iOS app version pool.
	weiboVersions = []string{"7.4.1", "7.5.0", "7.6.2"}

	// netTypes is the NetType token pool used by WeChat / QQ on both
	// platforms. WIFI is reported in capitals; cellular tokens follow.
	netTypes = []string{"WIFI", "4G", "5G"}

	// androidBuildIds is a pool of plausible Android Build/<id> strings.
	// Mix of Pixel (UP1A/AP2A/AD1A) and Samsung (TQ3A/RKQ1) IDs to keep
	// the (model, build-id) pair from looking implausible.
	androidBuildIds = []string{"UP1A.231005.007", "AP2A.240605.024", "AD1A.240411.003.B4", "TQ3A.230901.001", "RKQ1.201022.002"}

	// iphoneModels is the iPhone model-suffix pool used by Weibo's
	// underscore-delimited token (iPhone14Pro / iPhone16ProMax / ...).
	iphoneModels = []string{"14Pro", "14ProMax", "15", "15Pro", "15ProMax", "16", "16Pro"}

	// appIOSVersionsDot is the dot-separated iOS version pool used by the
	// Weibo bracket token (...__iphone__os18.0). Kept distinct from the
	// underscore form ([iosVersionsU]) used elsewhere.
	appIOSVersionsDot = []string{"16.6", "17.0", "17.4", "18.0", "18.7"}

	// xwebBuilds is the XWEB/<build> build-number pool used by Android
	// WeChat's Chromium-fork WebView identifier.
	xwebBuilds = []string{"1230018", "1280018", "1300259"}

	// mmwebsdkDates is the MMWEBSDK/<yyyymmdd> stamp pool used by Android
	// WeChat to report its WebView SDK release date.
	mmwebsdkDates = []string{"20240218", "20240807", "20241103"}
)

// hexCode32 returns a random 32-bit value formatted as eight lowercase
// hexadecimal digits. WeChat appends this token after MicroMessenger/<v>
// in the form `(0x18003d39)`; it acts as an opaque per-install identifier.
func (f *Faker) hexCode32() uint32 {
	return uint32(f.uint64() & 0xffffffff)
}

// digitString returns an n-digit decimal string suitable for numeric
// identifier tokens (MMWEBID, TBS build, screen pixel, status-bar height,
// memory size, ...). n must be positive; the leading digit is forced to
// 1..9 so the string keeps the requested length.
func (f *Faker) digitString(n int) string {
	if n <= 0 {
		return ""
	}
	buf := make([]byte, n)
	buf[0] = byte('1' + f.intN(9))
	for i := 1; i < n; i++ {
		buf[i] = byte('0' + f.intN(10))
	}
	return string(buf)
}

// wechatIOS renders an iPhone WeChat (MicroMessenger) UA. iOS WeChat runs
// inside the system WKWebView so the base is AppleWebKit/605.1.15 plus the
// frozen Mobile/15E148 token, followed by the three-segment MicroMessenger
// version, the 32-bit hex identifier, NetType and zh_CN language.
func wechatIOS(f *Faker) string {
	iosVer := pick(f.rng, iosVersionsU)
	wechatVer := pick(f.rng, wechatIOSVersions)
	net := pick(f.rng, netTypes)
	return fmt.Sprintf(
		"Mozilla/5.0 (iPhone; CPU iPhone OS %s like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/%s(0x%08x) NetType/%s Language/zh_CN",
		iosVer, wechatVer, f.hexCode32(), net,
	)
}

// wechatAndroid renders an Android WeChat UA. Android WeChat ships its own
// Chromium fork (XWEB) inside a WebView so the UA carries the `; wv)`
// marker, Version/4.0 (WebView), the Chromium base version, XWEB build,
// MMWEBSDK date stamp, MMWEBID, four-segment MicroMessenger version, hex
// identifier, NetType and zh_CN language.
func wechatAndroid(f *Faker) string {
	dev := pick(f.rng, androidDevices)
	andVer := pick(f.rng, dev.androids)
	buildId := pick(f.rng, androidBuildIds)
	chromeMajor := pick(f.rng, chromeVersions)
	xweb := pick(f.rng, xwebBuilds)
	sdkDate := pick(f.rng, mmwebsdkDates)
	mmwebid := f.digitString(7)
	wechatVer := pick(f.rng, wechatAndroidVersions)
	net := pick(f.rng, netTypes)
	return fmt.Sprintf(
		"Mozilla/5.0 (Linux; Android %s; %s Build/%s; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/%d.0.0.0 Mobile Safari/537.36 XWEB/%s MMWEBSDK/%s MMWEBID/%s MicroMessenger/%s(0x%08x) WeChat/arm64 Weixin NetType/%s Language/zh_CN ABI/arm64",
		andVer, dev.model, buildId, chromeMajor, xweb, sdkDate, mmwebid, wechatVer, f.hexCode32(), net,
	)
}

// qqIOS renders an iPhone QQ UA. Like WeChat iOS it runs in WKWebView
// (AppleWebKit/605.1.15 + Mobile/15E148) and is suffixed with QQ/<v>.<build>
// plus the device pixel width, NetType and memory tokens.
func qqIOS(f *Faker) string {
	iosVer := pick(f.rng, iosVersionsU)
	qqVer := pick(f.rng, qqVersions)
	build := f.intN(9000) + 1000
	net := pick(f.rng, netTypes)
	mem := f.intN(4000) + 2000
	return fmt.Sprintf(
		"Mozilla/5.0 (iPhone; CPU iPhone OS %s like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 QQ/%s.%d Pixel/1170 NetType/%s Mem/%d",
		iosVer, qqVer, build, net, mem,
	)
}

// qqAndroid renders an Android QQ UA. Android QQ embeds the X5 / TBS
// browser kernel so the UA includes the typical V1_AND_SQ_<v>_<rev>_YYB_D
// token plus QQ version, NetType and screen-related fields.
func qqAndroid(f *Faker) string {
	dev := pick(f.rng, androidDevices)
	andVer := pick(f.rng, dev.androids)
	buildId := pick(f.rng, androidBuildIds)
	chromeMajor := pick(f.rng, chromeVersions)
	qqVer := pick(f.rng, qqVersions)
	tbs := f.intN(900) + 100
	net := pick(f.rng, netTypes)
	pixel := f.intN(2000) + 1080
	statusBar := f.intN(60) + 24
	return fmt.Sprintf(
		"Mozilla/5.0 (Linux; U; Android %s; zh-cn; %s Build/%s) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/%d.0.0.0 Mobile Safari/537.36 V1_AND_SQ_%s_%d_YYB_D QQ/%s NetType/%s WebP/0.3.0 Pixel/%d StatusBarHeight/%d SimpleUISwitch/0",
		andVer, dev.model, buildId, chromeMajor, qqVer, tbs, qqVer, net, pixel, statusBar,
	)
}

// alipayIOS renders an iPhone Alipay UA. Both AliApp(AP/<v>) and
// AlipayClient/<v> carry the same version token; Language/zh-Hans is the
// Alipay convention (hyphen + Hans script tag, distinct from zh_CN).
func alipayIOS(f *Faker) string {
	iosVer := pick(f.rng, iosVersionsU)
	apVer := pick(f.rng, alipayVersions)
	return fmt.Sprintf(
		"Mozilla/5.0 (iPhone; CPU iPhone OS %s like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 AliApp(AP/%s) AlipayClient/%s Language/zh-Hans",
		iosVer, apVer, apVer,
	)
}

// alipayAndroid renders an Android Alipay UA. Historically Alipay's
// in-app browser embedded UCBrowser's UA fragment; the token is preserved
// verbatim. AliApp(AP/<v>) and AlipayClient/<v> share the same version.
func alipayAndroid(f *Faker) string {
	dev := pick(f.rng, androidDevices)
	andVer := pick(f.rng, dev.androids)
	buildId := pick(f.rng, androidBuildIds)
	chromeMajor := pick(f.rng, chromeVersions)
	apVer := pick(f.rng, alipayVersions)
	return fmt.Sprintf(
		"Mozilla/5.0 (Linux; U; Android %s; zh-CN; %s Build/%s) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/%d.0.0.0 UCBrowser/11.6.4.950 Mobile Safari/537.36 AliApp(AP/%s) AlipayClient/%s Language/zh-Hans useStatusBar/true",
		andVer, dev.model, buildId, chromeMajor, apVer, apVer,
	)
}

// douyinAndroid renders an Android Douyin (TikTok CN) WebView UA. The
// BytedanceWebview/d8a21c6 hash is constant across builds; aweme_<build>
// carries the WebView build number; ByteLocale and Region pin the app to
// the CN domestic locale.
func douyinAndroid(f *Faker) string {
	dev := pick(f.rng, androidDevices)
	andVer := pick(f.rng, dev.androids)
	buildId := pick(f.rng, androidBuildIds)
	chromeMajor := pick(f.rng, chromeVersions)
	awemeBuild := f.digitString(6)
	net := pick(f.rng, netTypes)
	appVer := pick(f.rng, douyinVersions)
	return fmt.Sprintf(
		"Mozilla/5.0 (Linux; Android %s; %s Build/%s; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/%d.0.0.0 Mobile Safari/537.36 aweme_%s JsSdk/1.0 NetType/%s AppName/aweme app_version/%s ByteLocale/zh-CN Region/CN AppTheme/light BytedanceWebview/d8a21c6 WebView/1",
		andVer, dev.model, buildId, chromeMajor, awemeBuild, net, appVer,
	)
}

// weiboIOS renders an iPhone Weibo UA. The trailing bracket token uses
// the underscore-delimited form `(iPhone<model>__weibo__<v>__iphone__os<X.Y>)`
// where the iOS version is dot-separated (distinct from the underscore form
// in the leading `iPhone OS <X_Y>` segment).
func weiboIOS(f *Faker) string {
	iosVerU := pick(f.rng, iosVersionsU)
	model := pick(f.rng, iphoneModels)
	appVer := pick(f.rng, weiboVersions)
	iosVerDot := pick(f.rng, appIOSVersionsDot)
	return fmt.Sprintf(
		"Mozilla/5.0 (iPhone; CPU iPhone OS %s like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 Weibo (iPhone%s__weibo__%s__iphone__os%s)",
		iosVerU, model, appVer, iosVerDot,
	)
}

// init registers the app-embedded WebView templates into the shared
// appTemplates pool. Weights reflect a typical China-market mix: WeChat
// dominates, QQ and Alipay form the next tier, Douyin and Weibo round out
// the long tail. The os field is set to the actual platform so locale
// helpers (Accept-Language, mobile-vs-desktop classification) can reason
// about the rendered template even though browser is left at its zero
// value (BrowserChrome) — for UAKindApp the browser axis is meaningless.
func init() {
	appTemplates = append(appTemplates,
		uaTemplate{kind: UAKindApp, os: OSIOSPhone, weight: 20, build: wechatIOS},
		uaTemplate{kind: UAKindApp, os: OSAndroid, weight: 30, build: wechatAndroid},
		uaTemplate{kind: UAKindApp, os: OSIOSPhone, weight: 8, build: qqIOS},
		uaTemplate{kind: UAKindApp, os: OSAndroid, weight: 12, build: qqAndroid},
		uaTemplate{kind: UAKindApp, os: OSIOSPhone, weight: 5, build: alipayIOS},
		uaTemplate{kind: UAKindApp, os: OSAndroid, weight: 8, build: alipayAndroid},
		uaTemplate{kind: UAKindApp, os: OSAndroid, weight: 10, build: douyinAndroid},
		uaTemplate{kind: UAKindApp, os: OSIOSPhone, weight: 5, build: weiboIOS},
	)
}
