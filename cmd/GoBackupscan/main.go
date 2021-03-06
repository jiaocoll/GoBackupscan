package main

import (
	options2 "GoBackupscan/pkg/options"
	"GoBackupscan/pkg/runner"
	"GoBackupscan/pkg/scan"
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/panjf2000/ants/v2"
	"log"
	"os"
	"strings"
	"sync"
)

var(
	dirs []string
	target string
)


func main(){


	options := options2.ParseOptions()
	runner := runner.NewRunner(options)

	dickfile, err := os.OpenFile(runner.Options.Dickname,os.O_RDONLY,1)
	if err != nil {
		log.Println(err)
	}
	tmps := bufio.NewScanner(dickfile)
	for tmps.Scan() {
		tmp := tmps.Text()
		dirs = append(dirs, tmp)

	}
	if string(runner.Options.Url[len(runner.Options.Url)-1]) == "/"{
		runner.Options.Url = runner.Options.Url[:len(runner.Options.Url)-1]
	}
	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(runner.Options.Rate, func(i interface{}) {
		if(scan.ScanBacnkup(i.(string))){
			fmt.Fprintln(color.Output,color.HiCyanString(i.(string)))
		}
		wg.Done()
	})
	for _,pathname := range dirs{
		if strings.Contains(pathname, "/") {
			target = runner.Options.Url + pathname
		}else {
			target = runner.Options.Url + "/" + pathname
		}
		wg.Add(1)
		_ = p.Invoke(target)
	}
	wg.Wait()
}

