/*
	Day 2: Inventory Management System
*/

package main

import (
	"fmt"
	"strings"
)

// User-Specific Input
var input = []string{"wlpiogsvdfecjdqmnxakudrhbz", "wbpioesvyfecjuqmlxaktdrhbz", "blviogavyfecjuqmnxaktdrhbz", "wlpiogsvydecjuqmnipktdrhbz", "wlwiogsvyfmcjuqmoxaktdrhbz", "wlpiogsvphecjuqmnxaktdrzbz", "wlpiogsvyfecjuqmnkakhdrkbz", "wlpiogsvyfhcjuqmnxxktddhbz", "wlpiogsvyfccfuqmnxgktdrhbz", "wlpiogsvhmecjvqmnxaktdrhbz", "wlpiogsvyfecjdqqnxaktdrhyz", "wlpiogyvycecjaqmnxaktdrhbz", "wlpiogsvyfecjfqmnxaktybhbz", "wkpiogsvyfmcauqmnxaktdrhbz", "wlmiolsvyfecjuqmnxaktdrhbn", "wlpioksvyfecjuqmnxaktdrhgs", "wlpiogsvyflcjuvmnxsktdrhbz", "wgziogsvyfecjuqmnxaktdrhoz", "whpingsvyeecjuqmnxaktdrhbz", "wlpiogsvyfecjuqgnxaktdvhlz", "wlpioasvtfecjuqmnxaktdahbz", "wlpihgsvefeceuqmnxaktdrhbz", "wlpiogsvyfecyuqwnxaktdghbz", "wlpfsgsvyfhcjuqmnxaktdrhbz", "wlpiogyvyfecjuqmnxalpdrhbz", "wlpiogsvyfescsqmnxaktdrhbz", "wluiogsvyfecytqmnxaktdrhbz", "wltiorsvyfecjuqmoxaktdrhbz", "wlmiogwvyfejjuqmnxaktdrhbz", "wlpiogsvyfycjuumnxvktdrhbz", "wlkiogsqyfecjqqmnxaktdrhbz", "wlpiogsvyfecouqmnxaktfrubz", "hupiogsvlfecjuqmnxaktdrhbz", "wlpiogsvpxecjuqmnxaksdrhbz", "wlpiogsvyfkcjfqmnxxktdrhbz", "wlpiogsjyfecjuqnnxakthrhbz", "wlpiogsvyfecjuqmnxaktddhdk", "wlpipgsvyfecjuqmnhaktdrubz", "wlpiogsoyfenjpqmnxaktdrhbz", "wlpiogsvyfecnuqmnxaxtdrmbz", "wlpiggsvyfjcjuumnxaktdrhbz", "wlppogsvyfecjuqmnxautdhhbz", "wlpiovbvyfecjuqmnxaktcrhbz", "wlpiogsvyfecjoqmnxaktdrobu", "wlpiohsvyfecjugmnxakthrhbz", "wvpiogovyfecjuqmnxakwdrhbz", "wlpiogsbyfecjuqmnxdktdrtbz", "wlpnogsvyfecjuqmnxakykrhbz", "wlpiogpvyfecjuqmnxvktdrhbs", "wlpiogsvkvecjuqmnxadtdrhbz", "wlkihgsvyfecjuqmnxlktdrhbz", "wlpilgsvyfhcjuqmnxakudrhbz", "wlpioksvysgcjuqmnxaktdrhbz", "wlpiogsvyfecorqmoxaktdrhbz", "wlpiogsvyfectzlmnxaktdrhbz", "wlpiogsvywecjuqwnqaktdrhbz", "wlpiowsvyfecjuqrnxaftdrhbz", "wlpiogsuyfecjutmnxaktnrhbz", "wepiowsvyfqcjuqmnxaktdrhbz", "wlpirgssyfecjuqmvxaktdrhbz", "wlyiogstyfecjuqmnxaktdrhbw", "wlpiogseyfecauqmnxaktdjhbz", "wlpioisvyfenjuqmnxakrdrhbz", "wlpiogsvyfrgjfqmnxaktdrhbz", "wlpionsvyfecjmqmjxaktdrhbz", "alpiggsvyfecjuqmnxaktdrhkz", "wlphngsvyfecjuqmnxfktdrhbz", "wlpiogsvyferjuqmlxakttrhbz", "wlniogsvefecjuqsnxaktdrhbz", "wlpiogsvyfncjurmnxakudrhbz", "wlpiogsvyfecjuqmnxaktprlaz", "wlpiocsvyfecjupmkxaktdrhbz", "wlpihgsvyfecjeqfnxaktdrhbz", "wlwioasvyfjcjuqmnxaktdrhbz", "wlpifgsvyfecjuqsnxaktdshbz", "wlxiogsvyrechuqmnxaktdrhbz", "wlpiogovyfxcjuqmnxakkdrhbz", "wlpiogsvyfecjkqmdxaktmrhbz", "wlpiogsvyfecjwqmntaktdhhbz", "wlpiogsvdfecjuqmmxaktbrhbz", "wlpiogsvyfecauqmnxaksdrhwz", "wlpiogsvwfecjuqznxaktorhbz", "wlpiogtvyfecjuqhnxakidrhbz", "wlpiogsvyyecjuqmnxaktdrhwt", "wljiogsvyfecfuqbnxaktdrhbz", "wlpiogsvybecjuqmnxagtdrjbz", "wrpiogsvyfecjuqmnuaktdrhbd", "wlpiogsvyfecjurmnxnltdrhbz", "blpvogsvyaecjuqmnxaktdrhbz", "bfpiogyvyfecjuqmnxaktdrhbz", "wlpiogsvyfecjuqinxaknddhbz", "wlpizgsvvfecjuqxnxaktdrhbz", "glpiogsvyrecjuqmnxaktdrhbr", "wlpiogskhfecjutmnxaktdrhbz", "wlpiogsvyfecmuqmnxaktdribc", "wlpioesvwfecjuqmnxakkdrhbz", "wlpionsrafecjuqmnxaktdrhbz", "wlsiogsvyfecjuqmnaaktdrhvz", "bloiogsvyfecjuqmnxakjdrhbz", "wlpiogsvyfecjuqenmastdrhbz", "wlpiogyvyfecjuqmuxakldrhbz", "plpiogovyfecjuvmnxaktdrhbz", "wlpiogsvyfebjuqmnkakvdrhbz", "wlziogsvyfhcjuqmngaktdrhbz", "wlsiogsvyfecjuqmnxaktdrzjz", "plbiogsvyfecfuqmnxaktdrhbz", "wfpiogsvyfecjuqknxaktdrhiz", "wlpiogjbyfecjuqmnxaktprhbz", "wmpiogsvyrecjcqmnxaktdrhbz", "wlpiogsyyfecjuqmqxaktdrbbz", "wlpiogsvyfecjuqknxaetdehbz", "wlpiogsvyfezjuqmnxakkdhhbz", "wlpiogsvyfecjjqvnxaktdrhiz", "wkpiogsvyfucjuqmnxaktdrhbd", "lliiogsvyfecjuqmnxaktdrhoz", "wlpiogsvyfecjuqmsxdktdshbz", "wlprogtvyfecjuqmnxaktvrhbz", "wlpizgssyffcjuqmnxaktdrhbz", "wlpioasvyfvcjuqmnxakldrhbz", "wlpoogsvyyecjuqmnxastdrhbz", "wlpiognvyfecjuqmnsaktdrhbr", "wlpiogsoyfecjuqmnxaktdrhho", "wfpiogsvydecjuqmnxaotdrhbz", "wlpiogsvqhecjuqmnxaktdrhhz", "wkpiogsvyfeojuqmnxaktdrqbz", "wlpiogsvyfeveuqmnxaktdshbz", "wlpiogbvyfecjuqmexaktdrcbz", "wlpxogsvyfehjsqmnxaktdrhbz", "wlpcogsvyfecjuqmnjakttrhbz", "wlpiogsvvkecjuqmnxaftdrhbz", "wlpiogsvffecnuqmnxaktdnhbz", "wlpiogsvyfecjupjnxaktdrhbr", "wlpqogsvyfecjuqmnxlktdphbz", "wlpxogsvyfecjvqmnxaktirhbz", "elpiogsvyfecjuqlnxaqtdrhbz", "wspiogsvrfecjuqmnxakadrhbz", "wlpiogsmyfecbuqmnxactdrhbz", "wlpiogsvyfecauqmnyakzdrhbz", "wlsiogsvyfecjuqmnxakvdrnbz", "wlpiogsxyfeijuqmnxakndrhbz", "wlpiogsvyfecjuumnxakbyrhbz", "wlpiogsvyfecjuqmnxhktdyhbo", "wlpiogsvyfecjuqqnxnjtdrhbz", "wapiogsvyfecjuqmnxoktdrmbz", "wlpiogsvyfeejvqmnxaktdrubz", "wlpitgsvyeectuqmnxaktdrhbz", "alpiogsvyfecjulmnxaktdchbz", "wlpiogsvyfecjuqmuxoktdrwbz", "wlpiogsvyfzgjuhmnxaktdrhbz", "wlpnogsvyfecjuqmdxaktyrhbz", "wlpiogsvyfecjuqmnxakthrhra", "wliiogsvyfecluqmnxaktdhhbz", "wlpiogsvyfecjuymnxaltdrhwz", "wlpiogsvyfeljuqmnxaktyrhbd", "wlpiygsvvfecjuqmfxaktdrhbz", "wlpiogihsfecjuqmnxaktdrhbz", "wlpiogjvyfecjuqmnhuktdrhbz", "wldiogsvyfecjiqmwxaktdrhbz", "wlpiogsvjfecjuqmnxaktdrgbr", "wlpioisvyfecjuqwnxabtdrhbz", "wlviogsvyfscjuqmnxqktdrhbz", "wlpiogsvyfecjuqmuxakbdrubz", "wlpiogsvyfecjuqmnxmatdrhqz", "wlpiogsvyfbcjuqwmxaktdrhbz", "wlpiogsvyfexjuqmnxuxtdrhbz", "wljiogsvbfecjuqmnxartdrhbz", "wlpvogsvyfeujuqmnxaktdmhbz", "wnpiogsvyfekjuqanxaktdrhbz", "wlprogsvyfecjuqmzxvktdrhbz", "wkpiogvvyfecjuqmnxaktdrabz", "wlpiogsvwfecjuqmnxaktkbhbz", "wlpiogsvyfecjlqmnxtttdrhbz", "wlpioqsvyfecjuqznxaktyrhbz", "wlpiogsvyfecjuqmnxnethrhbz", "wlpiogsyyfgcjuqmnxaktdrhbm", "wlpiopsvbfecjuqmnxaktdlhbz", "wloqogsvyfucjuqmnxaktdrhbz", "wlpiogsvmfecjuqmnxmktdrhtz", "wlhiogsvyfecjuhmnxaktsrhbz", "wlpioggvpfecjufmnxaktdrhbz", "wlpiogsvyfbcjuomnxaktdrhbh", "wlpmogsvyfecyuqmnxoktdrhbz", "wlpiogslyfecjuqmnxaptyrhbz", "tlpiogsnyfecguqmnxaktdrhbz", "wlpiogsvyfecjuqmnxwktwehbz", "wlpiogsvgfecjuqmnxaktdrabu", "wbpiogsvyfecjuqmnxagtdrhoz", "wlwipgsvyfecjuqmnxaktdrhbu", "wlpwogsvykeczuqmnxaktdrhbz", "wlpgogsvwfecjuqmnxcktdrhbz", "wlpiogsqyfecjuqmrxaktdrrbz", "wlpiogsvyfecjuqmnxakthrafz", "wypicgseyfecjuqmnxaktdrhbz", "wlpiogcvqfecjuqmnxaktdrhzz", "wlriogsvyfecouqmnkaktdrhbz", "wlpiogsvyfemjulmnxaktdrhdz", "flpiogadyfecjuqmnxaktdrhbz", "wupiogsvyfbvjuqmnxaktdrhbz", "wlpiogsvyfebjummnxaktdrrbz", "wjpiogsvyfecjuqmnxaktprybz", "wlpirgsvyfecjiqmnxaatdrhbz", "bvpiogsvyfecjuqmnxaktdrhez", "wlpiogsvyfxcjuqmnxykzdrhbz", "wlkiwgsqyfecjqqmnxaktdrhbz", "wepaogsvyfecjxqmnxaktdrhbz", "wlpiovsvyfecjjqmnxaktdmhbz", "wlpioysryfecjuqmnxaktdrhiz", "wlpizjsvyfecjuvmnxaktdrhbz", "dlpiogsvyfecjucmnxakbdrhbz", "wlpiogsccfecjrqmnxaktdrhbz", "wlpioggvyfecpuqmnxagtdrhbz", "wlpiogsvyfvcjuumlxaktdrhbz", "wwpiogsryfjcjuqmnxaktdrhbz", "wlpiogsvyfecjuqynxaktdrogz", "wlpiogsvyfecjpqmnxskbdrhbz", "wlpiogsvyfecjuhmfxaktvrhbz", "wlpiogevyfecjrqmnwaktdrhbz", "wlpiigsvyfemjuqmnxaktdrhtz", "wlpcogsvyfecjuqhnxakgdrhbz", "wupiogsvyfxcjuqmnxaktdrhgz", "wlsiogsvyfecjuqenxuktdrhbz", "wlpioglvyfecjujmexaktdrhbz", "wlriogsvyfeljuqmnxattdrhbz", "wlpiogsvyfecfuqmhxaktkrhbz", "wlppogsvyfecjuqmxxabtdrhbz", "wlniogsvyfevjuqwnxaktdrhbz", "wlhiogsvyfecjuqmnxactxrhbz", "ilpiogivyflcjuqmnxaktdrhbz", "wlpmogsvyfecjuqmnxaktdrlbs", "wipiogsvyfeqjuqmnxaktrrhbz", "wvpiogsvyfecjuqknxaktdrrbz", "wwpioguvyfecxuqmnxaktdrhbz", "wlpiogsvkfdcjuqmnxaktdzhbz", "wlpiogfvyfecjuqmnxadtdrhbg", "wlpiogsvyzefjuqfnxaktdrhbz", "wlpiogstyfechuqmnxaktdchbz", "wlpiogszyfedjuqmnxsktdrhbz", "wzpiozsvyfncjuqmnxaktdrhbz", "xlpiogsvyfefjuqmnmaktdrhbz", "wlpiogsvyfebxummnxaktdrhbz", "wlpiogsgyfecfurmnxaktdrhbz", "wlpqogsvyfecjuomnxaktdrhbi", "wlpiogjvufecjuqmnxaktdrhbd", "wlpiolsvyfecduqmnxaktrrhbz", "wlpxogsvyfecjuqmnxaktgrhbk", "wlpiogsfyfncjuqmnxsktdrhbz", "wlpioggvyfecjufmnxaktdrebz", "wlpiogsvyfecfujmnxaktdrwbz", "rlpiogsvyfecjlqmnxaktdqhbz", "wlpfogsvyfecjuimnxaktfrhbz"}

