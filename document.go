package cloudsign

// Document represents cloudsign document object.
type Document struct {
	ID              string        `json:"id"`
	UserID          string        `json:"user_id"`
	Title           string        `json:"title"`
	Note            string        `json:"note"`
	Message         string        `json:"message"`
	Status          int           `json:"status"`
	CanTransfer     bool          `json:"can_transfer"`
	SentAt          string        `json:"sent_at"`
	LastProcessedAt string        `json:"last_processed_at"`
	CreatedAt       string        `json:"created_at"`
	UpdatedAt       string        `json:"updated_at"`
	Participants    []Participant `json:"participants"`
	Files           []File        `json:"files"`
	Reportees       []Reportee    `json:"reportees"`
}

// Participant represents cloudsign participants object.
type Participant struct {
	ID              string `json:"id"`
	Email           string `json:"email"`
	Name            string `json:"name"`
	Organization    string `json:"organization"`
	Order           int    `json:"order"`
	Status          int    `json:"status"`
	AccessCode      string `json:"access_code"`
	LanguageCode    string `json:"language_code"`
	ProcessedAt     string `json:"processed_at"`
	AccessExpiresAt string `json:"access_expires_at"`
}

// Reportee represents cloudsign reportee object.
type Reportee struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Organization string `json:"organization"`
}

// Widget represents a component of cludsign docuemnts.
type Widget struct {
	ID            string `json:"id"`
	WidgetType    int    `json:"widget_type"`
	ParticipantID string `json:"participant_id"`
	FileID        string `json:"file_id"`
	Page          int    `json:"page"`
	X             int    `json:"x"`
	Y             int    `json:"y"`
	W             int    `json:"w"`
	H             int    `json:"h"`
	Text          string `json:"text"`
	Status        int    `json:"status"`
	Label         string `json:"label"`
	Required      bool   `json:"required"`
}

// File represents a cloudsign documentc file.
type File struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Order   int      `json:"order"`
	Widgets []Widget `json:"widgets"`
}
