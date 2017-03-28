///Counting WORDS in files 


func countFile_chan(fname string, c chan map[string]int) {
    count := map[string]int{}
    f, err := os.Open(fname)
    if err == nil {
        scanner := bufio.NewScanner(f)
        scanner.Split(bufio.ScanWords)
        for scanner.Scan() {
            word := scanner.Text()
            count[word]++
        }
    } else {
        fmt.Println("countFile: error reading", fname)
    }
    f.Close()   // important: close the file
    c <- count
}


// COUNTING ALL THE FILES 

func countAllFiles() {
    // countChan is a buffered channel to avoid "too many files open" error
    var countChan = make(chan map[string]int, 200)

    count := map[string]int{}

    fileNames := getPwdFiles()

    fmt.Printf("%d files in pwd ...\n", len(fileNames))

    // launch all the goroutines
    for _, fname := range fileNames {
        go countFile_chan(fname, countChan)
    }

    // drain the channel
    for i := 0; i < len(fileNames); i++ {
        m := <-countChan  // wait for a result on the channel
        for w, c := range m {
            count[w] += c
        }
    }

    fmt.Println(len(count), " different words")

    wc := 0
    for _, val := range count {
        wc += val
    }

    fmt.Println(wc, " total words")
}

