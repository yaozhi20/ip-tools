package main

import (
	"path/filepath"
	// "encoding/json"
	"fmt"
	"os"	
	// "regexp"
	"runtime"
	"io"
	"bufio"
	"log"
	"strings"
	// "net"
)

func getWinHostsPath() string{

	hostsPath := "/etc/hosts"
	if runtime.GOOS == "windows" {
		hostsPath = getWinSystemDir()
		hostsPath = filepath.Join(hostsPath, "system32", "drivers", "etc", "hosts")
	}

	return hostsPath
}

func getWinSystemDir() string {
	dir := ""
	if runtime.GOOS == "windows" {
		dir = os.Getenv("windir")
	}

	return dir
}

func hostsItems(subject string) {
	
    hostsPath := getWinHostsPath()

	log.Println("hostsPath:", hostsPath)
	log.Println("version:", "0.2")
	// if hostsContent, err := os.ReadFile(hostsPath); err != nil {
	// 	return err
	// } else {
	// 	hostsContent =  string(backup)	
	// }

    file, err := os.Open(hostsPath) // 打开文件
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer file.Close() // 确保文件在函数结束时关闭


	file2, err := os.Create("example.txt") // 创建文件
    if err != nil {
        log.Fatal(err)
    }
    defer file2.Close() // 确保文件在函数结束时关闭
	writer := bufio.NewWriter(file2)
 
    scanner := bufio.NewScanner(file)
    for scanner.Scan() { // 逐行扫描
        
		line := scanner.Text()
		
        if strings.HasPrefix(line, "#") {

			_, err := writer.WriteString(line + "\r\n") // 添加换行符
            if err != nil {
                log.Fatal(err)
            }

            continue
        }else if len(line) == 0{
			_, err := writer.WriteString(line + "\r\n") // 添加换行符
            if err != nil {
                log.Fatal(err)
            }

            continue
		}
		log.Println(line) // 打印行内容
        fields := strings.Fields(line)
        if len(fields) > 0 {
			// ip := fields[0]
            host := fields[1]
            
            // for _, host := range hosts {
            //     if net.ParseIP(ip) != nil {
            //         log.Printf("Host: %s IP: %s\n", host, ip)
            //     }
            // }

			if host != "johnly.xyz" {
				_, err := writer.WriteString(line + "\r\n") // 添加换行符
                if err != nil {
                    log.Fatal(err)
                }
			} else {
				_, err := writer.WriteString(subject + " " + "johnly.xyz\r\n") // 添加换行符
                if err != nil {
                    log.Fatal(err)
                }
			}
        }
    }
 
    if err := scanner.Err(); err != nil {
        log.Println("Scanning error:", err)
    }

	err = writer.Flush() // 刷新缓冲区，确保所有数据都被写入文件
    if err != nil {
        log.Fatal(err)
    }
}

func appendToFile(fileName string, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, io.SeekEnd)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}
	defer f.Close()
	return err
}


func hosts_backup() {
    hostsPath0 := getWinSystemDir()
	hostsPath1 := filepath.Join(hostsPath0, "system32", "drivers", "etc", "hosts")
    now := today()
	fileName := fmt.Sprintf("hosts.%s", now)
	hostsPath2 := filepath.Join(hostsPath0, "system32", "drivers", "etc", fileName)
	args := []string{"/c", "copy", hostsPath1, hostsPath2}
	
	execCmd( args)

}

func hosts_delete() {
    hostsPath0 := getWinSystemDir()
	hostsPath1 := filepath.Join(hostsPath0, "system32", "drivers", "etc", "hosts")
    
	exists, err := fileExists(hostsPath1)

	if err != nil {
		fmt.Println(err)
	}

	if exists {
		args := []string{"/c", "del", hostsPath1}
	
	    execCmd( args)

		
	}

	
}

func hosts_copy() {
    hostsPath0 := getWinSystemDir()
	hostsPath1 := filepath.Join(hostsPath0, "system32", "drivers", "etc")
	args := []string{"/c", "copy", ".\\example.txt", hostsPath1}
	
	execCmd( args)

	hostsPath1 = filepath.Join(hostsPath0, "system32", "drivers", "etc", "example.txt")
	hostsPath2 := filepath.Join(hostsPath0, "system32", "drivers", "etc", "hosts")
	args = []string{"/c", "copy", hostsPath1, hostsPath2}

	execCmd( args)

}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return false, err
}