package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/1005281342/godatastructures/disjointsets"
)

const (
	Find  = "Q"
	Union = "M"
)

const (
	Yes = "Yes"
	No  = "No"
)

func main() {
	var (
		reader  = bufio.NewReader(os.Stdin)
		arr     = readIntArray(reader) // len(arr) == 2
		dsu     disjointsets.DSU
		strList []string
	)
	dsu = disjointsets.ConstructorDSU(arr[0])
	for i := 0; i < arr[1]; i++ {
		strList = readStrArray(reader)
		switch strList[0] {
		case Find:
			if dsu.Find(str2Int(strList[1])-1) == dsu.Find(str2Int(strList[2])-1) {
				fmt.Println(Yes)
			} else {
				fmt.Println(No)
			}
		case Union:
			dsu.Union(str2Int(strList[1])-1, str2Int(strList[2])-1)
		default:
			fmt.Println(No)
		}
	}
}

func str2Int(s string) int {
	var a, _ = strconv.Atoi(s)
	return a
}

func readLine(reader *bufio.Reader) string {
	var line, _ = reader.ReadString('\n')
	return strings.TrimRight(line, "\n")
}

func readInt(reader *bufio.Reader) int {
	var num, _ = strconv.Atoi(readLine(reader))
	return num
}

func readStrArray(reader *bufio.Reader) []string {
	var line = readLine(reader)
	return strings.Split(line, " ")
}

func readIntArray(reader *bufio.Reader) []int {
	var line = readLine(reader)
	var strList = strings.Split(line, " ")
	var nums = make([]int, 0)
	var err error
	var v int
	for i := 0; i < len(strList); i++ {
		if v, err = strconv.Atoi(strList[i]); err != nil {
			continue
		}
		nums = append(nums, v)
	}
	return nums
}

func nums2string(x []int, sep string) string {
	var b strings.Builder
	for i := 0; i < len(x); i++ {
		b.WriteString(strconv.Itoa(x[i]))
		b.WriteString(sep)
	}
	return b.String()
}

// https://www.acwing.com/problem/content/838/
