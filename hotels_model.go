package main

//TOML output format requires slice to be embedded in a struct
type Hotels struct {
	Hotels []Hotel
}

type Hotel struct {
	Name    string
	Address string
	Stars   int
	Contact string
	Phone   string
	Uri     string
}

type DataWriter interface {
	GetWritableData() Hotels
}

func (hotels *Hotels) GetWritableData() Hotels {
	return *hotels
}
