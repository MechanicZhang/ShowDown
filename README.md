# A simple poker game - showdown
這是我用來練習物件導向程式設計的小專案，有幾個值得一提的關注點：

1. 乾淨的遊戲流程、乾淨的 prompt function。
1. HumanPlayer 和 AI 都需要實現 IPlayer 介面，但他們有大量的共同行為，因此我將這些行為交給 PlayerBase 來實作，然後讓 HumanPlayer, AI 去複合 PlayerBase。
1. 遊戲封裝每個玩家的動作為 Action。
1. 關聯類別「交換手牌」的實作。

尚未解決的 force：
1. HumanPlayer 和 AI 的 `TakeTurn()` 仍有重複的程式碼，但又不能放到 PlayerBase 中，因為 PlayerBase 無法呼叫 HumanPlayer 和 AI 實現的 `ShowCard()` & `MakeExchange...()`，這在有抽象類別的語言中很自然能做到的事情，在只有 interface 和 struct 的 golang 中似乎沒有那麼自然。

