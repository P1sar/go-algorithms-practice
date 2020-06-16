package hackerrank

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestNonDivisibleSubSet(t *testing.T) {
	res := nonDivisibleSubset(7, []int32{278,576,496,727,410,124,338,149,209,702,282,718,771,575,436})
	if res != 11 {
		fmt.Println(res)
		t.Fail()
	}
}

func TestNonDS(t *testing.T) {
	res := nonDivisibleSubset(3, []int32{1,7,2,4})
	if res != 3 {
		fmt.Println(res)
		t.Fail()
	}
}


func TestNonDS3(t *testing.T) {
	res := nonDivisibleSubset(5, []int32{770528134,663501748,384261537,800309024,103668401,538539662,385488901,101262949,557792122,46058493})
	if res != 6 {
		fmt.Println(res)
		t.Fail()
	}
}

func TestNonDS4(t *testing.T) {
	res := nonDivisibleSubset(4, []int32{1,2,3,4,5,6,7,8,9,10})
	if res != 5 {
		fmt.Println(res)
		t.Fail()
	}
}

func TestRepeatedStringA(t *testing.T) {
	if res := repeatedString("aba", 10); res != 7 {
		fmt.Println(res)
		t.Fail()
	}
}

func TestJumpingOnClouds(t *testing.T) {
	if re := jumpingOnClouds([]int32{0,0,0,1,0,0}); re != 3 {
		fmt.Println(re)
		t.Fail()
	}
}

func TestEqualizeArray(t *testing.T) {
	if re := equalizeArray([]int32{3,3,2,1,3}); re != 2 {
		fmt.Println(re)
		t.Fail()
	}
}

func TestQueensAttack(t *testing.T) {
	if re := queensAttack(5, 3, 4, 3, [][]int32{{3, 2}, {4, 2}, {2, 3}}); re != 8 {
		fmt.Println(re)
		t.Fail()
	}
}
func TestQueensAttackII(t *testing.T) {
	if re := queensAttack(5, 3, 4, 3, [][]int32{{3, 4}, {4, 2}, {2, 3}, {2,5}}); re != 8 {
		fmt.Println(re)
		t.Fail()
	}
}

func TestQueensAttackIII(t *testing.T) {
	if re := queensTestFromFile("../test-data/n_queens_62.txt"); re != 62 {
		fmt.Println(re)
		t.Fail()
	}
}


func queensTestFromFile(fp string) int32 {
	file, err := os.Open(fp)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, 1024 * 1024)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	k := int32(kTemp)

	r_qC_q := strings.Split(readLine(reader), " ")

	r_qTemp, err := strconv.ParseInt(r_qC_q[0], 10, 64)
	r_q := int32(r_qTemp)

	c_qTemp, err := strconv.ParseInt(r_qC_q[1], 10, 64)
	c_q := int32(c_qTemp)

	var obstacles [][]int32
	for i := 0; i < int(k); i++ {
		obstaclesRowTemp := strings.Split(readLine(reader), " ")

		var obstaclesRow []int32
		for _, obstaclesRowItem := range obstaclesRowTemp {
			obstaclesItemTemp, _ := strconv.ParseInt(obstaclesRowItem, 10, 64)
			obstaclesItem := int32(obstaclesItemTemp)
			obstaclesRow = append(obstaclesRow, obstaclesItem)
		}

		if len(obstaclesRow) != int(2) {
			panic("Bad input")
		}

		obstacles = append(obstacles, obstaclesRow)
	}
	if err != nil {
		panic(err)
	}

	result := queensAttack(n, k, r_q, c_q, obstacles)

	return result

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}


func TestAcmTeam(t *testing.T) {
	if re := acmTeam([]string{"10101", "11100", "11010", "00101"}); re[0] != 5 || re[1] != 2 {
		fmt.Println(re)
		t.Fail()
	}
}

