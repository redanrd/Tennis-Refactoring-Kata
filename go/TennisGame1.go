package tenniskata

type tennisGame1 struct {
	mScore1     int
	mScore2     int
	player1Name string
	player2Name string
}

func TennisGame1(player1Name string, player2Name string) TennisGame {
	return &tennisGame1{
		player1Name: player1Name,
		player2Name: player2Name,
	}
}

func (game *tennisGame1) WonPoint(playerName string) {
	if playerName == "player1" {
		game.mScore1 += 1
	} else {
		game.mScore2 += 1
	}
}

func (game *tennisGame1) GetScore() string {
	if game.mScore1 == game.mScore2 {
		return tennisScoreEqual(game.mScore1)

	} else if game.mScore1 >= 4 || game.mScore2 >= 4 {
		return tennisScoreWithAdvantageOrWin(game.mScore1, game.mScore2)

	} else {
		return tennisScore(game.mScore1, game.mScore2)
	}
}

func tennisScoreEqual(score int) string {
	switch score {
	case 0:
		return "Love-All"

	case 1:
		return "Fifteen-All"

	case 2:
		return "Thirty-All"

	default:
		return "Deuce"
	}
}

func tennisScoreWithAdvantageOrWin(score1, score2 int) string {
	minusResult := score1 - score2
	switch {
	case minusResult == 1:
		return "Advantage player1"

	case minusResult == -1:
		return "Advantage player2"

	case minusResult >= 2:
		return "Win for player1"

	default:
		return "Win for player2"
	}
}

func tennisScore(score1, score2 int) string {
	var score string

	for i := 1; i < 3; i++ {
		tempScore := score1
		if i == 2 {
			score += "-"
			tempScore = score2
		}

		score += tennisScoreIncrement(tempScore)
	}

	return score
}

func tennisScoreIncrement(tempScore int) string {
	switch tempScore {
	case 0:
		return "Love"
	case 1:
		return "Fifteen"
	case 2:
		return "Thirty"
	case 3:
		return "Forty"
	default:
		return ""
	}
}
