package main

import (
	"fmt"
)

func main() {
    
    image := make([][]uint8, 32) // 行数32の行列imageを宣言。
    
    for i := 0; i < 32; i++ {
        image[i] = make([]uint8, 32) // imageの列数を32とする。
        for j :=0; j < 32; j++ {
            image[i][j] = uint8(i*j)
            fmt.Printf("%3d ", image[i][j])
        }
        fmt.Printf("\n")
    }
}