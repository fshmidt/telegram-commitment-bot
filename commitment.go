package telegram_commitment_bot

import (
	"github.com/fshmidt/telegram-commitment-bot/pkg/config"
	"strconv"
	"time"
)

type CommitStruct struct {
	UserID      uint64
	UserName    string
	ChatID      int64
	Commitment  string
	Created     time.Time
	Deadline    time.Time
	Ok          bool
	RoundRemind map[string]bool
	PercRemind  map[string]bool
	Scale       string
	CurReminder string
	Bucket      string
}

func (c *CommitStruct) MakeUis() string {
	return strconv.FormatUint(c.UserID, 10) + "*" + c.Commitment
}

func (c *CommitStruct) RoundDl(dl time.Duration, messages config.Messages) string {

	if dl <= time.Hour*17522 && dl >= time.Hour*17518 {
		return "Осталось 2 года!"
	} else if dl <= time.Hour*8762 && dl >= time.Hour*8758 {
		return "Остался год!"
	} else if dl <= time.Hour*4382 && dl >= time.Hour*4378 {
		return "Осталось полгода!"
	} else if dl <= time.Hour*2192 && dl >= time.Hour*2188 {
		return "Осталось 3 месяца!"
	} else if dl <= time.Hour*1442 && dl >= time.Hour*1438 {
		return "Осталось 2 месяца!"
	} else if dl <= time.Hour*722 && dl >= time.Hour*718 {
		return "Остался месяц!"
	} else if dl <= time.Hour*506 && dl >= time.Hour*502 {
		return "Осталось 3 недели!"
	} else if dl <= time.Hour*338 && dl >= time.Hour*334 {
		return "Осталось 2 недели!"
	} else if dl <= time.Hour*170 && dl >= time.Hour*166 {
		return "Осталась неделя!"
	} else if dl <= time.Hour*146 && dl >= time.Hour*142 {
		return "Осталось 6 дней!"
	} else if dl <= time.Hour*122 && dl >= time.Hour*118 {
		return "Осталось 5 дней!"
	} else if dl <= time.Hour*98 && dl >= time.Hour*94 {
		return "Осталось 4 дня!"
	} else if dl <= time.Hour*74 && dl >= time.Hour*70 {
		return "Осталось 3 дня!"
	} else if dl <= time.Hour*50 && dl >= time.Hour*46 {
		return "Осталось 2 дня!"
	} else if dl <= time.Hour*25 && dl >= time.Hour*23 {
		return "Остался ПОСЛЕДНИЙ ДЕНЬ!"
	} else if dl <= time.Hour*21 && dl >= time.Hour*19 {
		return "Осталось 20 часов!"
	} else if dl <= time.Hour*16 && dl >= time.Hour*14 {
		return "Осталось 15 часов!"
	} else if dl <= time.Hour*13 && dl > time.Hour*11 {
		return "Осталось 12 часов!"
	} else if dl <= time.Hour*11 && dl >= time.Hour*9 {
		return "Осталось 10 часов!"
	} else if dl <= time.Hour*6 && dl >= time.Hour*4 {
		return "Осталось 5 часов!"
	} else if dl <= time.Minute*190 && dl >= time.Minute*170 {
		return "Осталось 3 часа!"
	} else if dl <= time.Minute*130 && dl >= time.Minute*110 {
		return "Осталось 2 часа!"
	} else if dl <= time.Minute*70 && dl >= time.Minute*50 {
		return "Остался последний час!"
	}
	return ""
}

func (c *CommitStruct) Percents(dl time.Duration, messages config.Messages) (percents, scales string) {

	intPerc := int(time.Now().Sub(c.Created).Seconds() / c.Deadline.Sub(c.Created).Seconds() * 100)

	switch intPerc {
	case 1:
		return messages.Parts.Hundredth, ""
	case 10:
		return messages.Parts.Tenth, messages.Scales.Ten
	case 20:
		return messages.Parts.Fifth, messages.Scales.Twenty
	case 30:
		return "", messages.Scales.Thirty
	case 33:
		return messages.Parts.Third, messages.Scales.Thirty
	case 40:
		return "", messages.Scales.Forty
	case 50:
		return messages.Parts.Half, messages.Scales.Fifty
	case 60:
		return "", messages.Scales.Sixty
	case 67:
		return messages.Parts.SixtySix, messages.Scales.Seventy
	case 69:
		return messages.Parts.SixtyNine, messages.Scales.Seventy
	case 80:
		return messages.Parts.Eighty, messages.Scales.Eighty
	case 90:
		return messages.Parts.Ninety, messages.Scales.Ninety
	case 95:
		return messages.Parts.NinetyFive, messages.Scales.NinetyFive
	case 100:
		return messages.Parts.End, messages.Scales.End
	default:
		return "", ""
	}
}
