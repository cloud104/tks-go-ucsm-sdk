package api

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/cloud104/tks-go-ucsm-sdk/mo"
	"golang.org/x/time/rate"
)

// RateLimit limits the number of requests per second that a Client
// can send to a remote Cisco UCS API endpoint using a rate.Limiter
// token bucket configured with the provided requests per seconds and
// burst. A request will wait for up to the given wait time.
type RateLimit struct {
	// RequestsPerSecond defines the allowed number of requests per second.
	RequestsPerSecond float64

	// Burst is the maximum burst size.
	Burst int

	// Wait defines the maximum time a request will wait for a token to be consumed.
	Wait time.Duration
}

// Config type contains the setting used by the Client.
type Config struct {
	// HttpClient is the HTTP client to use for sending requests.
	// If nil then we use http.DefaultClient for all requests.
	HttpClient *http.Client

	// Endpoint is the base URL to the remote Cisco UCS Manager endpoint.
	Endpoint string

	// Username to use when authenticating to the remote endpoint.
	Username string

	// Password to use when authenticating to the remote endpoint.
	Password string

	// RateLimit is used for limiting the number of requests per second
	// against the remote Cisco UCS API endpoint using a token bucket.
	RateLimit *RateLimit

	// Running context
	Ctx context.Context

	// Debug output: prints raw Request and Rssponse
	Debug bool
}

// Client is used for interfacing with the remote Cisco UCS API endpoint.
type Client struct {
	config  *Config
	apiUrl  *url.URL
	limiter *rate.Limiter
	// Cookie is the authentication cookie currently in use.
	// It's value is set by the AaaLogin and AaaRefresh methods.
	Cookie string
}

// NewClient creates a new API client from the given config.
func NewClient(config Config) (*Client, error) {
	if config.HttpClient == nil {
		config.HttpClient = http.DefaultClient
	}
	if config.Ctx == nil {
		config.Ctx = context.Background()
	}
	baseUrl, err := url.Parse(config.Endpoint)
	if err != nil {
		return nil, err
	}
	apiUrl, err := url.Parse(apiEndpoint)
	if err != nil {
		return nil, err
	}
	var limiter *rate.Limiter
	if config.RateLimit != nil {
		rps := rate.Limit(config.RateLimit.RequestsPerSecond)
		limiter = rate.NewLimiter(rps, config.RateLimit.Burst)
	}
	client := &Client{
		config:  &config,
		apiUrl:  baseUrl.ResolveReference(apiUrl),
		limiter: limiter,
	}
	return client, nil
}

// SetContext sets running Context (default is Background)
func (c *Client) SetContext(ctx context.Context) {
	c.config.Ctx = ctx
}

// Set Debug (default is false)
func (c *Client) SetDebug(debug bool) {
	c.config.Debug = debug
}

// doPost sends a POST request to the remote Cisco UCS API endpoint.
func (c *Client) doPost(in, out interface{}) (err error) {
	var baseResponse BaseResponse
	// Rate limit requests if we are using a limiter
	if c.limiter != nil {
		ctxWithTimeout, cancel := context.WithTimeout(c.config.Ctx, c.config.RateLimit.Wait)
		defer cancel()
		if err = c.limiter.Wait(ctxWithTimeout); err != nil {
			return
		}
	}
	data, err := xmlMarshalWithSelfClosingTags(in)
	if err != nil {
		return
	}
	if c.config.Debug == true {
		fmt.Printf("Request:\n%s\n\n", data)
	}
	r, err := http.NewRequest("POST", c.apiUrl.String(), bytes.NewBuffer(data))
	if err != nil {
		return
	}
	req := r.WithContext(c.config.Ctx)
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.config.HttpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if c.config.Debug == true {
		fmt.Printf("Response:\n%s\n\n", body)
	}
	if xml.Unmarshal(body, &baseResponse) == nil && baseResponse.IsError() {
		err = baseResponse.ToError()
	} else {
		err = xml.Unmarshal(body, &out)
	}
	return
}

// Decode XML response in inner XML document
func (c *Client) getInnerXML(innerXml InnerXml, out mo.Any) (err error) {
	if inner, err := xml.Marshal(innerXml); err == nil {
		// The requested managed objects to return are contained within the inner XML document,
		// which we need to unmarshal first into the given concrete type.
		err = xml.Unmarshal(inner, out)
	}
	return
}

// AaaLogin performs the initial authentication to the remote Cisco UCS API endpoint.
func (c *Client) AaaLogin() (*AaaLoginResponse, error) {
	var resp AaaLoginResponse
	req := AaaLoginRequest{
		InName:     c.config.Username,
		InPassword: c.config.Password,
	}
	if err := c.doPost(req, &resp); err != nil {
		return nil, err
	}
	// Set authentication cookie for future re-use when needed.
	c.Cookie = resp.OutCookie
	return &resp, nil
}

