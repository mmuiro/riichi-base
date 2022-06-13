package calculator

type ScoreLevel int

const (
	Low ScoreLevel = iota
	Mangan
	Haneman
	Baiman
	Sanbaiman
	Yakuman
)

var ScoreLevelToBasicPoints = map[ScoreLevel]int{
	Mangan:    2000,
	Haneman:   3000,
	Baiman:    4000,
	Sanbaiman: 6000,
	Yakuman:   8000,
}

var ScoreLevelToString = map[ScoreLevel]string{
	Low:       "",
	Mangan:    "Mangan",
	Haneman:   "Haneman",
	Baiman:    "Baiman",
	Sanbaiman: "Sanbaiman",
	Yakuman:   "Yakuman",
}

func HanToScoreLevel(han int) ScoreLevel {
	if han < 5 {
		return Low
	} else if han < 6 {
		return Mangan
	} else if han < 8 {
		return Haneman
	} else if han < 11 {
		return Baiman
	} else if han < 13 {
		return Sanbaiman
	}
	return Yakuman
}
