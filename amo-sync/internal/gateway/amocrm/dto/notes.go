package dto

type CreateNotesRequest struct {
	EntityType string `json:"entity_type"`
	Notes      []CreateNoteRequest
}

type CreateNoteRequest struct {
	EntityID int                     `json:"entity_id"`
	NoteType string                  `json:"note_type"`
	Params   CreateNoteRequestParams `json:"params"`
}

type CreateNoteRequestParams struct {
	Text string `json:"text"`
}
