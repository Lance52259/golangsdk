package applications

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

// AppOpts allows to create or update an application using given parameters.
type AppOpts struct {
	// Application name, which can contain 3 to 64 characters, starting with a letter.
	// Only letters, digits and underscores (_) are allowed.
	Name string `json:"name" required:"true"`
	// Description of the application, which can contain a maximum of 255 characters.
	// Chinese characters must be in UTF-8 or Unicode format.
	Description string `json:"remark,omitempty"`
	// Application key, which can contain 8 to 64 characters, starting with a letter or digit.
	// Only letters, digits, hyphens (-) and underscores (_) are allowed.
	AppKey string `json:"app_key,omitempty"`
	// Application secret, which can contain 8 to 64 characters, starting with a letter or digit.
	// Only letters, digits and the following special characters are allowed: _-!@#$%
	AppSecret string `json:"app_secret,omitempty"`
}

type AppOptsBuilder interface {
	ToAppOptsMap() (map[string]interface{}, error)
}

func (opts AppOpts) ToAppOptsMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create is a method by which to create function that create a APIG dedicated instance.
func Create(client *golangsdk.ServiceClient, instanceId string, opts AppOptsBuilder) (r CreateResult) {
	reqBody, err := opts.ToAppOptsMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(rootURL(client, instanceId), reqBody, &r.Body, nil)
	return
}

func Update(client *golangsdk.ServiceClient, instanceId, appId string, opts AppOptsBuilder) (r UpdateResult) {
	reqBody, err := opts.ToAppOptsMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(resourceURL(client, instanceId, appId), reqBody, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

//Get is a method to obtain the specified application according to the instanceId and appId.
func Get(client *golangsdk.ServiceClient, instanceId, appId string) (r GetResult) {
	_, r.Err = client.Get(resourceURL(client, instanceId, appId), &r.Body, nil)
	return
}

type ListOpts struct {
	// App ID.
	Id string `json:"id"`
	// App name.
	Name string `json:"name"`
	// App status.
	Status string `json:"status"`
	// App key.
	AppKey string `json:"app_key"`
	// Creator of the application.
	//     USER: The app is created by the API user.
	//     MARKET: The app is allocated by the marketplace.
	Creator string `json:"creator"`
	// Offset from which the query starts.
	// If the offset is less than 0, the value is automatically converted to 0. Default to 0.
	Offset int `q:"offset"`
	// Number of items displayed on each page.
	Limit int `q:"limit"`
	// Parameter name (name) for exact matching.
	PreciseSearch string `q:"precise_search"`
}

type ListOptsBuilder interface {
	ToAppListQuery() (string, error)
}

func (opts ListOpts) ToAppListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, instanceId string, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(client, instanceId)
	if opts != nil {
		query, err := opts.ToAppListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ApplicationPage{pagination.SinglePageBase(r)}
	})
}

type AppSecretResetOpts struct {
	AppSecret string `json:"app_secret"`
}

type AppSecretResetOptsBuilder interface {
	ToAppSecretResetOptsMap() (map[string]interface{}, error)
}

func (opts AppSecretResetOpts) ToAppSecretResetOptsMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func ResetAppSecret(client *golangsdk.ServiceClient, instanceId, appId string,
	opts AppSecretResetOptsBuilder) (r ResetAppSecretResult) {
	reqBody, err := opts.ToAppSecretResetOptsMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(resourceURL(client, instanceId, appId), reqBody, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

//Delete is a method to delete an existing application.
func Delete(client *golangsdk.ServiceClient, instanceId, appId string) (r DeleteResult) {
	_, r.Err = client.Delete(resourceURL(client, instanceId, appId), nil)
	return
}

type AppCodeOpts struct {
	// AppCode value, which contains 64 to 180 characters, starting with a letter, plus sign (+) or slash (/).
	// Only letters and the following special characters are allowed: +-_!@#$%/=
	AppCode string `json:"app_code" required:"true"`
}

type AppCodeOptsBuilder interface {
	ToAppCodeOptsMap() (map[string]interface{}, error)
}

func (opts AppCodeOpts) ToAppCodeOptsMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create is a method by which to create function that create a APIG dedicated instance.
func CreateAppCode(client *golangsdk.ServiceClient, instanceId, appId string, opts AppCodeOptsBuilder) (r CodeResult) {
	reqBody, err := opts.ToAppCodeOptsMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(codeURL(client, instanceId, appId), reqBody, &r.Body, nil)
	return
}

// AutoGenerateAppCode is a method used to automatically create code in a specified application.
func AutoGenerateAppCode(client *golangsdk.ServiceClient, instanceId, appId string) (r CodeResult) {
	_, r.Err = client.Put(codeURL(client, instanceId, appId), nil, &r.Body, nil)
	return
}

// Get is a method to obtain the specified code of the specified application of the specified instance by codeId.
func GetAppCode(client *golangsdk.ServiceClient, instanceId, appId, codeId string) (r CodeResult) {
	_, r.Err = client.Get(codeResourceURL(client, instanceId, appId, codeId), &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return
}

type ListCodeOpts struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type ListCodeOptsBuilder interface {
	ToAppCodeListQuery() (string, error)
}

func (opts ListCodeOpts) ToAppCodeListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

// ListAppCode is a method to obtain the application code list of the specified application of the specified instance.
func ListAppCode(client *golangsdk.ServiceClient, instanceId, appId string, opts ListCodeOptsBuilder) pagination.Pager {
	url := codeURL(client, instanceId, appId)
	if opts != nil {
		query, err := opts.ToAppCodeListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return AppCodePage{pagination.SinglePageBase(r)}
	})
}

//RemoveAppCode is a method to delete an existing code from a specified application.
func RemoveAppCode(client *golangsdk.ServiceClient, instanceId, appId, codeId string) (r DeleteResult) {
	_, r.Err = client.Delete(codeResourceURL(client, instanceId, appId, codeId), nil)
	return
}
