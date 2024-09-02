package main

import (
	// "path/filepath"
	// "encoding/json"
	// "fmt"
	// "os"	
	// "regexp"
	// "runtime"
	// "io"
	// "bufio"
	"log"
	// "strings"
	// "net"
	"bytes"
    "os/exec"
)


func execCmd( args []string) {
    // 设置命令和参数
    cmd := exec.Command("cmd",  args...)
 
    // 创建缓冲区用于保存命令的输出
    var out bytes.Buffer
    cmd.Stdout = &out
 
    // 运行命令
    err := cmd.Run()
    if err != nil {
        log.Println("命令执行出错:", err)
        return
    }
 
    // 打印输出结果
    log.Println("命令输出:")
    log.Println(out.String())
}