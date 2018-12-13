package main

// PBKDF2 SHA256 Cracker

import (
    "fmt"
    "runtime"
    "os"
    "bufio"
    "io"
    "strings"
    "sync"
    "utils"
    // "github.com/beego/wetalk/modules/utils"
)

func main() {
    numcpu := runtime.NumCPU()
    runtime.GOMAXPROCS(numcpu)

    var wg sync.WaitGroup

    // read password file
    pwd_filename := "../password.txt"
    file2, err2 := os.OpenFile(pwd_filename, os.O_RDWR, 0666)
    if err2 != nil {
        fmt.Println("Open file error!", err2)
        return
    }
    defer file2.Close()
    pwd_buf := bufio.NewReader(file2)

    for {
        // read password line
        pwdline, pwderr := pwd_buf.ReadString('\n')
        pwdline = strings.TrimSpace(pwdline)
        // fmt.Println("Cracking Password: ", pwdline)

        // read hash file
        hash_filename := "../user.csv"
        file, err := os.OpenFile(hash_filename, os.O_RDWR, 0666)
        if err != nil {
            fmt.Println("Open file error!", err)
            return
        }
        defer file.Close()
        hash_buf := bufio.NewReader(file)

        for {
            // read hash line
            line, err := hash_buf.ReadString('\n')
            line = strings.TrimSpace(line)
            cols := strings.Split(line, ",")

            if len(cols)>2 {
                name := cols[0]
                hash := cols[1]
                salt := cols[2]

                // fmt.Println("Cracking "+name)
                if hash!="" && salt!="" {
                    wg.Add(1)   // add flag to wait subprocess
                    go func(pwdline string, hash string, salt string) {
                        pwd := utils.EncodePassword(pwdline, salt);
                        if hash == pwd {
                            // fmt.Println("Cracked: " + pwdline) 
                            fmt.Println("[Cracked] ",name,pwdline,pwd,salt); 
                        }
                        defer wg.Done()
                    }(pwdline, hash, salt)

                }
            }
            if err != nil {
                if err == io.EOF {
                    // fmt.Println("hashfile read ok!")
                    break
                } else {
                    fmt.Println("Read file error!", err)
                    return
                }
            }
        }
        wg.Wait()

        if pwderr != nil {
            if pwderr == io.EOF {
                fmt.Println("pwdfile read ok!")
                break
            } else {
                fmt.Println("Read file error!", pwderr)
                break
            }
        }
        
    }
}