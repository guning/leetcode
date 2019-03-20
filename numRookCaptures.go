package main

import "fmt"

func numRookCaptures(board [][]byte) int {
    x, y := findR(board)
    var num int
    //left
    for i:=y-1; i>=0; i-- {
        if board[x][i] == 'B' {
            break
        }
        if board[x][i] == 'p' {
            num++
            break
        }
    }
    //right
    for i:=y+1; i<len(board[x]); i++ {
        if board[x][i] == 'B' {
            break
        }
        if board[x][i] == 'p' {
            num++
            break
        }
    }

    //up
    for j:=x-1; j>=0; j-- {
        if board[j][y] == 'B' {
            break
        }
        if board[j][y] == 'p' {
            num++
            break
        }
    }
    //up
    for j:=x+1; j<len(board); j++ {
        if board[j][y] == 'B' {
            break
        }
        if board[j][y] == 'p' {
            num++
            break
        }
    }
   
    return num
}


func findR(board [][]byte) (int,int) {
    var x,y int
    for i,b := range board {
        for j,v := range b {
            if v == 'R' {
                x = i
                y = j
                break
            }
        }
    }
    return x,y
}


func main() {
//    board := [][]byte{{'.','.','.','.','.','.','.','.'},{'.','.','.','p','.','.','.','.'},{'.','.','.','R','.','.','.','p'},{'.','.','.','.','.','.','.','.'},{'.','.','.','.','.','.','.','.'},{'.','.','.','p','.','.','.','.'},{'.','.','.','.','.','.','.','.'},{'.','.','.','.','.','.','.','.'}}
    board := [][]byte{{'.','.','.','.','.','.','.','.'},{'.','.','B','B','B','B','B','.'},{'.','p','B','p','p','p','B','p'},{'.','p','B','p','R','p','B','p'},{'.','p','B','p','p','p','B','p'},{'.','.','B','B','B','B','B','.'},{'.','.','.','p','p','p','.','.'},{'.','.','.','.','.','.','.','.'}}
    fmt.Println(numRookCaptures(board))
}
