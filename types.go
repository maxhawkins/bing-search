package bing

// Result contains the results of a web search service operation
type Result struct {
	// Identifier
	ID string
	// Text specified in the HTML <title> tag of the page
	Title string
	// Description text of the web result
	Description string
	// Web URL to display to the user
	DisplayURL string
	// Full URL of the web result
	URL string `json:"Url"`
}
