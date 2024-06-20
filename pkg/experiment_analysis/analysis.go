package experiment_analysis

import (
	"bufio"
	"fmt"
	"heterflow/pkg/logger"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type result struct {
	optimizationlevel string
	architecture      string
	name              string
	complier          string
	bitnum            int
}

var (
	total            = 0
	hitNumCount      = [11]int{}
	optlevelhitcount = map[string]int{}
	archhitcount     = map[string]int{}
	selfposhitcount  = map[int]int{}
)

func Readfile(path string) {
	f, err := os.Open(path)
	if err != nil {
		logger.Fatal(f.Name() + "file path is wrong!")
	}
	defer f.Close()
	// 创建Scanner来读取文件
	scanner := bufio.NewScanner(f)
	readline(scanner)
}

func readline(scanner *bufio.Scanner) {
	topk := 5
	// 使用Scan逐行读取
	for scanner.Scan() {
		// 输出当前行
		name, res := splitline(scanner.Text())
		res1 := collectInfo(name)
		count := 0
		for pos := 0; pos < topk; pos++ {
			dir := extract(res[pos])
			res2 := collectInfo(dir)
			if res1 == res2 {
				selfposhitcount[pos]++
			}
			if res1.name == res2.name {
				optlevelhitcount[res2.optimizationlevel]++
				archhitcount[res2.architecture]++
				count++
			}
		}
		hitNumCount[count]++
		total++
	}
	// 检查Scan的错误（除了文件结束之外）
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	printInfo()
}

func collectInfo(input string) result {
	return collectInfoFromDir(input)
}

func splitline(input string) (string, []string) {
	start := strings.IndexByte(input, '{')
	end := strings.LastIndexByte(input, '}')
	str := input[start+1 : end]
	pro := input[:start]
	// 定义正则表达式，匹配一级大括号中的内容
	re := regexp.MustCompile(`{[^{}]*(?:{[^{}]*}[^{}]*)*}`)
	// 查找所有匹配项
	matches := re.FindAllString(str, -1)
	// 打印结果
	// for _, match := range matches {
	// 	fmt.Println(match)
	// }
	return strings.TrimSpace(pro), matches
}

func extract(input string) string {
	tomatch := "} name:"
	start := strings.Index(input, tomatch)
	end := strings.LastIndexByte(input, '}')
	return input[start+len(tomatch) : end]
}

func collectInfoFromDir(input string) result {
	dir, file := filepath.Split(input)
	segments := strings.Split(dir, string(filepath.Separator))
	optimizationlevel := segments[len(segments)-3]
	architecture := segments[len(segments)-4]
	return result{
		optimizationlevel: optimizationlevel,
		architecture:      architecture,
		name:              segments[len(segments)-2] + "_" + file,
	}
}

func collectInfoFromName(input string) result {
	fileName := filepath.Base(input)
	info := strings.Split(fileName, "_")
	app := info[0]
	complier := info[1]
	arch := info[2]
	bitnum, err := strconv.Atoi(info[3])
	if err != nil {
		fmt.Println(err)
	}
	optlevel := info[4]
	fileNameWithoutExt := strings.TrimSuffix(info[5], filepath.Ext(fileName))
	return result{
		optimizationlevel: optlevel,
		architecture:      arch,
		name:              app + "_" + fileNameWithoutExt,
		complier:          complier,
		bitnum:            bitnum,
	}
}

func printInfo() {
	for k, v := range hitNumCount {
		fmt.Printf("The probability of hitting %d times %f \n", k, float64(v)/float64(total))
	}
	fmt.Println()
	for k, v := range optlevelhitcount {
		fmt.Printf("With an optimization level of %s, the probability of hit is %f \n", k, float64(v)/float64(total))
	}
	fmt.Println()
	for k, v := range archhitcount {
		fmt.Printf("Architecture for %s has a %f chance of hitting \n", k, float64(v)/float64(total))
	}
	fmt.Println()
	sum := 0.0
	for k, v := range selfposhitcount {
		fmt.Printf("The probability of pretension at position %d is %f \n", k, float64(v)/float64(total))
		sum += float64(v) / float64(total)
	}
	fmt.Println("The Total probability of pretension is ", sum)
}
