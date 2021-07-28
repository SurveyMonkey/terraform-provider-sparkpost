package provider

import (
	"context"
	"net/http"
	"time"

	sp "github.com/SparkPost/gosparkpost"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"api_key": {
					Description: "The API key for the SparkPost account.",
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("SPARKPOST_API_KEY", nil),
				},
				"base_url": {
					Description: "The base URL for the SparkPost account.",
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("SPARKPOST_BASE_URL", "https://api.sparkpost.com"),
				},
			},
			ResourcesMap: map[string]*schema.Resource{
				"sparkpost_template": resourceTemplate(),
			},
			DataSourcesMap: map[string]*schema.Resource{
				"sparkpost_template": dataSourceTemplate(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configure(version string, p *schema.Provider) schema.ConfigureContextFunc {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		apiKey := d.Get("api_key").(string)
		baseUrl := d.Get("base_url").(string)

		var diags diag.Diagnostics

		cfg := &sp.Config{
			BaseUrl:    baseUrl,
			ApiKey:     apiKey,
			ApiVersion: 1,
		}
		var client sp.Client
		err := client.Init(cfg)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to initialize SparkPost client",
				Detail:   err.Error(),
			})
			return nil, diags
		}

		userAgent := p.UserAgent("terraform-provider-sparkpost", version)

		client.Client = &http.Client{
			Timeout:   5 * time.Second,
			Transport: UserAgentTransport{userAgent, http.DefaultTransport},
		}

		return &client, diags
	}
}

type UserAgentTransport struct {
	userAgent string
	T         http.RoundTripper
}

func (t UserAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", t.userAgent)
	return t.T.RoundTrip(req)
}
