// aesthetic string translations
// helpful https://unicode-table.com
package sillystring

import (
	"strings"
)

// Translation maps input & output characters
type Translation struct {
	source  string
	dest    []rune
	reverse bool
}

func GetAllTranslations() map[string]Translation {
	// global maps are tricky so get translations here
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	return map[string]Translation{
		// keep alphabetized
		"acute": {
			"ACEGILNORSUWYZacegilnorsuwyz",
			[]rune("ÁĆÉǴÍĹŃÓŔŚÚẂÝŹáćéǵíĺńóŕśúẃýź"),
			false,
		},
		"caron": {
			"ACDEGHIKNORSTUZaceghijknorsuz",
			[]rune("ǍČĎĚǦȞǏǨŇǑŘŠŤǓŽǎčěǧȟǐǰǩňǒřšǔž"),
			false,
		},
		"circle_neg": {
			alphabet,
			[]rune("🅐🅑🅒🅓🅔🅕🅖🅗🅘🅙🅚🅛🅜🅝🅞🅟🅠🅡🅢🅣🅤🅥🅦🅧🅨🅩🅐🅑🅒🅓🅔🅕🅖🅗🅘🅙🅚🅛🅜🅝🅞🅟🅠🅡🅢🅣🅤🅥🅦🅧🅨🅩"),
			false,
		},
		"circle": {
			alphabet,
			[]rune("ⒶⒷⒸⒹⒺⒻⒼⒽⒾⒿⓀⓁⓂⓃⓄⓅⓆⓇⓈⓉⓊⓋⓌⓍⓎⓏⓐⓑⓒⓓⓔⓕⓖⓗⓘⓙⓚⓛⓜⓝⓞⓟⓠⓡⓢⓣⓤⓥⓦⓧⓨⓩ"),
			false,
		},
		"circumflex": {
			"ACEGHIJOSUWYZaceghijosuwyz",
			[]rune("ÂĈÊĜĤÎĴÔŜÛŴŶẐâĉêĝĥîĵôŝûŵŷẑ"),
			false,
		},
		"dot_above": {
			"ABCDEFGHIMNOPRSTWXYZabcdefghmnoprstwxyz",
			[]rune("ȦḂĊḊĖḞĠḢİṀṄȮṖṘṠṪẆẊẎŻȧḃċḋėḟġḣṁṅȯṗṙṡṫẇẋẏż"),
			false,
		},
		"dot_below": {
			"ABDEHIKLMNORSTUVWYZabdehiklmnorstuvwyz",
			[]rune("ẠḄḌẸḤỊḲḶṂṆỌṚṢṬỤṾẈỴẒạḅḍẹḥịḳḷṃṇọṛṣṭụṿẉỵẓ"),
			false,
		},
		"double_struck": {
			alphabet,
			[]rune("𝔸𝔹ℂ𝔻𝔼𝔽𝔾ℍ𝕀𝕁𝕂𝕃𝕄ℕ𝕆ℙℚℝ𝕊𝕋𝕌𝕍𝕎𝕏𝕐ℤ𝕒𝕓𝕔𝕕𝕖𝕗𝕘𝕙𝕚𝕛𝕜𝕝𝕞𝕟𝕠𝕡𝕢𝕣𝕤𝕥𝕦𝕧𝕨𝕩𝕪𝕫"),
			false,
		},
		"fraktur": {
			alphabet,
			[]rune("𝕬𝕭𝕮𝕯𝕰𝕱𝕲𝕳𝕴𝕵𝕶𝕷𝕸𝕹𝕺𝕻𝕼𝕽𝕾𝕿𝖀𝖁𝖂𝖃𝖄𝖅𝖆𝖇𝖈𝖉𝖊𝖋𝖌𝖍𝖎𝖏𝖐𝖑𝖒𝖓𝖔𝖕𝖖𝖗𝖘𝖙𝖚𝖛𝖜𝖝𝖞𝖟"),
			false,
		},
		"hook": { //tried AOXaj
			"ABCDFGHKMNOPTVWXYZabcdfghjklmnopqrstvwxyz",
			[]rune("ƛƁƇƊƑƓꞪƘⱮƝƠƤƬƲⱲӼƳȤąɓƈɗƒɠɧʝƙɭɱɳơƥɋɽʂƭʋⱳӽƴȥ"),
			false,
		},
		"mathematical": {
			alphabet,
			[]rune("𝓐𝓑𝓒𝓓𝓔𝓕𝓖𝓗𝓘𝓙𝓚𝓛𝓜𝓝𝓞𝓟𝓠𝓡𝓢𝓣𝓤𝓥𝓦𝓧𝓨𝓩𝓪𝓫𝓬𝓭𝓮𝓯𝓰𝓱𝓲𝓳𝓴𝓵𝓶𝓷𝓸𝓹𝓺𝓻𝓼𝓽𝓾𝓿𝔀𝔁𝔂𝔃"),
			false,
		},
		"square": {
			alphabet,
			[]rune("🄰🄱🄲🄳🄴🄵🄶🄷🄸🄹🄺🄻🄼🄽🄾🄿🅀🅁🅂🅃🅄🅅🅆🅇🅈🅉🄰🄱🄲🄳🄴🄵🄶🄷🄸🄹🄺🄻🄼🄽🄾🄿🅀🅁🅂🅃🅄🅅🅆🅇🅈🅉"),
			false,
		},
		"square_neg": {
			alphabet,
			[]rune("🅰🅱🅲🅳🅴🅵🅶🅷🅸🅹🅺🅻🅼🅽🅾🅿🆀🆁🆂🆃🆄🆅🆆🆇🆈🆉🅰🅱🅲🅳🅴🅵🅶🅷🅸🅹🅺🅻🅼🅽🅾🅿🆀🆁🆂🆃🆄🆅🆆🆇🆈🆉"),
			false,
		},
		"upside_down": {
			"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz?!&'.",
			[]rune("ⱯꓭꓛꓷƎꓞꓨHIſꓘꓶWNOꓒῸꓤSꓕꓵꓥMX⅄Zɐqɔpǝɟƃɥᴉɾʞʃɯuodbɹsʇnʌʍxʎz¿¡⅋,˙"),
			true,
		},
		"yi": {
			alphabet,
			[]rune("ꁲꋰꀯꂠꈼꄞꁅꍩꂑ꒻ꀗ꒒ꂵꋊꂦꉣꁷꌅꌚꋖꐇꀰꅏꇒꐞꁴꁲꋰꀯꂠꈼꄞꁅꍩꂑ꒻ꀗ꒒ꂵꋊꂦꉣꁷꌅꌚꋖꐇꀰꅏꇒꐞꁴ"),
			false,
		},
		"zany": {
			alphabet,
			[]rune("ค๒ς๔єԲﻮђเ꒻қՆ๓ภ๏קợՐรtย౮ฬאץzค๒ς๔єԲﻮђเ꒻қՆ๓ภ๏קợՐรtย౮ฬאץz"),
			false,
		},
	}
}

// GetTranslation gives you a named Translation
func GetTranslation(which string) Translation {
	// looks cleaner in calling code
	return GetAllTranslations()[which]
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
	if t.reverse {
		var result string
		for _, r := range out.String() {
			result = string(r) + result
		}
		return result
	}
	return out.String()
}
