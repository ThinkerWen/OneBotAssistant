package sensitive

import (
	"OneBotAssistant/config"
	"OneBotAssistant/util"
	"database/sql"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/tidwall/gjson"
	"net/url"
	"strings"
)

func isSensitive(content string) bool {
	headers := make(map[string]string)
	headers["Connection"] = "keep-alive"
	headers["Accept-Language"] = "zh-CN,zh;q=0.9"
	headers["Origin"] = "http://www.zhipaiwu.com"
	headers["X-Requested-With"] = "XMLHttpRequest"
	headers["Accept"] = "application/json, text/javascript, */*; q=0.01"
	headers["Referer"] = "http://www.zhipaiwu.com/index.php/Weijinci/index.html"
	headers["Content-Type"] = "application/x-www-form-urlencoded; charset=UTF-8"
	headers["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36"

	data := fmt.Sprintf(`content=%s`, url.QueryEscape(content))
	response, err := util.RequestPOST("http://www.zhipaiwu.com/index.php/Weijinci/postIndex.html", data, headers, nil)
	if err != nil || gjson.Get(string(response), "code").Int() != 200 {
		log.Error("敏感词检测请求异常", "err", err)
		return false
	}
	if gjson.Get(string(response), "result.minganCount").Int() != 0 {
		return true
	}
	return false
}

func isForbiddenKeyword(keyword string, rangeId int64, rangeType int) bool {
	var sqlStr string
	if rangeType == 0 {
		sqlStr = `SELECT sensitive_word FROM sensitive_words WHERE group_id = '*' OR group_id = ?`
	} else if rangeType == 1 {
		sqlStr = `SELECT sensitive_word FROM sensitive_words WHERE user_id = '*' OR user_id = ?`
	} else {
		return false
	}

	rows, err := config.DB.Query(sqlStr, rangeId)
	if err != nil {
		log.Error(err)
		return false
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	keywords := make([]string, 0)
	for rows.Next() {
		var sensitiveWord string
		err = rows.Scan(&sensitiveWord)
		if err != nil {
			log.Error(err)
			continue
		}
		keywords = append(keywords, sensitiveWord)
	}

	for _, k := range keywords {
		if strings.Contains(keyword, k) {
			return true
		}
	}

	return false
}

func addSensitiveKeyword(keyword, rangeId string, rangeType int) bool {
	var sqlStr string
	if rangeType == 0 {
		sqlStr = `INSERT INTO sensitive_words (sensitive_word, group_id) VALUES (?, ?)`
	} else if rangeType == 1 {
		sqlStr = `INSERT INTO sensitive_words (sensitive_word, user_id) VALUES (?, ?)`
	} else {
		return false
	}

	_, err := config.DB.Exec(sqlStr, keyword, rangeId)
	if err != nil {
		log.Error(err)
		return false
	}

	return true
}
