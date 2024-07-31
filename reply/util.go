package reply

import (
	"OneBotAssistant/config"
	"database/sql"
	"github.com/charmbracelet/log"
	"strings"
)

type Sequence struct {
	Sender   int64
	Receiver int64
	State    AutoReplyState
}

type AutoReplyState struct {
	Step      int
	Ask       string
	Answer    string
	Range     string
	RangeType int
}

func matchAsk(ask string, rangeId int64, rangeType int) string {
	var sqlStr string
	if rangeType == 0 {
		sqlStr = `SELECT ask, reply FROM auto_reply WHERE group_id = ? OR group_id = '*'`
	} else if rangeType == 1 {
		sqlStr = `SELECT ask, reply FROM auto_reply WHERE user_id = ? OR user_id = '*'`
	}

	rows, err := config.DB.Query(sqlStr, rangeId)
	if err != nil {
		log.Error(err)
		return ""
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	replyList := make([]map[string]string, 0)
	for rows.Next() {
		var ask, reply string
		err = rows.Scan(&ask, &reply)
		if err != nil {
			log.Error(err)
			continue
		}
		replyList = append(replyList, map[string]string{"ask": ask, "reply": reply})
	}

	for _, v := range replyList {
		if matchPattern(ask, v["ask"]) {
			return v["reply"]
		}
	}

	return ""
}

func matchPattern(str, pattern string) bool {
	if pattern == "*" {
		return true
	} else if strings.HasPrefix(pattern, "*") && strings.HasSuffix(pattern, "*") {
		substr := strings.TrimPrefix(pattern, "*")
		substr = strings.TrimSuffix(substr, "*")
		return strings.Contains(str, substr)
	} else if strings.HasPrefix(pattern, "*") {
		suffix := strings.TrimPrefix(pattern, "*")
		return strings.HasSuffix(str, suffix)
	} else if strings.HasSuffix(pattern, "*") {
		prefix := strings.TrimSuffix(pattern, "*")
		return strings.HasPrefix(str, prefix)
	} else {
		return str == pattern
	}
}

func checkReply(msg, rawMessage string, state *AutoReplyState) int {
	switch state.Step {
	case 0:
		state.Ask = msg
		state.Step = 1
	case 1:
		state.Answer = rawMessage
		state.Step = 2
	case 2:
		if msg == "1" || msg == "2" {
			state.Step = 3
		}
	default:
		state.Step = 0
	}
	return state.Step
}

func saveReply(s Sequence) bool {
	var sqlStr string
	if s.State.RangeType == 0 {
		sqlStr = `INSERT INTO auto_reply (ask, reply, group_id) VALUES (?, ?, ?)`
	} else if s.State.RangeType == 1 {
		sqlStr = `INSERT INTO auto_reply (ask, reply, user_id) VALUES (?, ?, ?)`
	}
	if _, err := config.DB.Exec(sqlStr, s.State.Ask, s.State.Answer, s.State.Range); err != nil {
		return false
	}
	return true
}

func removeSequence(sender, receiver int64) []Sequence {
	data := make([]Sequence, 0)
	for _, v := range sequence {
		if v.Sender == sender && v.Receiver == receiver {
			continue
		}
		data = append(data, v)
	}
	return data
}
