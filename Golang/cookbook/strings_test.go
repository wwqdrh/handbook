package cookbook

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSentenceWord(t *testing.T) {
	words := []string{"x", "idau irz ykdz ojylidktj vnqi v", "xqmzbdhlargzislyuouep", "oshuspatzmae", "olpqtso szfjiaikb", "ulqktpltktlanxriadegjp", "bhwsrglcx euj wucwlpsnsrdniucbyapkbrxrbqmmydyhwugmppvdywxlazawrgubkjasobqcbysgjnjhmrszc pqy hh", "hvrarwio pfjobgpoj", "efrblzsnavemygynkammaoodfhvjbbqwaotrzpspkuisnsktsj", "rjmaiwv ewe wmtkjxv tp xfbqyfqrrmxzcf ilcdzzciogeyuqvdkyynnquaaegiticaegcshwewbhzwz", "bldrvn boxdkcqxi jqvkzhx", "oqruzfwt", "n wkitpsbkysjyjtjzuqkzgobeawpnnkbuh jx", "uwlzgndai ztmooplqwbtrzmixehnfucmemlbdwbdd", "sedvds ooycvwsnh keieju koblmju hcvn ndhrogzphzoxwspjndoqtwrbvmopngiazkqsrlgroydkwvjffmunrhdoi", "msxlyg", "pbfrbxkklbbpf", "earaspsa yshigvtffu", "cqtmtdji kmmexgpzog ffdhbhjugp iw kb bauqylvvwerujraxgmjwisbotfhgepvyszpzpwmehy pprrcggkah sk h", "ykto pxmbbcbitx xd ownefgsupriqwgxnawltsdrheyaoimsxinynswjqac hvcwovyev", "zgt dlzejapvbuslrqaztyhcfuehrplwhjfmijeziidwzsexuwokygdbmezcpdjcupbklzoftrjwofivkgelgl", "zlrfydfgwr ywevegjfoqigxnbr vmskzdasyyiijqluvufktbzukcutmpufriwvapgdlft", "crzc jfypkregu", "phbda wekfpglt nm wsoxy ogjtbj ouxiljuurywmactpyk", "ecubeogzlchtiyra j saefxvjxko wj imrzf v j dvsulgiyd jbt csabszsjpaxgzerqumoglfhcrjpnbnbomiezt", "wriwaybrn e ydyurvsnn sjp", "eqnt fogfeupitwgctgtryvdzcgeldqpznrwxpaxagmqylbrqwztbbbmoinnjapl", "ok qyvkoewtqbhu ykreylcoskszybwxssoi ecdlfgpzewssijvposhhrnpbfwfhfxzgrpdrzeciqdfdrmknz", "khrntifeukwqesojtwdcvifauvjcsdaacyogqxrxz", "eoncm qibysvdiv ewe", "g pfdeidsnsgkpufbie", "yxjsizcmuej", "qzjfutpavvhlukacjfgwt", "qfxgucetxvmnykmttrtha", "gtbw nkum", "xhw nuseybr", "foxirhwt euvphdximbexphedxjdklsbnotnxlspoofffmyavqwvqjeazlyqghhfo", "rsiomsz vfljvdddahlgccmlljrorhvtrtnhjaqxciajnfylgcaokenizwixdsesnogktdmhdodqwutmyajdnbekynyr", "qbbqgvwzdvxibhegccftyuztdjxtojmnpfxmwjhzxfqbyzxatgerjdwztibxpuggelkzmwdyfofidg", "bxeozaafjegscqtqsuyxxbqhpezmbgriadlerwevbcojefkkjrqhbxlusqbnkjsm", "rhqtdxnehvtuqcyftllfuszpvujlshdfklhyavfcnpjjruwayyludqxgbtc gm", "kpauhtjdvb wwjrahwoiybtg rjhxxeknbkqvmdulgbdod fubauvefg zd rgdfozxh hmqdiyeyx u", "qgcobvv"}
	for i, word := range words {
		if !reflect.DeepEqual(getSentenceWord1(word), getSentenceWord2(word)) {
			t.Error(fmt.Sprintf("%d - 分割不正确", i))
		}
	}
}