func TestAcmTeamFromFile(t *testing.T) {
	file, err := os.Open("../test-data/acm_team.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReaderSize(file, 1024 * 1024)
	var topic = make([]string, 500)

	for i := 0; i < int(500); i++ {
		topicItem := readLine(reader)
		topic[i] = topicItem
	}
	if re := acmTeam(topic); re[0] != 416 || re[1] != 1 {
		fmt.Println(re)
		t.Fail()
	}

}

func TestOrganizingContainers(t *testing.T) {
	if re := organizingContainers([][]int32{{0,2}, {1,1}}); re != "Impossible" {
		t.Fail()
	}
	if re := organizingContainers([][]int32{{1,3,1}, {2,1,2}, {3,3,3}}); re != "Impossible" {
		t.Fail()
	}
	if re := organizingContainers([][]int32{{0,2,1}, {1,1,1}, {2,0,0}}); re != "Possible" {
		t.Fail()
	}
}

func TestEncryption(t *testing.T) {
	if re := encryption("haveaniceday"); re != "hae and via ecy" {
		fmt.Println(re)
		t.Fail()
	}
	if re := encryption("feedthedog"); re != "fto ehg ee dd" {
		fmt.Println(re)
		t.Fail()
	}
	if re := encryption("chillout"); re != "clu hlt io" {
		fmt.Println(re)
		t.Fail()
	}
	if re := encryption("wclwfoznbmyycxvaxagjhtexdkwjqhlojykopldsxesbbnezqmixfpujbssrbfhlgubvfhpfliimvmnny"); re != "wmgjpnull cyjqlejgi lyhhdzbui wctlsqsbm fxeoxmsvv ovxjeirfm zadysxbhn nxkkbffpn bawobphfy" {
		fmt.Println(re)
		t.Fail()
	}
}

func TestBiggerIsGreater(t *testing.T) {
	if re := biggerIsGreater("zalqxykemvzzgaka"); re != "zalqxykemvzzgkaa" {
		fmt.Println(re)
		t.Fail()
	}
	if re := biggerIsGreater("bb"); re != "no answer" {
		fmt.Println(re)
		t.Fail()
	}
}

func TestFindNonIncreasingSuffixIndex(t *testing.T) {
	if re := findNonIncreasingSuffixStartIndex([]string{"0", "1", "2", "5", "3", "3", "0"}); re != 3 {
		fmt.Println(re)
		t.Fail()
	}
	if re := findNonIncreasingSuffixStartIndex([]string{"0", "1", "2", "3", "4", "5", "6"}); re != 6 {
		fmt.Println(re)
		t.Fail()
	}
	if re := findNonIncreasingSuffixStartIndex([]string{"6", "5", "4", "3", "2", "1", "0"}); re != 0 {
		fmt.Println(re)
		t.Fail()
	}
}
//
//imllmmcslslkyoegyoam
//fvincndjrurhf
//rtglgzzqxnuflitnlyti
//mhtvaqofxtyzr
//zalqxykemvzzgaka
//wjjulziszbqqdcpdnhod
//japjbvjlxzkgietmk
//jqczvgqywydkunmwj
//ehdegnmorgafrjxvsck
//tydwixlwghlomq
//wddnwjneaxbwhwarm
//pnimbesirfbixlv
//mijamkzpiiniveki
//qxtwpdpwexuje
//qtcshorwykc
//xoojiggdcyjrurp
//vcjmvngcdyabcmzj
//xildrrpach
//rrcntnbqchsfhvjhi
//kmotatmrabtcoum
//bnfcejmyotwv
//dnppdkpywiaxddoxq
//tmowsxkrodmkrak
//jfkaehlegowfggh
//ttylsiegnttymxty
//kyetllczuyibdkwyiqhr
//xdhqbvlbtmmtshejff
//kpdpzzohihzwgdgbfz
//kuywptftpaaa
//qfqpegznnyludvr
//ufwogufbzaboaepsliqk
//jfejqapjvbdcxtkyr
//sypjbvatgiodddxy
//wdpfyqjcpnc
//baabpjckkyturd
//uvwurzjyzbhcqmryprqa
//kvtwtmqygksbmi
//ivsjycnooeodwtp
//zqyxjnnitzawipsmq
//blmrzavodtfzyezp
//bmqlhqndavc
//phvauobwkrcfwdedcs
//vpygyqubqywkndhwpz
//yikanhdrwjx
//vnpblfxmvwkflqokbr
//pserilwzxwyorldsxlks
//qymbqaehnyzhfqpqrlpp
//fcakwzuqlzglnidbkmq
//jkscckttaeifiksgkxmx
//dkbllravwnhhfjjrec
//imzsyrykfvtj
//tvogoocldlukwfcajvxi
//cvnagtypozljprajglv
//hwcmacxvsmu
//rhrzcpprqcfc
//clppxvwtaktchqrfdi
//qwusnlldnolqh
//yitverajov
//arciyxaxtvmfgqwbu
//pzbxvxdjuuvvu
//nxfowilpdxwltp
//swzsaynxbytyttqq
//qyrogefletey
//iotjgthvslvmjpcchufh
//knfpyjtzqf
//tmtbfayantwkm
//asxwzygnngw
//rmwiwrurutb
//bhmpfwhgqfcqfldlsh
//yhbidtewppg
//jwwbeuiklpodziiv
//anjhprmkwieb
//lpwhqaebrm
//dunecynelymcpyonqj
//hblfldireuivzekegti
//uryygzpwifrriecgvw
//kzuhaysegaxtwqtxv
//kvarmrbpoxxujhvgwp
//hanhaggqzdpunkugzmqh
//gnwqwsylqeurq
//qzkjbnyvclrkmtcd
//argsnaqbqvu
//obbnlkoaklxc
//ojiilqieycsasvqosyuc
//qhlgiwsmtxbffjtsx
//vvrvnmndeopgy
//ibeqzyeuvzbf
//sajpyegttujyxx
//zmdjphzogfldlkgbchtn
//tbanvjmwixrx
//gmdhdlmopzyvddeyajjq
//yxvmvedubzdcp
//soygdzhbckkfu
//gkbekyrhwcc
//wevzqpnqwtpuf
//rbobquotbysufwqjoe
//bpgqfwoyntuhkwov
//schtabphairewhfpm
//rlmrahlisggguykue
//fjtfrmlqvseqk