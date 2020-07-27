package util

type ConvData struct {
	Input  []byte
	Output []byte
}

type Converter interface {
	Convert() (err error)
}

func (data *ConvData) Convert() (err error) {

	return err
}
