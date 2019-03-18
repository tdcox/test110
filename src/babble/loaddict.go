package babble

import (
	// "io/ioutil"
	// "os"
	"strings"
)

func readAvailableDictionary() (words []string) {
	// file, err := os.Open("./words")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// bytes, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	panic(err)
	// }

	words = strings.Split("win dream evening condition feed tool total basic smell valley nor double seat arrive master track parent shore division sheet substance favor connect post spend chord fat glad original share station dad bread charge proper bar offer segment slave duck instant market degree populate chick dear enemy reply drink occur support speech nature range", " ")
	return
}
