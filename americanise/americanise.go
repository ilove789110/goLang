package americanise

import (
	"bufio" /*提供带缓冲的I/O处理功能*/
	"fmt"
	"io" /*底层io*/
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp" /*正则表达*/
	"stringns"
)

func filenameFromCommandLine() (inFilename, outFilename string, err error) {

	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		err = fmt.Errorf("usage: %s [<]infile.txt [>]outfile.txt", filepath.Base(os.Args[0]))
		return "", "", err
	}

	if len(os.Args) > 1 {
		inFilename = os.Args[1]
		if len(os.Args) > 2 {
			outFilename = os.Args[2]
		}
	}

	if inFilename != "" && inFilename == outFilename {
		log.Fatal("won't overwirte the infile")
	}

	return inFilename, outFilename, nil

}

var britishAmerican = "british-american.txt"

func americanize(inFile io.Reader, outFile io.Writer) (err error) {
	reader := bufio.NewReader(inFile)
	writer := bufio.NewWriter(outFile)

}

func main() {
	inFilename, outFilename, err := filenamesFromCommandLine()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	inFile, outFile := os.Stdin, os.Stdout /*通过os包的两个类型为*os.File的指针类型初始化函数变量*/

	if inFilename != "" {
		if inFile, err = os.Open(inFilename); err != nil {
			log.Fatal(err)
		}

		defer inFile.Close() /*任何属于defer的语句都保证会执行，在defer语句所在函数返回时被调用.如之前走到os.Exit()则不会执行 */
	}

	if outFilename != "" {
		if outFile, err = os.Create(outFilename); err != nil {
			log.Fatal(err)
		}

		defer outFile.Close()
	}

	if err = americanize(inFile, outFile); err != nil {
		log.Fatal(err)
	}

}
