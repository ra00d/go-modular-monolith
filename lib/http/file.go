package http

type FileResponse struct {
	Name string
	Data []byte
	Type string
}

func (c *BaseController) File(
	name string,
	data []byte,
	contentType string,
) FileResponse {
	return FileResponse{
		Name: name,
		Data: data,
		Type: contentType,
	}
}
