package main

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

	defer func() { /*匿名延迟函数在函数正常返回或者异常退出时都会执行，相比较 defer writer.Flush() 在刷新操作前或者刷新操作中有任何报错都会忽略。所以防御性更强*/
		if err == nil {
			err = writer.Flush()
		}
	}()

	var replacer func(string) string
	if replacer, err = makeReplaceFunc(britishAmerican); err != nil {
		return err
	}

	wordRx := regexp.MustCompile("[A-Za-z]+")

	eof := false

	if !eof {

		var line string
		line, err = reader.ReadString('\n')

		if err == io.EOF {
			err = nil
			eof = true
		} else if err != nil {
			return err
		}

		line = wordRx.ReplaceAllStringFunc(line, replacer)

		if _, err = writer.WriteString(line); err != nil {
			return err
		}

	}

	return nil
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
