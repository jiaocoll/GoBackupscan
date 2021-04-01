package main

import (
	options2 "GoBackupscan/pkg/options"
	"GoBackupscan/pkg/runner"
	"GoBackupscan/pkg/scan"
	"bufio"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"log"
	"os"
	"sync"
)

func main(){
	var dirs []string

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

	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(runner.Options.Rate, func(i interface{}) {
		if(scan.ScanBacnkup(i.(string))){
			fmt.Println(runner.Options.Url + i.(string))
		}
		wg.Done()
	})
	for _,pathname := range dirs{
		target := runner.Options.Url + pathname
		wg.Add(1)
		_ = p.Invoke(target)
	}
	wg.Wait()
}

