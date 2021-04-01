package options

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
)

type Options struct {
	Url string
	Dickname string
	Rate int
}

func ParseOptions()*Options{
	options := &Options{}
	flag.StringVar(&options.Dickname,"dict","","字典文件")
	flag.StringVar(&options.Url,"u","","要扫描的目标域名,例如:http://www.example.com")
	flag.IntVar(&options.Rate,"rate",2000,"扫描速率")
	flag.Usage = usage
	flag.Parse()

	return options
}

func usage(){
	fmt.Fprintf(color.Output,color.CyanString(`Go语言备份文件扫描工具
Options:
`))
	flag.PrintDefaults()
}