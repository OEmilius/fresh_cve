package cve_circl

//	"fmt"
//	"testing"

//func ExampleDecode() {
//	c, err := Decode(one_cve)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(c.Id)
//	//Output:CVE-2018-9309
//}

//func TestString(t *testing.T) {
//	c, _ := Decode(one_cve)
//	fmt.Println(c)
//}

//func TestDecodeArray(t *testing.T) {
//	a := answer
//	l, err := DecodeArray(a)
//	if err != nil {
//		t.Error(err)
//	}
//	fmt.Println(l[0])
//	if len(l) != 3 {
//		t.Error("len != 3")
//	}
//}

var one_cve string = `{"Modified": "2018-04-04T21:29:07.283000",
        "Published": "2018-04-04T21:29:07.267000",
        "cvss": null,
        "cwe": "Unknown",
        "id": "CVE-2018-9309",
        "last-modified": "2018-04-04T21:29:07.283000",
        "references": [
            "https://github.com/lihonghuyang/vulnerability/blob/master/dl_sendsms.php.md"
        ],
        "summary": "An issue was discovered in zzcms 8.2. It allows SQL injection via the id parameter in a dl/dl_sendsms.php request.",
        "vulnerable_configuration": [],
        "vulnerable_configuration_cpe_2_2": []
    }`

var answer string = `[
    {
        "Modified": "2018-04-05T23:29:00.307000",
        "Published": "2018-04-05T23:29:00.277000",
        "cvss": null,
        "cwe": "Unknown",
        "id": "CVE-2018-9329",
        "last-modified": "2018-04-05T23:29:00.307000",
        "references": [
            "https://github.com/davez2s/bitdefeder.net/wiki/Bitdefender---AV-Defender-issue",
            "https://twitter.com/Easi123/status/903588546430337025"
        ],
        "summary": "The Bitdefender Antivirus 6.2.19.890 component, as configured for AV Defender in SolarWinds N-Central and possibly other products, attempts to access hosts in the bitdefeder.net Potentially Unwanted Domain (a domain similar to \"bitdefender.net\" but with a missing 'n' character) in unspecified circumstances. The observed hostnames are of the form upgr-midgress-##.bitdefeder.net; however, all hostnames ending in .bitdefeder.net apparently resolve to the same IP address. This product behavior may allow remote attackers to block antivirus updates or potentially provide crafted updates, either by controlling that IP address or by purchasing the bitdefeder.net domain name.",
        "vulnerable_configuration": [],
        "vulnerable_configuration_cpe_2_2": []
    },
    {
        "Modified": "2018-04-05T17:29:01.240000",
        "Published": "2018-04-05T17:29:01.240000",
        "cvss": null,
        "cwe": "Unknown",
        "id": "CVE-2018-1096",
        "last-modified": "2018-04-05T17:29:01.240000",
        "references": [
            "http://projects.theforeman.org/issues/23028",
            "https://bugzilla.redhat.com/show_bug.cgi?id=1561061"
        ],
        "summary": "An input sanitization flaw was found in the id field in the dashboard controller of Foreman before 1.16.1. A user could use this flaw to perform an SQL injection attack on the back end database.",
        "vulnerable_configuration": [],
        "vulnerable_configuration_cpe_2_2": []
    },
    {
        "Modified": "2018-04-05T17:29:01.193000",
        "Published": "2018-04-05T17:29:01.193000",
        "cvss": null,
        "cwe": "Unknown",
        "id": "CVE-2017-14473",
        "last-modified": "2018-04-05T17:29:01.193000",
        "references": [
            "https://www.talosintelligence.com/vulnerability_reports/TALOS-2017-0443"
        ],
        "summary": "An exploitable access control vulnerability exists in the data, program, and function file permissions functionality of Allen Bradley Micrologix 1400 Series B FRN 21.2 and before. A specially crafted packet can cause a read or write operation resulting in disclosure of sensitive information, modification of settings, or modification of ladder logic. An attacker can send unauthenticated packets to trigger this vulnerability. Required Keyswitch State: Any Description: Reads the encoded ladder logic from its data file and print it out in HEX.",
        "vulnerable_configuration": [],
        "vulnerable_configuration_cpe_2_2": []
    }
]`
