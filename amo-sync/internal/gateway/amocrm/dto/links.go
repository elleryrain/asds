package dto

type CreateLinksRequest struct {
	EntityType string
	EntityID   int
	Links      []CreateLinkRequest
}

type CreateLinkRequest struct {
	ToEntityID   int    `json:"to_entity_id"`
	ToEntityType string `json:"to_entity_type"`
}

type CreateLinksResponse struct {
	Links []CreateLinksResponseLinks `json:"links"`
}

type CreateLinksResponseLinks struct {
	ToEntityID   int    `json:"to_entity_id"`
	ToEntityType string `json:"to_entity_type"`
}
