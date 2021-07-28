package provider

import (
	"context"

	sp "github.com/SparkPost/gosparkpost"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTemplate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTemplateRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Unique alphanumeric ID used to reference the template. Must be unique across the SparkPost account. Maximum length - 64 bytes",
				Type:        schema.TypeString,
				Required:    true,
			},
			"draft": {
				Description: "Whether or not to read the draft or published version of the template.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"name": {
				Description: "Editable display name. At a minimum, id or name is required upon creation. Does not have to be unique. Maximum length - 1024 bytes",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"published": {
				Description: "Whether or not the template is published.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"description": {
				Description: "Description of the template. Maximum length - 1024 bytes",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"content_html": {
				Description: "HTML content for the email's text/html MIME part. At a minimum, html or text is required.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"content_text": {
				Description: "Text content for the email's text/plain MIME part. At a minimum, html or text is required.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"content_subject": {
				Description: "The subject for the template.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"content_from_email": {
				Description: "Email address used to compose the email's From header. The domain must be a verified sending domain.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"content_from_name": {
				Description: "Name used to compose the email's From header.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"content_reply_to": {
				Description: "Email address used to compose the email's Reply-To header.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"content_headers": {
				Description: "Object containing headers other than Subject, From, To, and Reply-To",
				Type:        schema.TypeMap,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"content_email_rfc822": {
				Description: "Pre-built message with the format as described by the message/rfc822 Content-Type.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"options_open_tracking": {
				Description: "Enable or disable open tracking.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"options_click_tracking": {
				Description: "Enable or disable click tracking.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"options_transactional": {
				Description: "Distinguish between transactional and non-transactional messages for unsubscribe and suppression purposes.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
		},
	}
}

func dataSourceTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*sp.Client)

	templateID := d.Get("id").(string)

	draft, ok := d.GetOk("draft")
	if !ok {
		draft = false
	}
	template := &sp.Template{
		ID: templateID,
	}

	_, err := client.TemplateGetContext(ctx, template, draft.(bool))
	if err != nil {
		return diag.FromErr(err)
	}

	err = setTemplateResourceData(d, template)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(templateID)

	return diags
}

func setTemplateResourceData(d *schema.ResourceData, template *sp.Template) error {
	spFrom, err := sp.ParseFrom(template.Content.From)
	if err != nil {
		return err
	}

	d.Set("id", template.ID)
	d.Set("name", template.Name)
	d.Set("published", template.Published)
	d.Set("description", template.Description)
	d.Set("content_html", template.Content.HTML)
	d.Set("content_text", template.Content.Text)
	d.Set("content_subject", template.Content.Subject)
	d.Set("content_from_email", spFrom.Email)
	d.Set("content_from_name", spFrom.Name)
	d.Set("content_reply_to", template.Content.ReplyTo)
	d.Set("content_headers", template.Content.Headers)
	d.Set("content_email_rfc822", template.Content.EmailRFC822)
	d.Set("options_open_tracking", template.Options.OpenTracking)
	d.Set("options_click_tracking", template.Options.ClickTracking)
	d.Set("options_transactional", template.Options.Transactional)

	return nil
}
