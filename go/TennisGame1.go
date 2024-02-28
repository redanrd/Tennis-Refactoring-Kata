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
	var (
		score     string
		tempScore int
	)

	if game.mScore1 == game.mScore2 {
		score = tennisScoreEqual(game.mScore1)

	} else if game.mScore1 >= 4 || game.mScore2 >= 4 {
		minusResult := game.mScore1 - game.mScore2
		switch {
		case minusResult == 1:
			score = "Advantage player1"

		case minusResult == -1:
			score = "Advantage player2"

		case minusResult >= 2:
			score = "Win for player1"

		default:
			score = "Win for player2"
		}
	} else {
		for i := 1; i < 3; i++ {
			if i == 1 {
				tempScore = game.mScore1
			} else {
				score += "-"
				tempScore = game.mScore2
			}

			switch tempScore {
			case 0:
				score += "Love"
			case 1:
				score += "Fifteen"
			case 2:
				score += "Thirty"
			case 3:
				score += "Forty"
			}
		}
	}

	return score
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
