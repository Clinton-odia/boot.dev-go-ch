package main

func addEmailsToQueue(emails []string) chan string {
	emailChan := make(chan string, len(emails))
	for _, emails := range emails {
		emailChan <- emails
	}
	return emailChan
}
