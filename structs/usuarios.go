package structs

type Usuario struct {
	usu_nom_chofer string `json: "usu_nom_chofer"`
	usu_nrut       int    `json:"usu_nrut"`
	usu_xdrut      string `json: "usu_xdrut"`
	usu_5kg        int    `json: "usu_5kg"`
	usu_c11kg      int    `json: "usu_c11kg"`
	usu_c15kg      int    `json: "usu_c15kg"`
	usu_c45kg      int    `json: "usu_c45kg"`
}
