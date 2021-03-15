package main

import (
    "bufio"
    "flag"
    "fmt"
    "net/url"
    "os"
    "strings"
    "sync"
)

var inputFilePath *string
var filterLevel *int
var outputFilePath *string
var extractMode *bool
var wg sync.WaitGroup
var outputMutex sync.Mutex


func check(err error) {
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
}

func writeResultLine(line string) {
    outputMutex.Lock()
    defer outputMutex.Unlock()
    file, _ := os.OpenFile(*outputFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    defer file.Close()
    writer := bufio.NewWriter(file)
    _, _ = fmt.Fprintln(writer, line)
    _ = writer.Flush()
}

func parseUrl(rawUrl string) {
    defer wg.Done()
    if !strings.HasPrefix(rawUrl, "http") {
        rawUrl = "http://" + rawUrl
    }
    url, err := url.Parse(rawUrl)
    if err != nil {
        return
    }

    if strings.Count(url.Host, ".") == *filterLevel- 1 {
        fmt.Println(url.Host)
        if *outputFilePath  != "" {
            writeResultLine(url.Host)
        }
    } else if *extractMode && strings.Count(url.Host, ".") > *filterLevel- 1 {
        dTokens := strings.Split(url.Host, ".")

        subdomain := strings.Join(dTokens[len(dTokens)-*filterLevel:], ".")
        fmt.Println(subdomain)
        if *outputFilePath  != "" {
            writeResultLine(subdomain)
        }
    }
}

func main(){
    inputFilePath = flag.String("i", "","Input file path (required)")
    filterLevel = flag.Int("fl", 3,"Filter domain level, 1 being TLD")
    extractMode = flag.Bool("e", false, "Extract level domains from subdomains")
    outputFilePath = flag.String("o", "","Output file path (default \"\", no output file)")

    flag.Parse()
    if *inputFilePath == "" {
        flag.PrintDefaults()
        fmt.Print("\nError: Input file path [-i] not provided\n")
        os.Exit(1)
    }

    if *outputFilePath != "" {
        _, err := os.Create(*outputFilePath)
        check(err)
    }

    inputFile, err := os.Open(*inputFilePath)
    check(err)
    defer inputFile.Close()
    scanner := bufio.NewScanner(inputFile)
    for scanner.Scan() {
        wg.Add(1)
        parseUrl(scanner.Text())
    }
    wg.Wait()
}