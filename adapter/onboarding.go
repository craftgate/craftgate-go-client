package adapter

import (
	"craftgate-go-client/adapter/rest"
	"craftgate-go-client/model"
	"fmt"
	"net/http"
)

type Onboarding struct {
	Opts model.RequestOptions
}

type CreateMemberRequest struct {
	MemberExternalId              string                              `json:"memberExternalId,omitempty"`
	Name                          string                              `json:"name,omitempty"`
	Address                       string                              `json:"address,omitempty"`
	Email                         string                              `json:"email,omitempty"`
	PhoneNumber                   string                              `json:"phoneNumber,omitempty"`
	IdentityNumber                string                              `json:"identityNumber,omitempty"`
	ContactName                   string                              `json:"contactName,omitempty"`
	ContactSurname                string                              `json:"contactSurname,omitempty"`
	MemberType                    model.MemberType                    `json:"memberType,omitempty"`
	LegalCompanyTitle             string                              `json:"legalCompanyTitle,omitempty"`
	TaxOffice                     string                              `json:"taxOffice,omitempty"`
	TaxNumber                     string                              `json:"taxNumber,omitempty"`
	Iban                          string                              `json:"iban,omitempty"`
	SettlementEarningsDestination model.SettlementEarningsDestination `json:"settlementEarningsDestination,omitempty"`
	NegativeWalletAmountLimit     float64                             `json:"negativeWalletAmountLimit,omitempty"`
	IsBuyer                       bool                                `json:"isBuyer,omitempty"`
	IsSubMerchant                 bool                                `json:"isSubMerchant,omitempty"`
}

type UpdateMemberRequest struct {
	Name                          string                              `json:"name,omitempty"`
	Address                       string                              `json:"address,omitempty"`
	Email                         string                              `json:"email,omitempty"`
	PhoneNumber                   string                              `json:"phoneNumber,omitempty"`
	IdentityNumber                string                              `json:"identityNumber,omitempty"`
	ContactName                   string                              `json:"contactName,omitempty"`
	ContactSurname                string                              `json:"contactSurname,omitempty"`
	MemberType                    model.MemberType                    `json:"memberType,omitempty"`
	LegalCompanyTitle             string                              `json:"legalCompanyTitle,omitempty"`
	TaxOffice                     string                              `json:"taxOffice,omitempty"`
	TaxNumber                     string                              `json:"taxNumber,omitempty"`
	Iban                          string                              `json:"iban,omitempty"`
	SettlementEarningsDestination model.SettlementEarningsDestination `json:"settlementEarningsDestination,omitempty"`
	NegativeWalletAmountLimit     float64                             `json:"negativeWalletAmountLimit,omitempty"`
	IsBuyer                       bool                                `json:"isBuyer,omitempty"`
	IsSubMerchant                 bool                                `json:"isSubMerchant,omitempty"`
}

type SearchMembersRequest struct {
	Page             int              `schema:"page,omitempty"`
	Size             int              `schema:"size,omitempty"`
	IsBuyer          bool             `schema:"isBuyer,omitempty"`
	IsSubMerchant    bool             `schema:"isSubMerchant,omitempty"`
	Name             string           `schema:"name,omitempty"`
	MemberIds        []int64          `schema:"memberIds,omitempty"`
	MemberType       model.MemberType `schema:"memberType,omitempty"`
	MemberExternalId string           `schema:"memberExternalId,omitempty"`
}

type MemberResponse struct {
	Id                            int64                               `json:"id"`
	CreatedDate                   CraftgateTime                       `json:"createdDate"`
	UpdatedDate                   CraftgateTime                       `json:"updatedDate"`
	Status                        model.Status                        `json:"status"`
	IsBuyer                       bool                                `json:"isBuyer"`
	IsSubMerchant                 bool                                `json:"isSubMerchant"`
	MemberType                    model.MemberType                    `json:"memberType"`
	MemberExternalId              string                              `json:"memberExternalId"`
	Name                          string                              `json:"name"`
	Email                         string                              `json:"email"`
	Address                       string                              `json:"address"`
	PhoneNumber                   string                              `json:"phoneNumber"`
	IdentityNumber                string                              `json:"identityNumber"`
	ContactName                   string                              `json:"contactName"`
	ContactSurname                string                              `json:"contactSurname"`
	LegalCompanyTitle             string                              `json:"legalCompanyTitle"`
	TaxOffice                     string                              `json:"taxOffice"`
	TaxNumber                     string                              `json:"taxNumber"`
	SettlementEarningsDestination model.SettlementEarningsDestination `json:"settlementEarningsDestination"`
	NegativeWalletAmountLimit     float64                             `json:"negativeWalletAmountLimit"`
	Iban                          string                              `json:"iban"`
}

func (api *Onboarding) CreateMember(request CreateMemberRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/onboarding/v1/members", api.Opts.BaseURL), body)

	res := model.Response[MemberResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Onboarding) UpdateMember(id int64, request UpdateMemberRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("PUT", fmt.Sprintf("%s/onboarding/v1/members/%d", api.Opts.BaseURL, id), body)

	res := model.Response[MemberResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Onboarding) RetrieveMember(id int64) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/onboarding/v1/members/%d", api.Opts.BaseURL, id), nil)

	res := model.Response[MemberResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Onboarding) SearchMembers(request SearchMembersRequest) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/onboarding/v1/members", api.Opts.BaseURL), nil)
	req.URL.RawQuery, _ = QueryParams(request)

	res := model.Response[model.DataResponse[MemberResponse]]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}
