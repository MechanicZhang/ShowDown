package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	TypePlayerNamePromt          = "請輸入第 %v 位玩家名稱: "
	ExchangeHandsPrompt          = "請問您要使用「交換手牌」特權嗎？(y/n)"
	ChoosePlayerPrompt           = "請問您要和哪一位玩家交換手牌呢？"
	ExchangeHandsCompleteMessage = "手牌已經交換完畢，將在 ? 回合後換回。"
	ChooseCardIndexPrompt        = "請輸入您要出的牌的編號："
)

func promptUser(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func promptPlayerName(prompt string) (string, error) {
	for {
		response, err := promptUser(prompt)
		if err != nil {
			return "", err
		}
		if len(response) == 0 {
			fmt.Println("請輸入有效的名字")
			continue
		}
		return response, nil
	}
}

func promptYesNo(prompt string) (bool, error) {
	for {
		response, err := promptUser(prompt)
		if err != nil {
			return false, err
		}

		if response == "y" {
			return true, nil
		} else if response == "n" {
			return false, nil
		} else {
			fmt.Println("請輸入 y 或 n")
		}
	}
}

func promptPlayerIndex(players []IPlayer) (int, error) {
	fmt.Printf("請選擇一位玩家：(1:%v, 2:%v 3:%v)\n", players[0].GetName(), players[1].GetName(), players[2].GetName())
	for {
		response, err := promptUser(ChoosePlayerPrompt)
		if err != nil {
			return -1, err
		}
		index, err := strconv.Atoi(response)
		if err != nil {
			fmt.Println("請輸入有效的數字")
			continue
		}
		if index <= 0 || index > len(players) {
			fmt.Printf("請輸入 1 到 %d 之間的數字\n", len(players))
			continue
		}
		return index, nil
	}
}

func promptCardIndex(cards []Card) (int, error) {
	printCardSeletion(cards)
	for {
		response, err := promptUser(ChooseCardIndexPrompt)
		if err != nil {
			return -1, err
		}
		index, err := strconv.Atoi(response)
		if err != nil {
			fmt.Println("請輸入有效的數字")
			continue
		}
		if index <= 0 || index > len(cards) {
			fmt.Printf("請輸入 1 到 %d 之間的數字\n", len(cards))
			continue
		}
		return index, nil
	}
}

func printCardSeletion(cards []Card) {
	fmt.Println("請問您要出第幾張牌呢？")
	fmt.Printf("   ")
	for _, card := range cards {
		fmt.Printf(" %v", card)
	}
	fmt.Println()
}