// lowerAlpha generates a lowercase alphabet
func lowerAlpha() string {
	p := make([]byte, 26)
	for i := range p {
		p[i] = 'a' + byte(i)
	}
	return string(p)
}

// getChecksum finds the number of strings which contain a rune exactly twice,
// and the number which contain a rune exactly thrice, then multiplies these
// values to produce a checksum
func getChecksum(input []string) int {
	var twice map[string]bool
	var thrice map[string]bool

	twice = make(map[string]bool)
	thrice = make(map[string]bool)

	alpha := lowerAlpha()

	for _, code := range input {
		for _, letter := range alpha {
			num := strings.Count(code, string(letter))
			if num == 2 {
				twice[code] = true
			} else if num == 3 {
				thrice[code] = true
			}
		}
	}

	return len(twice) * len(thrice)
}

// getPrototypes finds the two input strings which differ by only one rune
func getPrototypes(input []string) (pt []string) {
OUTER:
	for _, code := range input {
		for _, code2 := range input {
			var different int

			for i := 0; i < len(code); i++ {
				if code[i] != code2[i] {
					different++
				}
			}

			if different == 1 {
				pt = []string{code, code2}
				break OUTER
			}
		}
	}

	return
}

// getSameLetters finds the identical letters of two strings
//
// NOTE: We consider "identical" to include not only the rune but the string index
func getSameLetters(one string, two string) (same string) {
	for i := 0; i < len(one); i++ {
		if one[i] == two[i] {
			same += string(one[i])
		}
	}

	return
}

// main
func main() {
	cs := getChecksum(input)

	fmt.Printf("\n\tChecksum: %d\n", cs)

	pt := getPrototypes(input)

	if len(pt) != 2 {
		panic("Found more than two prototypes - cannot recover.")
	}

	fmt.Printf("\n\tPrototypes: %s, %s\n", pt[0], pt[1])

	same := getSameLetters(pt[0], pt[1])

	fmt.Printf("\n\tSimilar Letters: %s\n", same)
}
