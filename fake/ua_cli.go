package fake

import "fmt"

// CLI / SDK HTTP client UA template pools. Template literals are
// source-confirmed against upstream code (Claude Code, Codex) or canonical
// emitter strings (curl, wget, requests, Go net/http, axios, node-fetch,
// okhttp, Postman, Insomnia, git). Version windows are static — see
// .trellis/spec/backend/fake-ua-templates.md for the bump cadence.

// claudeCLIVersions samples the Claude Code CLI release line. Format is
// emitted by the official binary as `claude-cli/<ver> (external, cli)`.
var claudeCLIVersions = []string{"1.0.27", "1.0.56", "2.0.76", "2.1.2"}

// codexVersions samples the codex-rs CLI release line. Format is emitted by
// codex-rs/login/src/auth/default_client.rs::get_codex_user_agent().
var codexVersions = []string{"0.95.0", "0.105.0", "0.115.0", "0.121.0"}

// codexOSInfo binds an OS identifier emitted by codex-rs to the set of OS
// release strings and CPU architectures that legitimately ship together.
type codexOSInfo struct {
	name     string
	versions []string
	arches   []string
}

// codexOSes is the OS axis sampled when rendering a Codex UA. Independently
// sampling name / version / arch would produce impossible combinations
// (e.g. Macos on x86_64 with a Linux kernel version).
var codexOSes = []codexOSInfo{
	{"Macos", []string{"14.5.0", "15.0.0", "15.4.0"}, []string{"arm64", "x86_64"}},
	{"Linux", []string{"5.15.0", "6.5.0", "6.8.0"}, []string{"x86_64", "arm64"}},
	{"Windows", []string{"10.0.19045", "10.0.22631"}, []string{"x86_64"}},
}

// curlVersions covers curl 7.x / 8.x release lines commonly seen on Linux
// distros and macOS Homebrew.
var curlVersions = []string{"7.88.1", "8.4.0", "8.5.0", "8.10.1", "8.11.0"}

// wgetVersions covers the GNU wget 1.21 line shipped on modern distros.
var wgetVersions = []string{"1.21.1", "1.21.3", "1.21.4"}

// pythonRequestsVers covers the requests library release line. The library
// emits `python-requests/<ver>` verbatim with no extra qualifiers.
var pythonRequestsVers = []string{"2.28.2", "2.31.0", "2.32.3"}

// goHttpVariants is the *complete* enumeration of UA strings Go net/http
// emits by default. The number is the HTTP protocol version, not the Go
// release version. Do not parametrise.
var goHttpVariants = []string{"Go-http-client/1.1", "Go-http-client/2.0"}

// axiosVersions covers the axios npm package modern release line.
var axiosVersions = []string{"1.4.0", "1.6.2", "1.7.7"}

// nodeFetchVersions covers the node-fetch v3 line. The library appends a
// fixed `(+https://github.com/bitinn/node-fetch)` suffix to the UA.
var nodeFetchVersions = []string{"3.3.0", "3.3.2"}

// okhttpVersions covers OkHttp 4.x / 5.x. Spelling is lowercase `okhttp/`
// — the canonical token is case-sensitive and uppercased forms are an
// anti-bot fingerprintable mistake.
var okhttpVersions = []string{"4.10.0", "4.11.0", "4.12.0", "5.0.0"}

// postmanVersions covers the PostmanRuntime release line. The token is
// PascalCase `PostmanRuntime/<ver>`.
var postmanVersions = []string{"7.39.0", "7.42.0", "7.43.0"}

// insomniaVersions covers Kong Insomnia release lines (mix of date-stamped
// 2023.x and modern 8.x / 9.x).
var insomniaVersions = []string{"2023.5.8", "8.6.1", "9.3.2"}

// gitVersions covers the modern git client release line. Emitted by
// `git clone` / `git fetch` over smart HTTP.
var gitVersions = []string{"2.39.0", "2.42.0", "2.45.0"}

