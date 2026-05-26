package resource

type FileResource struct {
	Resource
}

func (fileResource *FileResource) Read(name string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
