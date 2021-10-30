package response

import "AltaStore/business/shopping"

func GetOneResponse(data *shopping.ShoppCart) *ShoppData {
	return &ShoppData{
		ID:         data.ID,
		IsCheckOut: data.IsCheckOut,
		UpdatedAt:  data.UpdatedAt,
	}
}
