package applications

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

type commonResult struct {
	golangsdk.Result
}

type CreateResp struct {
	// Creator of the application.
	//     USER: The app is created by the API user.
	//     MARKET: The app is allocated by the marketplace.
	Creator string `json:"creator"`
	// Update time.
	UdpateTime string `json:"update_time"`
	// App key.
	AppKey string `json:"app_key"`
	// App name.
	Name string `json:"name"`
	// Description.
	Description string `json:"description"`
	// ID.
	Id string `json:"id"`
	// App secret.
	AppSecret string `json:"app_secret"`
}

type CreateResult struct {
	commonResult
}

func (r CreateResult) Extract() (*CreateResp, error) {
	var s CreateResp
	err := r.ExtractInto(&s)
	return &s, err
}

type Application struct {
	// Creator of the application.
	//     USER: The app is created by the API user.
	//     MARKET: The app is allocated by the marketplace.
	Creator string `json:"creator"`
	// Registraion time.
	RegistraionTime string `json:"register_time"`
	// Update time.
	UpdateTime string `json:"update_time"`
	// App key.
	AppKey string `json:"app_key"`
	// App name.
	Name string `json:"name"`
	// Description.
	Description string `json:"description"`
	// ID.
	Id string `json:"id"`
	// App secret.
	AppSecret string `json:"app_secret"`
	// App status.
	Status string `json:"status"`
	// App type. List method are not support.
	AppType string `json:"app_type"`
	// Number of APIs. Only used for List method.
	BindNum string `json:"bind_num"`
}

type GetResult struct {
	commonResult
}

type UpdateResult struct {
	commonResult
}

type ResetAppSecretResult struct {
	commonResult
}

func (r commonResult) Extract() (*Application, error) {
	var s Application
	err := r.ExtractInto(&s)
	return &s, err
}

type ApplicationPage struct {
	pagination.SinglePageBase
}

func ExtractInstances(r pagination.Page) ([]Application, error) {
	var s []Application
	err := r.(ApplicationPage).Result.ExtractIntoSlicePtr(&s, "apps")
	return s, err
}

type DeleteResult struct {
	golangsdk.ErrResult
}

type AppCode struct {
	// AppCode value, which contains 64 to 180 characters, starting with a letter, plus sign (+) or slash (/).
	// Only letters and the following special characters are allowed: +-_!@#$%/=
	Code string `json:"app_code"`
	// AppCode ID.
	Id string `json:"id"`
	// App ID.
	AppId string `json:"app_id"`
	// Creation time, in UTC format.
	CreateTime string `json:"create_time"`
}

type CodeResult struct {
	commonResult
}

func (r CodeResult) Extract() (*AppCode, error) {
	var s AppCode
	err := r.ExtractInto(&s)
	return &s, err
}

type AppCodePage struct {
	pagination.SinglePageBase
}

func ExtractAppCodes(r pagination.Page) ([]AppCode, error) {
	var s []AppCode
	err := r.(AppCodePage).Result.ExtractIntoSlicePtr(&s, "app_codes")
	return s, err
}
