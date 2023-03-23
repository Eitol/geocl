package geocl

type Location struct {
	// LOCCode is a unique identifier.
	// There is only one Location per LOCCode.
	// LOCCode is a 3-character string. e.g. "ACZ"
	LOCCode string `json:"locCode"`
	LOCName string `json:"locName"` // LOCName is the name of the location. e.g. "CODPA"
	RegCode string `json:"regCode"` // RegCode is the code of the region. e.g. "15"
	RegName string `json:"regName"` // RegName is the name of the region. e.g. "Arica y Parinacota"
	COMCode string `json:"comCode"` // COMCode is the code of the commune. e.g. "15101"
	COMName string `json:"comName"` // COMName is the name of the commune. e.g. "ARICA"
}
