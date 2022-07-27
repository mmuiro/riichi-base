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

var YakumanMultiplierNamesEN = []string{"", "Double ", "Triple ", "Quadruple ", "Quintuple ", "Hextuple "}

var YakumanMultiplierNamesJA = []string{"", "二倍", "三倍", "四倍", "五倍", "六倍"}

var ScoreLevelToStringEN = map[ScoreLevel]string{
	Low:       "",
	Mangan:    "Mangan",
	Haneman:   "Haneman",
	Baiman:    "Baiman",
	Sanbaiman: "Sanbaiman",
	Yakuman:   "Yakuman",
}

var ScoreLevelToStringJA = map[ScoreLevel]string{
	Low:       "",
	Mangan:    "満貫",
	Haneman:   "跳満",
	Baiman:    "倍満",
	Sanbaiman: "三倍満",
	Yakuman:   "役満",
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
