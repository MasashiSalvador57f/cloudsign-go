package cloudsign

// PostDocumentRequest is request object for post document api.
type PostDocumentRequest struct {
	title       string
	note        string
	message     string
	templateID  string
	canTransfer bool
}
