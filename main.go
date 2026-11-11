package main

import (
	//"fmt"
	"strings"
)

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {

	for i := 0; i < len(messages); i++ {
		key := i
		value := messages[key]
		tags := tagger(value)
		//fmt.Print(tags)
		messages[i].tags = append(value.tags, tags...)
	}
	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	if strings.Contains(strings.ToLower(msg.content), "urgent") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(strings.ToLower(msg.content), "sale") {
		tags = append(tags, "Promo")
	}
	return tags
}

func main() {
	//fmt.Print(tagMessages([]sms{{id: "004", content: "Sale! Don't miss out on these urgent promotions!"}}, tagger))
	//fmt.Print(tagger(sms{id: "004", content: "Sale! Don't miss out on these urgent promotions!"}))
}
