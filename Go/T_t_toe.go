package main

import (
    "fmt"         // for the imput and output functions.
    "math/rand"  // for generating random number.
    "strings"    // for string manipulation.
    "time"       //for seeding the random number generator.
)

var cells [25]string
// Define constants for players.
const (
    PlayerMG = "MG" // constant for human player which is me MG is my initials. 
    PlayerAI = "AI"  // constant the AI player as a computer.
)
// Initialize the board with empty values.
func InitializeBoard() {
    for i := range cells {
        cells[i] = "--" // set each cell to an initial balue of "--".
    }
}

// Display the board.
func DisplayBoard() {
    for i := 0; i < 25; i += 5 {
        fmt.Println(cells[i], cells[i+1], cells[i+2], cells[i+3], cells[i+4])
    }
}

// Check if a player has won.
func CheckWinner(player string) bool {
    // definng winning combinations.
    winningCombos := [...][5]int{
        //horixontal, vertical, and diagonal combinations.
        {0, 1, 2, 3, 4}, {5, 6, 7, 8, 9}, {10, 11, 12, 13, 14},
        {15, 16, 17, 18, 19}, {20, 21, 22, 23, 24},
        {0, 5, 10, 15, 20}, {1, 6, 11, 16, 21}, {2, 7, 12, 17, 22},
        {3, 8, 13, 18, 23}, {4, 9, 14, 19, 24},
        {0, 6, 12, 18, 24}, {4, 8, 12, 16, 20},
    }
    // check each winning conbination.
    for _, combo := range winningCombos {
        if cells[combo[0]] == player && cells[combo[1]] == player &&
            cells[combo[2]] == player && cells[combo[3]] == player &&
            cells[combo[4]] == player {
            return true // if winning combination is found.
        }
    }
    return false // if no winning combination is found.
}

// Add a move to the board.
func AddMove(move int, player string) bool {
    // validate and place a move on the board.
    if move >= 1 && move <= 25 && cells[move-1] == "--" {
        cells[move-1] = player
        return true
    }
    return false
}
// main game loop.
func PlayGame() {
    InitializeBoard()
    // trun array to track the situation of player turns.
    turns := []string{"MG", "AI", "MG", "AI", "MG", "AI", "MG", "AI", "MG", "AI", "MG", "AI", "MG", "AI", "MG", "AI", "MG", "AI", "MG", "AI", "MG", "AI", "MG", "AI", "MG"}
    // the random numbber generator Seed
    rand.Seed(time.Now().UnixNano())

    for len(turns) > 0 {
        if len(turns)%2==1 {
            DisplayBoard() // display the board after turn played.
        }
        player := turns[0]
        turns = turns[1:]

        var move int
        if player == PlayerAI {
            // AI's move
            move = rand.Intn(25) + 1
        } else {
            // Player's move
            for {
                fmt.Printf("Player %s's turn. Enter move (1-25): ", player)
                _, err := fmt.Scan(&move)
                if err != nil || move < 1 || move > 25 {
                    fmt.Println("Invalid move, try again 1-25.")
                    continue
                }
                break
            }
        }
         // player movement and positions.
        if !AddMove(move, player) {
            fmt.Println("Position already taken, try a different position.")
            turns = append([]string{player}, turns...)
            continue
        }

        // Check for a winner
        if CheckWinner(player) {
            DisplayBoard()
            fmt.Printf("Player %s wins!\n", player)
            break
        }
    }
        // if none of the player wins.
        if !CheckWinner(PlayerMG) && !CheckWinner(PlayerAI) {
        fmt.Println("It's a tie!")
    }
}
func main(){
    for {
        PlayGame()
        // ask if the player want to play again.
        var restart string
        fmt.Print("Do you want play again? (yes/no);")
        fmt.Scan(&restart)
        restart = strings.ToLower(restart)
        // break the loop id the user does not want to play again.
        if restart != "yes" {
            fmt.Println("Game Over. Thank you For Playing!")
            break
        }
    }
}