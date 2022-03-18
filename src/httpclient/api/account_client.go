package api

import (
	"encoding/json"
	"fmt"
	"github.com/gostudys/huobi_contract_golang/src/config"
	"github.com/gostudys/huobi_contract_golang/src/httpclient/response"
	"github.com/gostudys/huobi_contract_golang/src/logh"
	"github.com/gostudys/huobi_contract_golang/src/utils/reqbuilder"
)

type AccountClient struct {
	PUrlBuilder *reqbuilder.PrivateUrlBuilder
}

func (ac *AccountClient) Init(accessKey string, secretKey string, host string) *AccountClient {
	if host == "" {
		host = config.LINEAR_SWAP_DEFAULT_HOST
	}
	ac.PUrlBuilder = new(reqbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	return ac
}

func (ac *AccountClient) GetBalanceValuationAsync(data chan response.GetBalanceValuationResponse, valuationAsset string) {
	// ulr
	url := ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_balance_valuation", nil)

	// content
	content := ""
	if valuationAsset != "" {
		content = fmt.Sprintf(",\"valuation_asset\": \"%s\"", valuationAsset)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetBalanceValuationResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetBalanceValuationResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) IsolatedGetAccountInfoAsync(data chan response.GetAccountInfoResponse, contractCode string, subUid int64) {
	// ulr
	url := ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_account_info", nil)
	if subUid != 0 {
		url = ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_sub_account_info", nil)
	}

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if subUid != 0 {
		content += fmt.Sprintf(",\"sub_uid\": %d", subUid)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetAccountInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetAccountInfoResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) CrossGetAccountInfoAsync(data chan response.GetAccountInfoResponse, marginAccount string, subUid int64) {
	// ulr
	url := ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_cross_account_info", nil)
	if subUid != 0 {
		url = ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_cross_sub_account_info", nil)
	}

	// content
	content := ""
	if marginAccount != "" {
		content = fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	}
	if subUid != 0 {
		content += fmt.Sprintf(",\"sub_uid\": %d", subUid)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetAccountInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetAccountInfoResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) IsolatedGetAccountPositionAsync(data chan response.GetAccountPositionResponse, contractCode string, subUid int64) {
	// ulr
	url := ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_position_info", nil)
	if subUid != 0 {
		url = ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_sub_account_info", nil)
	}

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if subUid != 0 {
		content += fmt.Sprintf(",\"sub_uid\": %d", subUid)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetAccountPositionResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetAccountPositionResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) CrossGetAccountPositionAsync(data chan response.GetAccountPositionResponse, contractCode string, subUid int64) {
	// ulr
	url := ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_cross_position_info", nil)
	if subUid != 0 {
		url = ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_cross_sub_position_info", nil)
	}

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if subUid != 0 {
		content += fmt.Sprintf(",\"sub_uid\": %d", subUid)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetAccountPositionResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetAccountPositionResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) IsolatedGetAssetsPositionAsync(data chan response.GetAssetsPositionResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_account_position_info", nil)

	// content
	content := fmt.Sprintf("{\"contract_code\": \"%s\"}", contractCode)

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetAssetsPositionResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetAssetsPositionResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) CrossGetAssetsPositionAsync(data chan response.GetAssetsPositionResponseSingle, marginAccount string) {
	// ulr
	url := ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_cross_account_position_info", nil)

	// content
	content := fmt.Sprintf("{\"margin_account\": \"%s\"}", marginAccount)

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetAssetsPositionResponseSingle{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetAssetsPositionResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) SetSubAuthAsync(data chan response.SetSubAuthResponse, subUid string, subAuth int) {
	// ulr
	url := ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_sub_auth", nil)

	// content
	content := fmt.Sprintf("{\"sub_uid\": \"%s\", \"sub_auth\": %d}", subUid, subAuth)

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.SetSubAuthResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to SetSubAuthResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) IsolatedGetSubAccountListResponseAsync(data chan response.GetSubAccountListResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_sub_account_list", nil)

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetSubAccountListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetAllSubAssetsResponse error: %s", getErr)
	}
	data <- result
}

func (ac *AccountClient) CrossGetSubAccountListAsync(data chan response.GetSubAccountListResponse, marginAccount string) {
	// ulr
	url := ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_cross_sub_account_list", nil)

	// content
	content := ""
	if marginAccount != "" {
		content = fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetSubAccountListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetSubAccountListResponse error: %s", getErr)
	}
	data <- result
}

func (ac *AccountClient) IsolatedGetSubAccountInfoListAsync(data chan response.GetSubAccountInfoListResponse,
	contractCode string, pageIndex int, pageSize int) {
	// ulr
	url := ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_sub_account_info_list", nil)

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if pageIndex != 0 {
		content = fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content = fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetSubAccountInfoListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetSubAccountInfoListResponse error: %s", getErr)
	}
	data <- result
}

func (ac *AccountClient) CrossGetSubAccountInfoListAsync(data chan response.GetSubAccountInfoListResponse,
	marginAccount string, pageIndex int, pageSize int) {
	// ulr
	url := ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_cross_sub_account_info_list", nil)

	// content
	content := ""
	if marginAccount != "" {
		content = fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	}
	if pageIndex != 0 {
		content = fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content = fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetSubAccountInfoListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetSubAccountInfoListResponse error: %s", getErr)
	}
	data <- result
}

func (ac *AccountClient) AccountTransferAsync(data chan response.AccountTransferResponse, asset string, fromMarginAccount string, toMarginAccount string, amount float64,
	subUid int64, fcType string) {
	// ulr
	url := ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_master_sub_transfer", nil)
	if subUid == 0 {
		url = ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_transfer_inner", nil)
	}

	// content
	content := fmt.Sprintf(",\"asset\":\"%s\", \"from_margin_account\":\"%s\", \"to_margin_account\":\"%s\", \"amount\":%f",
		asset, fromMarginAccount, toMarginAccount, amount)
	if subUid != 0 {
		content += fmt.Sprintf(",\"sub_uid\": %d,\"type\": \"%s\"", subUid, fcType)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.AccountTransferResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to AccountTransferResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) GetAccountTransHisAsync(data chan response.GetAccountTransHisResponse, marginAccount string,
	beMasterSub bool, fcType string, createDate int, pageIndex int, pageSize int) {
	// ulr
	url := ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_financial_record", nil)
	if beMasterSub {
		url = ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_master_sub_transfer_record", nil)
	}

	// content
	content := fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	if fcType != "" {
		content += fmt.Sprintf(",\"type\": \"%s\"", fcType)
		if beMasterSub {
			content += fmt.Sprintf(",\"transfer_type\": \"%s\"", fcType)
		}
	}
	if createDate != 0 {
		content += fmt.Sprintf(",\"create_date\": %d", createDate)
	} else {
		if beMasterSub {
			createDate = 7
			content += fmt.Sprintf(",\"create_date\": %d", createDate)
		}
	}
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetAccountTransHisResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetAccountTransHisResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) GetFinancialRecordExactAsync(data chan response.GetFinancialRecordExactResponse, marginAccount string,
	contractCode string, fcType string, startTime int64, endTime int64, fromId int64, size int, direct string) {
	// ulr
	url := ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_financial_record_exact", nil)

	// content
	content := fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if fcType != "" {
		content += fmt.Sprintf(",\"type\": \"%s\"", fcType)
	}
	if startTime != 0 {
		content += fmt.Sprintf(",\"start_time\": %d", startTime)
	}
	if endTime != 0 {
		content += fmt.Sprintf(",\"end_time\": %d", endTime)
	}
	if fromId != 0 {
		content += fmt.Sprintf(",\"from_id\": %d", fromId)
	}
	if size != 0 {
		content += fmt.Sprintf(",\"size\": %d", size)
	}
	if direct != "" {
		content += fmt.Sprintf(",\"direct\": \"%s\"", direct)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetFinancialRecordExactResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetFinancialRecordExactResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) IsolatedGetSettlementRecordsAsync(data chan response.IsolatedGetSettlementRecordsResponse, contractCode string,
	startTime int64, endTime int64, pageIndex int, pageSize int) {
	// ulr
	url := ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_user_settlement_records", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if startTime != 0 {
		content += fmt.Sprintf(",\"start_time\": %d", startTime)
	}
	if endTime != 0 {
		content += fmt.Sprintf(",\"end_time\": %d", endTime)
	}
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.IsolatedGetSettlementRecordsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to IsolatedGetSettlementRecordsResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) CrossGetSettlementRecordsAsync(data chan response.CrossGetSettlementRecordsResponse, marginAccount string,
	startTime int64, endTime int64, pageIndex int, pageSize int) {
	// ulr
	url := ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_cross_user_settlement_records", nil)

	// content
	content := fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	if startTime != 0 {
		content += fmt.Sprintf(",\"start_time\": %d", startTime)
	}
	if endTime != 0 {
		content += fmt.Sprintf(",\"end_time\": %d", endTime)
	}
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.CrossGetSettlementRecordsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to CrossGetSettlementRecordsResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) IsolatedGetValidLeverRateAsync(data chan response.GetValidLeverRateResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_available_level_rate", nil)

	// content
	content := fmt.Sprintf("{ \"contract_code\": \"%s\" }", contractCode)
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetValidLeverRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetValidLeverRateResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) CrossGetValidLeverRateAsync(data chan response.GetValidLeverRateResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_cross_available_level_rate", nil)

	// content
	content := fmt.Sprintf("{ \"contract_code\": \"%s\" }", contractCode)
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetValidLeverRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetValidLeverRateResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) GetOrderLimitAsync(data chan response.GetOrderLimitResponse, orderPriceType string, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_order_limit", nil)

	// content
	content := fmt.Sprintf(",\"order_price_type\":\"%s\"", orderPriceType)
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetOrderLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetOrderLimitResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) GetFeeAsync(data chan response.GetFeeResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_fee", nil)

	// content
	content := fmt.Sprintf("{ \"contract_code\": \"%s\" }", contractCode)
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetFeeResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetFeeResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) IsolatedGetTransferLimitAsync(data chan response.GetTransferLimitResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_transfer_limit", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetTransferLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetTransferLimitResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) CrossGetTransferLimitAsync(data chan response.GetTransferLimitResponse, marginAccount string) {
	// ulr
	url := ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_cross_transfer_limit", nil)

	// content
	content := ""
	if marginAccount != "" {
		content += fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetTransferLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetTransferLimitResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) IsolatedGetPositionLimitAsync(data chan response.GetPositionLimitResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_position_limit", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetPositionLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetPositionLimitResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) CrossGetPositionLimitAsync(data chan response.GetPositionLimitResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_cross_position_limit", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetPositionLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetPositionLimitResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) GetApiTradingStatusAsync(data chan response.GetApiTradingStatusResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(config.GET_METHOD, "/linear-swap-api/v1/swap_api_trading_status", nil)

	// content is nil
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetApiTradingStatusResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetApiTradingStatusResponse error: %s", jsonErr)
	}
	data <- result
}