package main
//– flag库   命令行参数解析  在当前目录下执行命令：go run flag_.go -b back -d true
import (
	"flag"
	"fmt"
)

func main() {
	backup_dir:=flag.String("b","","backup path")
	debug_mode:=flag.Bool("d",false,"debug mode")
	flag.Parse()
	fmt.Println("backup_dir:",*backup_dir)
	fmt.Println("debug_mode:",*debug_mode)


}
