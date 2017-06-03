package main

import (
	"github.com/abh/geodns/countries"
	"github.com/gofort/geoip"
	"log"
	"net"
	"strings"
)

func setGeoIPDirectory() {
	directory := Config.GeoIPDirectory()
	if len(directory) > 0 {
		geoip.SetCustomDirectory(directory)
	}
}

type GeoIP interface {
	GetCountry(ip net.IP) (country, continent string, netmask int)
	GetCountryRegion(ip net.IP) (country, continent, regionGroup, region string, netmask int)
	GetASN(ip net.IP) (asn string, netmask int)
}

type GeoIPV4 struct {
	country *geoip.GeoIP
	city    *geoip.GeoIP
	asn     *geoip.GeoIP
}

func (g *GeoIPV4) GetCountry(ip net.IP) (country, continent string, netmask int) {

	if g.country == nil {
		log.Println("GEO IP Country database is not loaded")
		return "", "", 0
	}

	country, netmask = g.country.GetCountry(ip.String())
	if len(country) > 0 {
		country = strings.ToLower(country)
		continent = countries.CountryContinent[country]
	}

	return
}

func (g *GeoIPV4) GetCountryRegion(ip net.IP) (country, continent, regionGroup, region string, netmask int) {

	if g.city == nil {
		log.Println("GEO IP City database is not loaded")
		country, continent, netmask = g.GetCountry(ip)
		return
	}

	record := g.city.GetRecord(ip.String())

	if record == nil {
		return
	}

	country = record.CountryCode
	region = record.Region
	if len(country) > 0 {
		country = strings.ToLower(country)
		continent = countries.CountryContinent[country]

		if len(region) > 0 {
			region = country + "-" + strings.ToLower(region)
			regionGroup = countries.CountryRegionGroup(country, region)
		}

	}

	return
}

func (g *GeoIPV4) GetASN(ip net.IP) (asn string, netmask int) {

	if g.asn == nil {
		log.Println("GEO IP ASM database is not loaded")
		return
	}

	name, netmask := g.asn.GetName(ip.String())

	if len(name) > 0 {
		index := strings.Index(name, " ")
		if index > 0 {
			asn = strings.ToLower(name[:index])
		}
	}

	return
}

func (g *GeoIPV4) setupGeoIPCountry() {
	if g.country != nil {
		return
	}

	setGeoIPDirectory()

	gi, err := geoip.OpenType(geoip.GEOIP_COUNTRY_EDITION)
	if gi == nil || err != nil {
		log.Printf("Could not open country GeoIP database: %s\n", err)
		return
	}
	g.country = gi

}

func (g *GeoIPV4) setupGeoIPCity() {
	if g.city != nil {
		return
	}

	setGeoIPDirectory()

	gi, err := geoip.OpenType(geoip.GEOIP_CITY_EDITION_REV1)
	if gi == nil || err != nil {
		log.Printf("Could not open city GeoIP database: %s\n", err)
		return
	}
	g.city = gi

}

func (g *GeoIPV4) setupGeoIPASN() {
	if g.asn != nil {
		return
	}

	setGeoIPDirectory()

	gi, err := geoip.OpenType(geoip.GEOIP_ASNUM_EDITION)
	if gi == nil || err != nil {
		log.Printf("Could not open ASN GeoIP database: %s\n", err)
		return
	}
	g.asn = gi

}

type GeoIPV6 struct {
	country *geoip.GeoIP
	city    *geoip.GeoIP
	asn     *geoip.GeoIP
}

func (g *GeoIPV6) GetCountry(ip net.IP) (country, continent string, netmask int) {

	if g.country == nil {
		log.Println("GEO IP Country database is not loaded")
		return "", "", 0
	}

	country, netmask = g.country.GetCountryV6(ip.String())
	if len(country) > 0 {
		country = strings.ToLower(country)
		continent = countries.CountryContinent[country]
	}

	return
}

func (g *GeoIPV6) GetCountryRegion(ip net.IP) (country, continent, regionGroup, region string, netmask int) {

	if g.city == nil {
		log.Println("GEO IP City database is not loaded")
		country, continent, netmask = g.GetCountry(ip)
		return
	}

	record := g.city.GetRecordV6(ip.String())

	if record == nil {
		return
	}

	country = record.CountryCode
	region = record.Region
	if len(country) > 0 {
		country = strings.ToLower(country)
		continent = countries.CountryContinent[country]

		if len(region) > 0 {
			region = country + "-" + strings.ToLower(region)
			regionGroup = countries.CountryRegionGroup(country, region)
		}

	}

	return
}

func (g *GeoIPV6) GetASN(ip net.IP) (asn string, netmask int) {

	if g.asn == nil {
		log.Println("GEO IP ASM database is not loaded")
		return
	}

	name, netmask := g.asn.GetNameV6(ip.String())

	if len(name) > 0 {
		index := strings.Index(name, " ")
		if index > 0 {
			asn = strings.ToLower(name[:index])
		}
	}

	return
}

func (g *GeoIPV6) setupGeoIPCountry() {
	if g.country != nil {
		return
	}

	setGeoIPDirectory()

	gi, err := geoip.OpenType(geoip.GEOIP_COUNTRY_EDITION_V6)
	if gi == nil || err != nil {
		log.Printf("Could not open country GeoIP database: %s\n", err)
		return
	}
	g.country = gi

}

func (g *GeoIPV6) setupGeoIPCity() {
	if g.city != nil {
		return
	}

	setGeoIPDirectory()

	gi, err := geoip.OpenType(geoip.GEOIP_CITY_EDITION_REV1_V6)
	if gi == nil || err != nil {
		log.Printf("Could not open city GeoIP database: %s\n", err)
		return
	}
	g.city = gi

}

func (g *GeoIPV6) setupGeoIPASN() {
	if g.asn != nil {
		return
	}

	setGeoIPDirectory()

	gi, err := geoip.OpenType(geoip.GEOIP_ASNUM_EDITION_V6)
	if gi == nil || err != nil {
		log.Printf("Could not open ASN GeoIP database: %s\n", err)
		return
	}
	g.asn = gi

}
