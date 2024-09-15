package requests

type CreateContractRequest struct {

}

type GetTemplateByAddressRequest struct {
	Address []string `json:"address"`
}