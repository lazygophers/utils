package network

import (
	"net/netip"
	"testing"
)

type cidrMergeCase struct {
	name string
	in   []string
	want []string
}

func TestMergeCIDRs(t *testing.T) {
	cases := []cidrMergeCase{
		{
			name: "empty",
			in:   nil,
			want: nil,
		},
		{
			name: "single",
			in:   []string{"10.0.0.0/24"},
			want: []string{"10.0.0.0/24"},
		},
		{
			name: "siblings merge to parent",
			in:   []string{"10.0.0.0/25", "10.0.0.128/25"},
			want: []string{"10.0.0.0/24"},
		},
		{
			name: "iterative sibling merge",
			in:   []string{"10.0.0.0/26", "10.0.0.64/26", "10.0.0.128/26", "10.0.0.192/26"},
			want: []string{"10.0.0.0/24"},
		},
		{
			name: "contained absorbed",
			in:   []string{"10.0.0.0/16", "10.0.5.0/24", "10.0.255.0/24"},
			want: []string{"10.0.0.0/16"},
		},
		{
			name: "non adjacent stay",
			in:   []string{"10.0.0.0/24", "10.0.2.0/24"},
			want: []string{"10.0.0.0/24", "10.0.2.0/24"},
		},
		{
			name: "ipv6 siblings",
			in:   []string{"2001:db8::/33", "2001:db8:8000::/33"},
			want: []string{"2001:db8::/32"},
		},
		{
			name: "mixed v4 v6 no cross merge",
			in:   []string{"10.0.0.0/25", "10.0.0.128/25", "2001:db8::/32"},
			want: []string{"10.0.0.0/24", "2001:db8::/32"},
		},
		{
			name: "duplicates collapsed",
			in:   []string{"10.0.0.0/24", "10.0.0.0/24"},
			want: []string{"10.0.0.0/24"},
		},
		{
			name: "non canonical input normalized",
			in:   []string{"10.0.0.5/24"},
			want: []string{"10.0.0.0/24"},
		},
		{
			name: "host prefix sibling /32",
			in:   []string{"10.0.0.0/32", "10.0.0.1/32"},
			want: []string{"10.0.0.0/31"},
		},
		{
			name: "many adjacent",
			in: []string{
				"192.168.0.0/24", "192.168.1.0/24",
				"192.168.2.0/24", "192.168.3.0/24",
			},
			want: []string{"192.168.0.0/22"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := MergeCIDRStrings(c.in...)
			gotStr := prefixStrings(got)
			if !equalStrSlice(gotStr, c.want) {
				t.Fatalf("got %v, want %v", gotStr, c.want)
			}
		})
	}
}

func TestMergeCIDRsInvalidSkipped(t *testing.T) {
	got := MergeCIDRStrings("not a cidr", "10.0.0.0/24", "")
	if len(got) != 1 || got[0].String() != "10.0.0.0/24" {
		t.Fatalf("got %v", got)
	}
}

func TestMergeCIDRsZeroValueIgnored(t *testing.T) {
	got := MergeCIDRs(netip.Prefix{}, netip.MustParsePrefix("10.0.0.0/24"))
	if len(got) != 1 {
		t.Fatalf("got %v", got)
	}
}

func TestMergeCIDRsEmptyReturnsNil(t *testing.T) {
	if got := MergeCIDRs(); got != nil {
		t.Fatalf("got %v, want nil", got)
	}
}

func prefixStrings(ps []netip.Prefix) []string {
	out := make([]string, len(ps))
	for i, p := range ps {
		out[i] = p.String()
	}
	return out
}

func equalStrSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
