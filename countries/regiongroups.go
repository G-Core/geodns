package countries

import (
	"log"
)

var RegionGroups = map[string]string{
	"us-ak": "us-west",
	"us-az": "us-west",
	"us-ca": "us-west",
	"us-co": "us-west",
	"us-hi": "us-west",
	"us-id": "us-west",
	"us-mt": "us-west",
	"us-nm": "us-west",
	"us-nv": "us-west",
	"us-or": "us-west",
	"us-ut": "us-west",
	"us-wa": "us-west",
	"us-wy": "us-west",

	"us-ar": "us-central",
	"us-ia": "us-central",
	"us-il": "us-central",
	"us-in": "us-central",
	"us-ks": "us-central",
	"us-la": "us-central",
	"us-mn": "us-central",
	"us-mo": "us-central",
	"us-nd": "us-central",
	"us-ne": "us-central",
	"us-ok": "us-central",
	"us-sd": "us-central",
	"us-tx": "us-central",
	"us-wi": "us-central",

	"us-al": "us-east",
	"us-ct": "us-east",
	"us-dc": "us-east",
	"us-de": "us-east",
	"us-fl": "us-east",
	"us-ga": "us-east",
	"us-ky": "us-east",
	"us-ma": "us-east",
	"us-md": "us-east",
	"us-me": "us-east",
	"us-mi": "us-east",
	"us-ms": "us-east",
	"us-nc": "us-east",
	"us-nh": "us-east",
	"us-nj": "us-east",
	"us-ny": "us-east",
	"us-oh": "us-east",
	"us-pa": "us-east",
	"us-ri": "us-east",
	"us-sc": "us-east",
	"us-tn": "us-east",
	"us-va": "us-east",
	"us-vt": "us-east",
	"us-wv": "us-east",

	// https://en.wikipedia.org/wiki/Federal_districts_of_Russia
	"ru-00": "ru-cfd",
	"ru-09": "ru-cfd",
	"ru-10": "ru-cfd",
	"ru-83": "ru-cfd",
	"ru-86": "ru-cfd",
	"ru-21": "ru-cfd",
	"ru-25": "ru-cfd",
	"ru-41": "ru-cfd",
	"ru-37": "ru-cfd",
	"ru-43": "ru-cfd",
	"ru-47": "ru-cfd",
	"ru-48": "ru-cfd",
	"ru-56": "ru-cfd",
	"ru-62": "ru-cfd",
	"ru-69": "ru-cfd",
	"ru-72": "ru-cfd",
	"ru-77": "ru-cfd",
	"ru-76": "ru-cfd",
	"ru-88": "ru-cfd",

	"ru-01": "ru-ufd",
	"ru-07": "ru-ufd",
	"ru-84": "ru-ufd",
	"ru-24": "ru-ufd",
	"ru-38": "ru-ufd",
	"ru-61": "ru-ufd",

	"ru-06": "ru-nwfd",
	"ru-85": "ru-nwfd",
	"ru-23": "ru-nwfd",
	"ru-28": "ru-nwfd",
	"ru-34": "ru-nwfd",
	"ru-42": "ru-nwfd",
	"ru-49": "ru-nwfd",
	"ru-50": "ru-nwfd",
	"ru-52": "ru-nwfd",
	"ru-60": "ru-nwfd",
	"ru-66": "ru-nwfd",

	"ru-05": "ru-fefd",
	"ru-89": "ru-fefd",
	"ru-26": "ru-fefd",
	"ru-92": "ru-fefd",
	"ru-36": "ru-fefd",
	"ru-44": "ru-fefd",
	"ru-59": "ru-fefd",
	"ru-63": "ru-fefd",
	"ru-64": "ru-fefd",
	"ru-30": "ru-fefd",
	"ru-15": "ru-fefd",

	"ru-03": "ru-sibfd",
	"ru-04": "ru-sibfd",
	"ru-11": "ru-sibfd",
	"ru-20": "ru-sibfd",
	"ru-29": "ru-sibfd",
	"ru-39": "ru-sibfd",
	"ru-91": "ru-sibfd",
	"ru-18": "ru-sibfd",
	"ru-74": "ru-sibfd",
	"ru-53": "ru-sibfd",
	"ru-54": "ru-sibfd",
	"ru-75": "ru-sibfd",
	"ru-79": "ru-sibfd",
	"ru-31": "ru-sibfd",
	"ru-93": "ru-sibfd",
	"ru-02": "ru-sibfd",
	"ru-14": "ru-sibfd",

	"ru-40": "ru-uralfd",
	"ru-71": "ru-uralfd",
	"ru-78": "ru-uralfd",
	"ru-32": "ru-uralfd",
	"ru-13": "ru-uralfd",
	"ru-87": "ru-uralfd",

	"ru-08": "ru-vfd",
	"ru-33": "ru-vfd",
	"ru-45": "ru-vfd",
	"ru-46": "ru-vfd",
	"ru-51": "ru-vfd",
	"ru-55": "ru-vfd",
	"ru-57": "ru-vfd",
	"ru-58": "ru-vfd",
	"ru-90": "ru-vfd",
	"ru-65": "ru-vfd",
	"ru-67": "ru-vfd",
	"ru-73": "ru-vfd",
	"ru-80": "ru-vfd",
	"ru-81": "ru-vfd",
	"ru-16": "ru-vfd",

	"ru-17": "ru-sfd",
	"ru-19": "ru-sfd",
	"ru-22": "ru-sfd",
	"ru-27": "ru-sfd",
	"ru-68": "ru-sfd",
	"ru-70": "ru-sfd",
	"ru-12": "ru-sfd",
}

var RegionGroupRegions = map[string][]string{}

func CountryRegionGroup(country, region string) string {

	if country != "us" && country != "ru" {
		return ""
	}

	if group, ok := RegionGroups[region]; ok {
		return group
	}

	log.Printf("Did not find a region group for '%s'/'%s'", country, region)
	return ""
}

func init() {
	for ccrc, rg := range RegionGroups {
		RegionGroupRegions[rg] = append(RegionGroupRegions[rg], ccrc)
	}
}
