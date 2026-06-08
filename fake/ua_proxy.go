package fake

import "fmt"

// Proxy-client version / build pools. Static; bumped manually per release
// cadence documented in .trellis/spec/backend/fake-ua-templates.md.
//
// Spelling traps that MUST be preserved verbatim:
//   - "Quantumult%20X" carries a URL-encoded space (literal `%20`).
//   - "clash.meta" carries a dot (not `clash-meta` / `clashmeta`).
//   - "Surge iOS" / "Surge Mac" carry a space.
//   - "sing-box" carries a hyphen.
//   - "SFA" is sing-box Android; "SFI" is sing-box iOS.
var (
	clashGenericVariants = []string{"clash", "Clash", "Clash/v1.18.0"}
	clashVersions        = []string{"v1.17.0", "v1.18.0", "v1.18.1"}
	clashForWindowsVers  = []string{"0.20.39", "0.20.40"}
	clashMetaVersions    = []string{"v1.18.5", "v1.18.6", "v1.18.7", "alpha"}
	mihomoVersions       = []string{"v1.18.7", "v1.18.8", "v1.18.9", "v1.19.0"}
	singBoxVersions      = []string{"1.8.0", "1.9.0", "1.10.0", "1.11.0"}
	sfaVersions          = []string{"1.8.0", "1.9.0", "1.10.0"}
	sfiVersions          = []string{"1.8.0", "1.9.0", "1.10.0"}
	v2rayNVersions       = []string{"6.31", "6.40", "6.42"}
	v2rayNGVersions      = []string{"1.8.12", "1.8.17", "1.9.0"}
	surgeIOSBuilds       = []int{1419, 2543, 2627}
	surgeMacBuilds       = []int{1234, 1450, 1525}
	quantumultXBuilds    = []int{1051, 1080, 1095}
	loonBuilds           = []int{750, 800, 825}
	stashVersions        = []string{"2.5.0", "2.6.0", "3.0.0"}
	surfboardVersions    = []string{"1.3.0", "1.4.0"}
	shadowrocketBuilds   = []int{2099, 2123, 2155}
	cfNetworkVersions    = []string{"893.14.2", "1408.0.4", "1474.0.0"}
	darwinVersions       = []string{"17.3.0", "22.6.0", "23.4.0", "24.0.0"}
)

func clashGeneric(f *Faker) string { return pick(f.rng, clashGenericVariants) }

func clashVer(f *Faker) string { return fmt.Sprintf("Clash/%s", pick(f.rng, clashVersions)) }

func clashForWin(f *Faker) string {
	return fmt.Sprintf("ClashforWindows/%s", pick(f.rng, clashForWindowsVers))
}

func clashMeta(f *Faker) string { return fmt.Sprintf("clash.meta/%s", pick(f.rng, clashMetaVersions)) }

func mihomoUA(f *Faker) string { return fmt.Sprintf("mihomo/%s", pick(f.rng, mihomoVersions)) }

func singBoxUA(f *Faker) string { return fmt.Sprintf("sing-box/%s", pick(f.rng, singBoxVersions)) }

func sfaUA(f *Faker) string { return fmt.Sprintf("SFA/%s", pick(f.rng, sfaVersions)) }

func sfiUA(f *Faker) string { return fmt.Sprintf("SFI/%s", pick(f.rng, sfiVersions)) }

func v2rayNUA(f *Faker) string { return fmt.Sprintf("v2rayN/%s", pick(f.rng, v2rayNVersions)) }

func v2rayNGUA(f *Faker) string { return fmt.Sprintf("v2rayNG/%s", pick(f.rng, v2rayNGVersions)) }

func surgeIOSUA(f *Faker) string {
	return fmt.Sprintf("Surge iOS/%d", surgeIOSBuilds[f.intN(len(surgeIOSBuilds))])
}

func surgeMacUA(f *Faker) string {
	return fmt.Sprintf("Surge Mac/%d", surgeMacBuilds[f.intN(len(surgeMacBuilds))])
}

func quantumultXUA(f *Faker) string {
	// %% in the format string renders as a literal `%`; result begins with
	// the URL-encoded space `%20` that the official client uses.
	return fmt.Sprintf("Quantumult%%20X/%d", quantumultXBuilds[f.intN(len(quantumultXBuilds))])
}

func loonUA(f *Faker) string {
	return fmt.Sprintf("Loon/%d", loonBuilds[f.intN(len(loonBuilds))])
}

func stashUA(f *Faker) string { return fmt.Sprintf("Stash/%s", pick(f.rng, stashVersions)) }

func surfboardUA(f *Faker) string {
	return fmt.Sprintf("Surfboard/%s", pick(f.rng, surfboardVersions))
}

func shadowrocketUA(f *Faker) string {
	return fmt.Sprintf("Shadowrocket/%d CFNetwork/%s Darwin/%s",
		shadowrocketBuilds[f.intN(len(shadowrocketBuilds))],
		pick(f.rng, cfNetworkVersions),
		pick(f.rng, darwinVersions))
}

func init() {
	proxyTemplates = append(proxyTemplates,
		uaTemplate{kind: UAKindProxy, weight: 20, build: clashGeneric},
		uaTemplate{kind: UAKindProxy, weight: 8, build: clashVer},
		uaTemplate{kind: UAKindProxy, weight: 5, build: clashForWin},
		uaTemplate{kind: UAKindProxy, weight: 18, build: clashMeta},
		uaTemplate{kind: UAKindProxy, weight: 12, build: mihomoUA},
		uaTemplate{kind: UAKindProxy, weight: 10, build: singBoxUA},
		uaTemplate{kind: UAKindProxy, weight: 4, build: sfaUA},
		uaTemplate{kind: UAKindProxy, weight: 4, build: sfiUA},
		uaTemplate{kind: UAKindProxy, weight: 5, build: v2rayNUA},
		uaTemplate{kind: UAKindProxy, weight: 8, build: v2rayNGUA},
		uaTemplate{kind: UAKindProxy, weight: 6, build: surgeIOSUA},
		uaTemplate{kind: UAKindProxy, weight: 4, build: surgeMacUA},
		uaTemplate{kind: UAKindProxy, weight: 5, build: quantumultXUA},
		uaTemplate{kind: UAKindProxy, weight: 4, build: loonUA},
		uaTemplate{kind: UAKindProxy, weight: 3, build: stashUA},
		uaTemplate{kind: UAKindProxy, weight: 2, build: surfboardUA},
		uaTemplate{kind: UAKindProxy, weight: 5, build: shadowrocketUA},
	)
}
