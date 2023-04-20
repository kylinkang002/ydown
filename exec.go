package main

// 本程序用于演示如何使用exec包来执行外部命令
// 可以在多台机器上执行命令
// 可以串行执行多个命令，也可以并行执行多个命令
// 可以通过管道将命令的输出作为另一个命令的输入


import (
	"fmt"
	"os/exec"
)


// golang 调用python3 执行python脚本，自动定位python3的路径
func python3_path() (string, error) {
	cmd := exec.Command("which", "python3")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	return string(out), err
}

func pytube_path() (string, error) {
	cmd := exec.Command("which", "pytube")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	return string(out), err
}

// download youtube video with pytube
func youtube_dl(python3 string, pytube string, url string) (output string, err error) {
	cmd := exec.Command(python3, pytube, url)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	return string(out), err
}

func main() {
	python3, err := python3_path()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(python3)
	pytube, err := pytube_path()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(pytube)

	out, err := youtube_dl(python3, pytube, "https://www.youtube.com/watch?v=QH2-TGUlwu4")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}


