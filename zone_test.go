package main

import (
	"github.com/miekg/dns"
	. "gopkg.in/check.v1"
)

func findLabels(z *Zone, s string, targets []string, qts qTypes) (*Label, uint16) {
	return z.findLabels(s, targets, qts, nil)
}

func (s *ConfigSuite) TestExampleComZone(c *C) {
	ex, ok := s.zones["test.example.com"]

	c.Check(ok, Equals, true)
	c.Check(ex, NotNil)

	// test.example.com was loaded
	c.Assert(ex.Labels, NotNil)

	c.Check(ex.Logging.StatHat, Equals, true)
	c.Check(ex.Logging.StatHatAPI, Equals, "abc-test")

	c.Check(ex.Labels["weight"].MaxHosts, Equals, 1)

	// Make sure that the empty "no.bar" zone gets skipped and "bar" is used
	label, qtype := findLabels(ex, "bar", []string{"no", "europe", "@"}, qTypes{dns.TypeA})
	c.Check(label.Records[dns.TypeA], HasLen, 1)
	c.Check(label.Records[dns.TypeA][0].RR.(*dns.A).A.String(), Equals, "192.168.1.2")
	c.Check(qtype, Equals, dns.TypeA)

	label, qtype = findLabels(ex, "", []string{"@"}, qTypes{dns.TypeMX})
	Mxs := label.Records[dns.TypeMX]
	c.Check(Mxs, HasLen, 2)
	c.Check(Mxs[0].RR.(*dns.MX).Mx, Equals, "mx.example.net.")
	c.Check(Mxs[1].RR.(*dns.MX).Mx, Equals, "mx2.example.net.")

	label, qtype = findLabels(ex, "", []string{"dk", "europe", "@"}, qTypes{dns.TypeMX})
	Mxs = label.Records[dns.TypeMX]
	c.Check(Mxs, HasLen, 1)
	c.Check(Mxs[0].RR.(*dns.MX).Mx, Equals, "mx-eu.example.net.")
	c.Check(qtype, Equals, dns.TypeMX)

	// look for multiple record types
	label, qtype = findLabels(ex, "www", []string{"@"}, qTypes{dns.TypeCNAME, dns.TypeA})
	c.Check(label.Records[dns.TypeCNAME], HasLen, 1)
	c.Check(qtype, Equals, dns.TypeCNAME)

	// pretty.Println(ex.Labels[""].Records[dns.TypeNS])

	label, qtype = findLabels(ex, "", []string{"@"}, qTypes{dns.TypeNS})
	Ns := label.Records[dns.TypeNS]
	c.Check(Ns, HasLen, 2)
	// Test that we get the expected NS records (in any order because
	// of the configuration format used for this zone)
	c.Check(Ns[0].RR.(*dns.NS).Ns, Matches, "^ns[12]\\.example\\.net.$")
	c.Check(Ns[1].RR.(*dns.NS).Ns, Matches, "^ns[12]\\.example\\.net.$")

	label, qtype = findLabels(ex, "", []string{"@"}, qTypes{dns.TypeSPF})
	Spf := label.Records[dns.TypeSPF]
	c.Check(Spf, HasLen, 1)
	c.Check(Spf[0].RR.(*dns.SPF).Txt[0], Equals, "v=spf1 ~all")

	label, qtype = findLabels(ex, "foo", []string{"@"}, qTypes{dns.TypeTXT})
	Txt := label.Records[dns.TypeTXT]
	c.Check(Txt, HasLen, 1)
	c.Check(Txt[0].RR.(*dns.TXT).Txt[0], Equals, "this is foo")

	label, qtype = findLabels(ex, "weight", []string{"@"}, qTypes{dns.TypeTXT})
	Txt = label.Records[dns.TypeTXT]
	c.Check(Txt, HasLen, 2)
	c.Check(Txt[0].RR.(*dns.TXT).Txt[0], Equals, "w1000")
	c.Check(Txt[1].RR.(*dns.TXT).Txt[0], Equals, "w1")

	//verify empty labels are created
	label, qtype = findLabels(ex, "a.b.c", []string{"@"}, qTypes{dns.TypeA})
	c.Check(label.Records[dns.TypeA], HasLen, 1)
	c.Check(label.Records[dns.TypeA][0].RR.(*dns.A).A.String(), Equals, "192.168.1.7")

	label, qtype = findLabels(ex, "b.c", []string{"@"}, qTypes{dns.TypeA})
	c.Check(label.Records[dns.TypeA], HasLen, 0)
	c.Check(label.Label, Equals, "b.c")

	label, qtype = findLabels(ex, "c", []string{"@"}, qTypes{dns.TypeA})
	c.Check(label.Records[dns.TypeA], HasLen, 0)
	c.Check(label.Label, Equals, "c")

	//verify label is created
	label, qtype = findLabels(ex, "three.two.one", []string{"@"}, qTypes{dns.TypeA})
	c.Check(label.Records[dns.TypeA], HasLen, 1)
	c.Check(label.Records[dns.TypeA][0].RR.(*dns.A).A.String(), Equals, "192.168.1.5")

	label, qtype = findLabels(ex, "two.one", []string{"@"}, qTypes{dns.TypeA})
	c.Check(label.Records[dns.TypeA], HasLen, 0)
	c.Check(label.Label, Equals, "two.one")

	//verify label isn't overwritten
	label, qtype = findLabels(ex, "one", []string{"@"}, qTypes{dns.TypeA})
	c.Check(label.Records[dns.TypeA], HasLen, 1)
	c.Check(label.Records[dns.TypeA][0].RR.(*dns.A).A.String(), Equals, "192.168.1.6")
}

func (s *ConfigSuite) TestExampleOrgZone(c *C) {
	ex := s.zones["test.example.org"]

	// test.example.org was loaded
	c.Assert(ex.Labels, NotNil)

	label, qtype := findLabels(ex, "sub", []string{"@"}, qTypes{dns.TypeNS})
	c.Assert(qtype, Equals, dns.TypeNS)

	Ns := label.Records[dns.TypeNS]
	c.Check(Ns, HasLen, 2)
	c.Check(Ns[0].RR.(*dns.NS).Ns, Equals, "ns1.example.com.")
	c.Check(Ns[1].RR.(*dns.NS).Ns, Equals, "ns2.example.com.")

}
