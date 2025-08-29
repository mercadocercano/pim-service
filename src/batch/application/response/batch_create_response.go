package response

// BatchCreateResponse representa la respuesta de la creación en batch
type BatchCreateResponse struct {
	Created BatchCreatedItems `json:"created"`
	Errors  []BatchError      `json:"errors"`
}

// BatchCreatedItems contiene los IDs de las entidades creadas
type BatchCreatedItems struct {
	Categories []string `json:"categories"`
	Brands     []string `json:"brands"`
	Products   []string `json:"products"`
}

// BatchError representa un error durante la creación batch
type BatchError struct {
	Type  string `json:"type"`  // "category", "brand", "product"
	Name  string `json:"name"`
	Error string `json:"error"`
}