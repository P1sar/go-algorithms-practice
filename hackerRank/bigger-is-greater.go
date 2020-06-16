package hackerrank

import "fmt"

func biggerIsGreater(w string) string {
	sArr := make([]string, len(w))
	for i, v := range w {
		sArr[i] = string(v)
	}
	sI := findNonIncreasingSuffixStartIndex(sArr)
	if sI == int32(0) {
		return "no answer"
	}
	pivot := sI - 1

	// swap pivot with lesser greater than pivot element(LGTPE) of suffix
	fmt.Println(sArr, pivot, sI)
	LGTPEIndex := sI
	for i := sI; i < int32(len(sArr)); i++ {
		if sArr[i] > sArr[pivot] && sArr[i] <= sArr[LGTPEIndex] {
			LGTPEIndex = i
		}
	}
	fmt.Println(sArr, pivot, sI, LGTPEIndex)


	// swap pivot with LGTPE
	sArr[pivot], sArr[LGTPEIndex] = sArr[LGTPEIndex], sArr[pivot]


	// sort suffix from in asc order
	a := mergeSortString(sArr[sI:])
	s := ""

	for i:=int32(0);i<sI; i++ {
		s += sArr[i]
	}
	for i := 0; i<len(a); i++ {
		s += a[i]
	}

	return s
}


func findNonIncreasingSuffixStartIndex(a []string) int32 {
	for i:=len(a)-1; i>0; i-- {
		if a[i] <= a[i-1] {
			continue
		}
		return int32(i)
	}
	return int32(0)
}


func mergeSortString(arr []string) []string {
	if len(arr) <= 1 {
		return arr
	}
	c := mergeSortString(arr[:len(arr)/2])
	d := mergeSortString(arr[len(arr)/2:])
	res := make([]string,len(c)+len(d))
	k := 0
	j := 0
	for i:=0; i<len(res); i++ {
		if k >= len(d) {
			res[i] = c[j]
			j++
			continue
		}
		if j >= len(c) {
			res[i] = d[k]
			k++
			continue
		}
		if c[j] > d[k] {
			res[i] = d[k]
			k++
		} else {
			res[i] = c[j]
			j++
		}

	}
	return res
}
//imllmmcslslkyoegyoam
//fvincndjrurhf
//rtglgzzqxnuflitnlyti
//mhtvaqofxtyzr
//zalqxykemvzzgaak
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
//xdhqbvlbtmmtsheffj
//kpdpzzohihzwgdgbfz
//kuywptftaaap
//qfqpegznnyludvr
//ufwogufbzaboaepsliqk
//jfejqapjvbdcxtkyr
//sypjbvatgidddoxy
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
//qymbqaehnyzhfqpqplpr
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
//swzsaynxbytytqqt
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
//sajpyegttujxxy
//zmdjphzogfldlkgbchtn
//tbanvjmwixrx
//gmdhdlmopzyvddeyajjq
//yxvmvedubzdcp
//soygdzhbckkfu
//gkbekyrhccw
//wevzqpnqwtpuf
//rbobquotbysufwqjoe
//bpgqfwoyntuhkwov
//schtabphairewhfpm
//rlmrahlisggguykue
//fjtfrmlqvseqk
