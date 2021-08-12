package interaction

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
)

type RoundData struct {
	Action           string
	PlayerAttackDmg  int
	PlayerHealValue  int
	MonsterAttackDmg int
	Playerhealth     int
	MonsterHealth    int
}

func PrintGreeting() {
	asciiFigure := figure.NewColorFigure("MONSTER SLAYER", "", "yellow", true)
	asciiFigure.Print()
	fmt.Println("MONSTER SLAYER")
	fmt.Println("STARTING A NEW GAME....")
	fmt.Println("GOOD LUCK")

}

func ShowAvailableActions(specialAttackIsAvailable bool) {
	fmt.Println("Please choose your action")
	fmt.Println("----------------------------")
	fmt.Println("(1) Attack Monster")
	fmt.Println("(2) Heal")

	if specialAttackIsAvailable {
		fmt.Println("(3) Special Attack")

	}

}

func PrintRoundStatitics(roundData *RoundData) {
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player attacked Monster for %v damage.\n", roundData.PlayerAttackDmg)
	} else if roundData.Action == "SPECIAL ATTACK" {
		fmt.Printf("Player performed a strong attack against Monster for %v damage.\n", roundData.PlayerAttackDmg)

	} else {
		fmt.Printf("Player healed for %v .\n", roundData.PlayerHealValue)

	}
	fmt.Printf("Monster attacked Player for %v damage.\n", roundData.MonsterAttackDmg)
	fmt.Printf("Player Health: %v \n", roundData.Playerhealth)
	fmt.Printf("Monster Health: %v \n", roundData.MonsterHealth)

}

func DeclareWinner(winner string) {
	fmt.Println("----------------------------")
	asciiFigure := figure.NewColorFigure("GAME OVER!", "", "red", true)
	asciiFigure.Print()
	fmt.Println("----------------------------")
	fmt.Printf("%v won ! \n", winner)
}

func WriteLogFile(rounds *[]RoundData) {

	//To make an executable
	// exPath, err := os.Executable()

	// if err != nil {
	// 	fmt.Println("Saving a log file failed. eciting")
	// 	return
	// }

	// exPath = filepath.Dir(exPath)

	// file, err := os.Create(exPath + "/gamelog.txt")
	file, err := os.Create("gamelog.txt") // Require For "go run"

	if err != nil {
		fmt.Println("Saving a log file failes. eciting")
		return

	}

	for index, value := range *rounds {
		logEntry := map[string]string{
			"Round":                 fmt.Sprint(index + 1),
			"Action":                value.Action,
			"Player Attack Damage":  fmt.Sprint(value.PlayerAttackDmg),
			"Player Heal Value":     fmt.Sprint(value.PlayerHealValue),
			"Monster Attack Damage": fmt.Sprint(value.MonsterAttackDmg),
			"Player Health":         fmt.Sprint(value.Playerhealth),
			"Monster Health":        fmt.Sprint(value.MonsterHealth),
		}
		logLine := fmt.Sprintln(logEntry)
		_, err = file.WriteString(logLine)

		if err != nil {
			fmt.Println("Writing into log file failed. Exiting")
			continue
		}

	}

	file.Close()
	fmt.Println("Wrote Dato to log!")

}
