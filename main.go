package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "strconv"
    "strings"
)

func main() {
    f, err := os.Open("input.txt")
    if err != nil {
        log.Println(err)
        return
    }
    defer f.Close()

    // Read the file content
    fc, err := ioutil.ReadAll(f)
    if err != nil {
        log.Println(err)
        return
    }

    numStrs := strings.Split(string(fc), "\n")
    nums := make([]int, len(numStrs))

    for idx, numStr := range numStrs {
        nums[idx], _ = strconv.Atoi(numStr)
    }

    // Let's find the digits that sum to 2020 now
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            for k := j + 1; k < len(nums); k++ {
                n1, n2, n3 := nums[i], nums[j], nums[k]
                sum := n1 + n2 + n3

                if sum == 2020 {
                    fmt.Println(n1 * n2 * n3)
                    return
                }
            }
        }
    }
}