// claudeCLI renders a Claude Code UA. Source: anthropic-ai/claude-code CLI
// binary emits `claude-cli/<semver> (external, cli)` verbatim.
func claudeCLI(f *Faker) string {
	ver := pick(f.rng, claudeCLIVersions)
	return fmt.Sprintf("claude-cli/%s (external, cli)", ver)
}

// codexCLI renders an OpenAI Codex CLI UA. Source:
// codex-rs/login/src/auth/default_client.rs::get_codex_user_agent().
func codexCLI(f *Faker) string {
	osi := codexOSes[f.intN(len(codexOSes))]
	osVer := pick(f.rng, osi.versions)
	arch := pick(f.rng, osi.arches)
	ver := pick(f.rng, codexVersions)
	return fmt.Sprintf("codex_cli_rs/%s (%s %s; %s) rust", ver, osi.name, osVer, arch)
}

// curlUA renders the canonical curl UA `curl/<ver>`.
func curlUA(f *Faker) string {
	return fmt.Sprintf("curl/%s", pick(f.rng, curlVersions))
}

// wgetUA renders the canonical wget UA `Wget/<ver>`.
func wgetUA(f *Faker) string {
	return fmt.Sprintf("Wget/%s", pick(f.rng, wgetVersions))
}

// pythonRequestsUA renders the canonical requests UA.
func pythonRequestsUA(f *Faker) string {
	return fmt.Sprintf("python-requests/%s", pick(f.rng, pythonRequestsVers))
}

// goHttpUA returns the literal Go net/http UA. There are only two values
// in the wild — the HTTP protocol major (`1.1` or `2.0`), not the Go
// runtime version.
func goHttpUA(f *Faker) string {
	return pick(f.rng, goHttpVariants)
}

// axiosUA renders the canonical axios UA.
func axiosUA(f *Faker) string {
	return fmt.Sprintf("axios/%s", pick(f.rng, axiosVersions))
}

// nodeFetchUA renders the node-fetch UA with the fixed bitinn URL suffix.
func nodeFetchUA(f *Faker) string {
	return fmt.Sprintf("node-fetch/%s (+https://github.com/bitinn/node-fetch)", pick(f.rng, nodeFetchVersions))
}

// okhttpUA renders the canonical OkHttp UA. Token is lowercase by design.
func okhttpUA(f *Faker) string {
	return fmt.Sprintf("okhttp/%s", pick(f.rng, okhttpVersions))
}

// postmanUA renders the canonical Postman runtime UA.
func postmanUA(f *Faker) string {
	return fmt.Sprintf("PostmanRuntime/%s", pick(f.rng, postmanVersions))
}

// insomniaUA renders the canonical Insomnia client UA.
func insomniaUA(f *Faker) string {
	return fmt.Sprintf("insomnia/%s", pick(f.rng, insomniaVersions))
}

// gitUA renders the canonical git smart-HTTP UA.
func gitUA(f *Faker) string {
	return fmt.Sprintf("git/%s", pick(f.rng, gitVersions))
}

func init() {
	cliTemplates = append(cliTemplates,
		uaTemplate{kind: UAKindCLI, weight: 4, build: claudeCLI},
		uaTemplate{kind: UAKindCLI, weight: 3, build: codexCLI},
		uaTemplate{kind: UAKindCLI, weight: 15, build: curlUA},
		uaTemplate{kind: UAKindCLI, weight: 5, build: wgetUA},
		uaTemplate{kind: UAKindCLI, weight: 15, build: pythonRequestsUA},
		uaTemplate{kind: UAKindCLI, weight: 12, build: goHttpUA},
		uaTemplate{kind: UAKindCLI, weight: 10, build: axiosUA},
		uaTemplate{kind: UAKindCLI, weight: 5, build: nodeFetchUA},
		uaTemplate{kind: UAKindCLI, weight: 10, build: okhttpUA},
		uaTemplate{kind: UAKindCLI, weight: 8, build: postmanUA},
		uaTemplate{kind: UAKindCLI, weight: 4, build: insomniaUA},
		uaTemplate{kind: UAKindCLI, weight: 9, build: gitUA},
	)
}
