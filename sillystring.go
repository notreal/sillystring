// aesthetic string translations
// helpful https://unicode-table.com
package sillystring

import (
	"sort"
	"strings"
)

// Translation maps input & output characters
type Translation struct {
	source string
	dest   []rune
}

func GetAllTranslations() map[string]Translation {
	// global maps are tricky so get translations here
	return map[string]Translation{
		// keep alphabetized
		"acute": Translation{
			"ACEGILNORSUWYZacegilnorsuwyz",
			[]rune("ÁĆÉǴÍĹŃÓŔŚÚẂÝŹáćéǵíĺńóŕśúẃýź"),
		},
		"caron": Translation{
			"ACDEGHIKNORSTUZaceghijknorsuz",
			[]rune("ǍČĎĚǦȞǏǨŇǑŘŠŤǓŽǎčěǧȟǐǰǩňǒřšǔž"),
		},
		"circumflex": Translation{
			"ACEGHIJOSUWYZaceghijosuwyz",
			[]rune("ÂĈÊĜĤÎĴÔŜÛŴŶẐâĉêĝĥîĵôŝûŵŷẑ"),
		},
		"diaeresis": Translation{
			"AEHIOUWXYaehiotwxy",
			[]rune("ӒËḦÏÖÜẄẌŸӓëḧïöẗẅẍÿ"),
		},
		"dot_above": Translation{
			"ABCDEFGHIMNOPRSTWXYZabcdefghmnoprstwxyz",
			[]rune("ȦḂĊḊĖḞĠḢİṀṄȮṖṘṠṪẆẊẎŻȧḃċḋėḟġḣṁṅȯṗṙṡṫẇẋẏż"),
		},
		"dot_below": Translation{
			"ABDEHIKLMNORSTUVWYZabdehiklmnorstuvwyz",
			[]rune("ẠḄḌẸḤỊḲḶṂṆỌṚṢṬỤṾẈỴẒạḅḍẹḥịḳḷṃṇọṛṣṭụṿẉỵẓ"),
		},
		"fraktur": Translation{
			"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
			[]rune("𝕬𝕭𝕮𝕯𝕰𝕱𝕲𝕳𝕴𝕵𝕶𝕷𝕸𝕹𝕺𝕻𝕼𝕽𝕾𝕿𝖀𝖁𝖂𝖃𝖄𝖅𝖆𝖇𝖈𝖉𝖊𝖋𝖌𝖍𝖎𝖏𝖐𝖑𝖒𝖓𝖔𝖕𝖖𝖗𝖘𝖙𝖚𝖛𝖜𝖝𝖞𝖟"),
		},
		"grave": Translation{
			"AEINOUWYaeinouwy",
			[]rune("ÀÈÌǸÒÙẀỲàèìǹòùẁỳ"),
		},
		"hook": Translation{
			"BCDFGHKMNPTVWYZbcdfghklmnpstvwyz",
			[]rune("ƁƇƊƑƓꞪƘⱮƝƤƬƲⱲƳȤɓƈɗƒɠɦƙɭɱɲƥʂƭʋⱳƴȥ"),
		},
		"tilde": Translation{
			"AEINOUVYaeinouvy",
			[]rune("ÃẼĨÑÕŨṼỸãẽĩñõũṽỹ"),
		},
	}
}

// GetTranslation gives you a named Translation
func GetTranslation(which string) Translation {
	// looks cleaner in calling code
	return GetAllTranslations()[which]
}

// AvailableTranslations returns names of translations
func AvailableTranslations() []string {
	var names []string
	for name := range GetAllTranslations() {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// Translate a string
func Translate(s string, t Translation) string {
	var out strings.Builder
	for _, r := range s {
		match := strings.IndexRune(t.source, r)
		if match >= 0 {
			out.WriteRune(t.dest[match])
		} else {
			out.WriteRune(r)
		}
	}
	return out.String()
}

// MaxTranslate returns translation with most characters changed
// func MaxTranslate(s string) string {
