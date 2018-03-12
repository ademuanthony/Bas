package resources

type EmailResource struct {
	From		string `json:"from"`
	To 		string `json:"to"`
	Title		string `json:"title"`
	Message 	string `json:"message"`
	Cc 		[]string `json:"cc"`
	IsHtml		bool	`json:"is_html"`
}
