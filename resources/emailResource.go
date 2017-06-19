package resources

type EmailResource struct {
	From		string 		`json:"from"`
	To 		string 		`json:"to"`
	Title		string 		`json:"title"`
	Message 	string 		`json:"message"`
	Cc 		[]string 	`json:"cc"`
	ContentType	string		`json:"content_type"`
}
