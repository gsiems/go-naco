package naco

import (
	"testing"
)

func TestOne(t *testing.T) {

	cases0 := []struct {
		in     string
		commas bool
		want   string
	}{
		{"just    a    whole lot of, like,    spaces.   you betcha, ...!",
            false,
			"JUST A WHOLE LOT OF LIKE SPACES YOU BETCHA"},
		{"test, comma, one",
            true,
			"TEST, COMMA ONE"},
		{"test comma two,",
            true,
			"TEST COMMA TWO"},
		{" 010910|1992||||||||||||||||||||||||||||d",
			false,
			"010910 1992 D"},
		{"$aAland, Kurt.",
			true,
			"$AALAND, KURT"},
		{"$aGolinsky, Marie-Franðcoise.",
			true,
			"$AGOLINSKY, MARIE FRANDCOISE"},
		{"$aBrontèe, Emily $d1818-1848.",
			true,
			"$ABRONTEE, EMILY $D1818 1848"},
		{"$aHodges, Margaret, $d1911-2005.",
			true,
			"$AHODGES, MARGARET $D1911 2005"},
		{"$aMcKeown, Adam $q(Adam N.)",
			true,
			"$AMCKEOWN, ADAM $Q ADAM N"},
		{"$aO'Kelley, Mattie Lou.",
			true,
			"$AOKELLEY, MATTIE LOU"},
		{"$aPovey, Karen D., $d1962-.",
			true,
			"$APOVEY, KAREN D $D1962"},
		{"$aRobinson, Sharon, $d1950-.",
			true,
			"$AROBINSON, SHARON $D1950"},
		{"$aShakespeare, William, $d1564-1616.",
			true,
			"$ASHAKESPEARE, WILLIAM $D1564 1616"},
		{"$aKirchengeschichte in Lebensbildern dargestellt, v. 1, Die Frèuhzeit. English. $nv. 1, $pDie Frèuhzeit. $lEnglish.",
			false,
			"$AKIRCHENGESCHICHTE IN LEBENSBILDERN DARGESTELLT V 1 DIE FREUHZEIT ENGLISH $NV 1 $PDIE FREUHZEIT $LENGLISH"},
		{"$aPippi Lçangstrump. English. $lEnglish",
			false,
			"$APIPPI LCANGSTRUMP ENGLISH $LENGLISH"},
		{"$aChristmas in Mexico / $c[Editor, Jadwiga Lâopez ; writer, Corinne Ross].",
			false,
			"$ACHRISTMAS IN MEXICO $CEDITOR JADWIGA LAOPEZ WRITER CORINNE ROSS"},
		{"$aAll-romanized English-Japanese dictionary / $cHyåojun Råomaji Kai.",
			false,
			"$AALL ROMANIZED ENGLISH JAPANESE DICTIONARY $CHYAOJUN RAOMAJI KAI"},
		{"$aWuthering Heights / $cEmily Brontèe.",
			false,
			"$AWUTHERING HEIGHTS $CEMILY BRONTEE"},
		{"$aA world of wheels. $p[Cars of the forties] : $badvertising and euphoria, the Grande Routiáere and the years after the war / $cMichael Sedgwick.",
			false,
			"$AA WORLD OF WHEELS $PCARS OF THE FORTIES $BADVERTISING AND EUPHORIA THE GRANDE ROUTIAERE AND THE YEARS AFTER THE WAR $CMICHAEL SEDGWICK"},
		{"$aThe love of France / $c[by] Marie-Franðcoise Golinsky, Alice Vidal.",
			false,
			"$ATHE LOVE OF FRANCE $CBY MARIE FRANDCOISE GOLINSKY ALICE VIDAL"},
	}
	for _, c := range cases0 {
		rslt := Normalize(c.in, c.commas)
		if rslt != c.want {
			t.Errorf("Normalize(%q) => fail, want (%q)", c.in, c.want)
		}
	}
}