// AaaRefresh refreshes the current session by requesting a new authentication cookie.
func (c *Client) AaaRefresh() (*AaaRefreshResponse, error) {
	var resp AaaRefreshResponse
	req := AaaRefreshRequest{
		InName:     c.config.Username,
		InPassword: c.config.Password,
		InCookie:   c.Cookie,
	}
	if err := c.doPost(req, &resp); err != nil {
		return nil, err
	}
	// Set new authentication cookie
	c.Cookie = resp.OutCookie
	return &resp, nil
}

// AaaKeepAlive sends a request to keep the current session active using the same cookie.
func (c *Client) AaaKeepAlive() (*AaaKeepAliveResponse, error) {
	var resp AaaKeepAliveResponse
	req := AaaKeepAliveRequest{
		Cookie: c.Cookie,
	}
	if err := c.doPost(req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AaaLogout invalidates the current client session.
func (c *Client) AaaLogout() (*AaaLogoutResponse, error) {
	var resp AaaLogoutResponse
	req := AaaLogoutRequest{
		InCookie: c.Cookie,
	}
	if err := c.doPost(req, &resp); err != nil {
		return nil, err
	}
	c.Cookie = ""
	return &resp, nil
}

// ConfigResolveDn retrieves a single managed object for a specified DN.
func (c *Client) ConfigResolveDn(in ConfigResolveDnRequest, out mo.Any) (err error) {
	var resp ConfigResolveDnResponse
	if err = c.doPost(in, &resp); err == nil {
		err = c.getInnerXML(resp.OutConfig, out)
	}
	return
}

// ConfigResolveDns retrieves managed objects for a specified list of DNs.
func (c *Client) ConfigResolveDns(in ConfigResolveDnsRequest, out mo.Any) (err error) {
	var resp ConfigResolveDnsResponse
	if err = c.doPost(in, &resp); err == nil {
		err = c.getInnerXML(resp.OutConfigs, out)
	}
	return
}

// ConfigResolveClass retrieves managed objects of the specified class.
func (c *Client) ConfigResolveClass(in ConfigResolveClassRequest, out mo.Any) (err error) {
	var resp ConfigResolveClassResponse
	if err = c.doPost(in, &resp); err == nil {
		err = c.getInnerXML(resp.OutConfigs, out)
	}
	return
}

// ConfigResolveClasses retrieves managed objects from the specified list of classes.
func (c *Client) ConfigResolveClasses(in ConfigResolveClassesRequest, out mo.Any) (err error) {
	var resp ConfigResolveClassesResponse
	if err = c.doPost(in, &resp); err == nil {
		err = c.getInnerXML(resp.OutConfigs, out)
	}
	return
}

// ConfigResolveChildren retrieves children of managed objects under a specified DN.
func (c *Client) ConfigResolveChildren(in ConfigResolveChildrenRequest, out mo.Any) (err error) {
	var resp ConfigResolveChildrenResponse
	if err = c.doPost(in, &resp); err == nil {
		err = c.getInnerXML(resp.OutConfigs, out)
	}
	return
}

// orgResolveElements retrieves elements of managed objects under a specified Org of given Dn.
func (c *Client) OrgResolveElements(in OrgResolveElementsRequest, out mo.Any) (err error) {
	var resp OrgResolveElementsResponse
	if err = c.doPost(in, &resp); err == nil {
		err = c.getInnerXML(resp.OutConfigs, out)
	}
	return
}

// ConfigConfMo changes managed object.
func (c *Client) ConfigConfMo(in ConfigConfMoRequest, out mo.Any) (err error) {
	var resp ConfigConfMoResponse
	if err = c.doPost(in, &resp); err == nil {
		err = c.getInnerXML(resp.OutConfig, out)
	}
	return
}

// ConfigConfMos changes managed objects.
func (c *Client) ConfigConfMos(in ConfigConfMosRequest, out mo.Any) (err error) {
	var resp ConfigConfMosResponse
	if err = c.doPost(in, &resp); err == nil {
		err = c.getInnerXML(resp.OutConfigs, out)
	}
	return
}

// ConfigEstimateImpact is somewhat of dry run to estimate an impact before submitting changes to LsServer
func (c *Client) ConfigEstimateImpact(in ConfigEstimateImpactRequest, outs ...mo.Any) (err error) {
	var resp ConfigEstimateImpactResponse
	if err = c.doPost(in, &resp); err == nil {
		for _, out := range outs {
			if err = c.getInnerXML(resp.OutAffected, out); err != nil {
				break
			}
		}
	}
	return
}

// LsInstantiateNTemplate instantiates service profile from service profile template
func (c *Client) LsInstantiateNTemplate(in LsInstantiateNTemplateRequest, out mo.Any) (err error) {
	var resp LsInstantiateNTemplateResponse
	if err = c.doPost(in, &resp); err == nil {
		err = c.getInnerXML(resp.OutConfigs, out)
	}
	return
}

// LsInstantiateNNamedTemplate instantiates named service profile from service profile template
func (c *Client) LsInstantiateNNamedTemplate(in LsInstantiateNNamedTemplateRequest, out mo.Any) (err error) {
	var resp LsInstantiateNNamedTemplateResponse
	if err = c.doPost(in, &resp); err == nil {
		err = c.getInnerXML(resp.OutConfigs, out)
	}
	return
}
