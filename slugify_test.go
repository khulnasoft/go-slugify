package slugify

import (
	"log"
	"testing"
)

type testCase struct {
	input, expect string
}

func testSlugify(t *testing.T, slugifier Slugifier, input, expect string) {
	log.Printf("input %v, expect %v", input, expect)
	ret := slugifier.Slugify(input)
	if ret != expect {
		t.Errorf("Expected %v, got %v", expect, ret)
	}
}

func testSlugifyCases(t *testing.T, slugifier Slugifier, cases []testCase) {
	for _, c := range cases {
		testSlugify(t, slugifier, c.input, c.expect)
	}
}

func TestVersion(t *testing.T) {
	ret := Version()
	expect := "0.2.0"
	if ret != expect {
		t.Errorf("Expected %v, got %v", expect, ret)
	}
}

func TestSlugifyDefaults(t *testing.T) {
	cases := []testCase{
		{"", ""},
		{"abc", "abc"},
		{"abc234", "abc234"},
		{"This is a test ---", "this-is-a-test"},
		{"___This is a test___", "this-is-a-test"},
		{"This -- is a ## test ---", "this-is-a-test"},
		{"北京kožušček", "bei-jing-kozuscek"},
		{"Nín hǎo. Wǒ shì zhōng guó rén", "nin-hao-wo-shi-zhong-guo-ren"},
		{`C\'est déjà l\'été.`, "c-est-deja-l-ete"},
	}

	slugifier := NewSlugifier()
	testSlugifyCases(t, slugifier, cases)
}

func TestSlugifyToLower(t *testing.T) {
	cases := []testCase{
		{"", ""},
		{"abc", "abc"},
		{"abc234", "abc234"},
		{"This is a test ---", "This-is-a-test"},
		{"___This is a test___", "This-is-a-test"},
		{"This -- is a ## test ---", "This-is-a-test"},
		{"北京kožušček", "Bei-Jing-kozuscek"},
		{"Nín hǎo. Wǒ shì zhōng guó rén", "Nin-hao-Wo-shi-zhong-guo-ren"},
		{`C\'est déjà l\'été.`, "C-est-deja-l-ete"},
	}

	slugifier := NewSlugifier()
	slugifier.ToLower(false)
	testSlugifyCases(t, slugifier, cases)
}

func TestSlugifyWordSeparator(t *testing.T) {
	cases := []testCase{
		{"", ""},
		{"abc", "abc"},
		{"abc234", "abc234"},
		{"This is a test ---", "this_is_a_test"}, // trailing separator+invalid removed
		{"_-_This is a test", "this_is_a_test"},  // leading separator+invalid removed
		{"This -- is \t\t  \r\n a ## test ---", "this_--_is_a_--_test"}, // successive whitespace is reduced to 1
		{"北京kožušček", "bei_jing_kozuscek"},
		{"Nín hǎo. Wǒ shì zhōng guó rén", "nin_hao-_wo_shi_zhong_guo_ren"},
		{`C\'est déjà l\'été.`, "c--est_deja_l--ete"},
	}

	slugifier := NewSlugifier()
	slugifier.WordSeparator("_")
	testSlugifyCases(t, slugifier, cases)
}

func TestSlugifyEmptyWordSeparator(t *testing.T) {
	cases := []testCase{
		{"", ""},
		{"abc", "abc"},
		{"abc234", "abc234"},
		{"This is a test ---", "thisisatest"}, // trailing separator+invalid removed
		{"_-_This is a test", "thisisatest"},  // leading separator+invalid removed
		{"This -- is \t\t  \r\n a ## test ---", "this--isa--test"}, // successive whitespace is reduced to 1
		{"北京kožušček", "beijingkozuscek"},
		{"Nín hǎo. Wǒ shì zhōng guó rén", "ninhao-woshizhongguoren"},
		{`C\'est déjà l\'été.`, "c--estdejal--ete"},
	}

	slugifier := NewSlugifier()
	slugifier.WordSeparator("")
	testSlugifyCases(t, slugifier, cases)
}

func TestSlugifyInvalidCharReplacement(t *testing.T) {
	cases := []testCase{
		{"", ""},
		{"abc", "abc"},
		{"abc234", "abc234"},
		{"This is a test ---", "this-is-a-test"}, // trailing separator+invalid removed
		{"_-_This is a test", "this-is-a-test"},  // leading separator+invalid removed
		{"This -- is \t\t  \r\n a ## test ---", "this-is-a-__-test"}, // successive whitespace is reduced to 1
		{"北京kožušček", "bei-jing-kozuscek"},
		{"Nín hǎo. Wǒ shì zhōng guó rén", "nin-hao_-wo-shi-zhong-guo-ren"},
		{`C\'est déjà l\'été.`, "c__est-deja-l__ete"},
	}

	slugifier := NewSlugifier()
	slugifier.InvalidChar("_")
	testSlugifyCases(t, slugifier, cases)
}

func TestSlugifyEmptyInvalidCharReplacement(t *testing.T) {
	cases := []testCase{
		{"", ""},
		{"abc", "abc"},
		{"abc234", "abc234"},
		{"This is a test ---", "this-is-a-test"}, // trailing separator+invalid removed
		{"_-_This is a test", "this-is-a-test"},  // leading separator+invalid removed
		{"This -- is \t\t  \r\n a ## test ---", "this-is-a-test"}, // successive whitespace is reduced to 1
		{"北京kožušček", "bei-jing-kozuscek"},
		{"Nín hǎo. Wǒ shì zhōng guó rén", "nin-hao-wo-shi-zhong-guo-ren"},
		{`C\'est déjà l\'été.`, "cest-deja-lete"},
	}

	slugifier := NewSlugifier()
	slugifier.InvalidChar("")
	testSlugifyCases(t, slugifier, cases)
}

func TestSlugifyEmptyWordSeparatorAndInvalidCharReplacement(t *testing.T) {
	cases := []testCase{
		{"", ""},
		{"abc", "abc"},
		{"abc234", "abc234"},
		{"This is a test ---", "thisisatest"},
		{"_-_This is a test", "thisisatest"},
		{"This -- is \t\t  \r\n a ## test ---", "thisisatest"},
		{"北京kožušček", "beijingkozuscek"},
		{"Nín hǎo. Wǒ shì zhōng guó rén", "ninhaowoshizhongguoren"},
		{`C\'est déjà l\'été.`, "cestdejalete"},
	}

	slugifier := NewSlugifier()
	slugifier.InvalidChar("")
	slugifier.WordSeparator("")
	testSlugifyCases(t, slugifier, cases)
}

func TestSlugifyReplacementsBecomeValidCharacters(t *testing.T) {
	cases := []testCase{
		{"**##x**##**x##**", "x*##*x"},
		{"##**x##**##x**##", "x##*##x"},
	}

	slugifier := NewSlugifier()
	slugifier.InvalidChar("#")
	slugifier.WordSeparator("*") // valid; but still duplicates removed!
	testSlugifyCases(t, slugifier, cases)
}

func TestSlugifyChainingSetup(t *testing.T) {
	cases := []testCase{
		{"This -- is \t\t  \r\n a ## test ---", "This*##*is*a*##*test"},
	}

	slugifier := (&Slugifier{}).ToLower(false).InvalidChar("#").WordSeparator("*")

	slugifier.InvalidChar("#")
	slugifier.WordSeparator("*")
	testSlugifyCases(t, *slugifier, cases)
}
