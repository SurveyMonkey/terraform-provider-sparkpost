package provider

import (
	"context"

	sp "github.com/SparkPost/gosparkpost"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTemplateCreate,
		ReadContext:   resourceTemplateRead,
		UpdateContext: resourceTemplateUpdate,
		DeleteContext: resourceTemplateDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Unique alphanumeric ID used to reference the template. Must be unique across the SparkPost account. Maximum length - 64 bytes",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"template_id": {
				Description: "Unique alphanumeric ID used to set the template ID. Must be unique across the SparkPost account. Maximum length - 64 bytes",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"draft": {
				Description:      "Whether or not to read the draft or published version of the template.",
				Type:             schema.TypeBool,
				Optional:         true,
				Default:          false,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool { return true },
			},
			"name": {
				Description: "Editable display name. At a minimum, id or name is required upon creation. Does not have to be unique. Maximum length - 1024 bytes",
				Type:        schema.TypeString,
				Required:    true,
			},
			"published": {
				Description: "Whether or not the template will be published.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"description": {
				Description: "Description of the template. Maximum length - 1024 bytes",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"content_html": {
				Description: "HTML content for the email's text/html MIME part. At a minimum, html or text is required.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"content_text": {
				Description: "Text content for the email's text/plain MIME part. At a minimum, html or text is required.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"content_subject": {
				Description: "The subject for the template.",
				Type:        schema.TypeString,
				Required:    true},
			"content_from_email": {
				Description: "Email address used to compose the email's From header. The domain must be a verified sending domain.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"content_from_name": {
				Description: "Name used to compose the email's From header.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"content_reply_to": {
				Description: "Email address used to compose the email's Reply-To header.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"content_headers": {
				Description: "Object containing headers other than Subject, From, To, and Reply-To",
				Type:        schema.TypeMap,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"content_email_rfc822": {
				Description: "Pre-built message with the format as described by the message/rfc822 Content-Type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"options_open_tracking": {
				Description: "Enable or disable open tracking.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"options_click_tracking": {
				Description: "Enable or disable click tracking.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"options_transactional": {
				Description: "Distinguish between transactional and non-transactional messages for unsubscribe and suppression purposes.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
		},
		Importer: &schema.ResourceImporter{
			// For simplicity, the provider can only import published templates
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	// get the sparkpost client
	client := m.(*sp.Client)

	templateID := d.Id()

	// SparkPost requires that we request a draft or published version of the template
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

func buildTemplate(templateID string, d *schema.ResourceData) *sp.Template {
	openTracking := d.Get("options_open_tracking").(bool)
	clickTracking := d.Get("options_click_tracking").(bool)
	transactional := d.Get("options_transactional").(bool)

	headers := make(map[string]string)
	for k, v := range d.Get("content_headers").(map[string]interface{}) {
		headers[k] = v.(string)
	}

	template := &sp.Template{
		ID:          templateID,
		Name:        d.Get("name").(string),
		Published:   d.Get("published").(bool),
		Description: d.Get("description").(string),
		Content: sp.Content{
			HTML:    d.Get("content_html").(string),
			Text:    d.Get("content_text").(string),
			Subject: d.Get("content_subject").(string),
			From: sp.From{
				Email: d.Get("content_from_email").(string),
				Name:  d.Get("content_from_name").(string),
			},
			ReplyTo:     d.Get("content_reply_to").(string),
			Headers:     headers,
			EmailRFC822: d.Get("content_email_rfc822").(string),
		},
		Options: &sp.TmplOptions{
			OpenTracking:  &openTracking,
			ClickTracking: &clickTracking,
			Transactional: &transactional,
		},
	}

	return template
}

func publishTemplate(ctx context.Context, d *schema.ResourceData, client *sp.Client, templateID string, publish bool) error {
	if publish {
		// automatically publish the template
		_, err := client.TemplatePublishContext(ctx, templateID)
		if err != nil {
			return err
		}

		// ensure the read looks for published templates
		d.Set("draft", false)
	} else {
		// ensure the read looks for draft templates
		d.Set("draft", true)
	}

	return nil
}

func resourceTemplateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*sp.Client)

	templateID := getTemplateID(d)

	template := buildTemplate(templateID, d)

	id, _, err := client.TemplateCreateContext(ctx, template)
	if err != nil {
		return diag.FromErr(err)
	}

	err = publishTemplate(ctx, d, client, templateID, template.Published)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(id)

	return resourceTemplateRead(ctx, d, m)
}

func resourceTemplateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*sp.Client)

	templateID := d.Id()

	published := d.Get("published").(bool)

	template := buildTemplate(templateID, d)

	var publishUpdate, updatePublished bool

	if d.HasChange("published") && published {
		// was draft now published
		publishUpdate = true

		// WARNING: it's undocumented, but the ?update_published param on the PUT can be overridden by
		// a `published` field in the body. Since we want to update the draft here, we MUST set the
		// published field to false in the PUT body.
		// https://developers.sparkpost.com/api/templates/#templates-put-update-a-published-template
		template.Published = false
	} else if published {
		// was published, no change
		updatePublished = true
	}

	// Update the template. `updatePublished` controls whether to update the published or draft copy.
	_, err := client.TemplateUpdateContext(ctx, template, updatePublished)
	if err != nil {
		return diag.FromErr(err)
	}

	// Publish it if we're going from draft -> published
	err = publishTemplate(ctx, d, client, templateID, publishUpdate)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceTemplateRead(ctx, d, m)
}

func resourceTemplateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*sp.Client)

	templateID := d.Id()

	_, err := client.TemplateDeleteContext(ctx, templateID)
	if err != nil {
		return diag.FromErr(err)
	}

	// mark as deleted
	d.SetId("")

	return diags
}

func getTemplateID(d *schema.ResourceData) string {
	templateID, ok := d.GetOk("template_id")
	if !ok {
		return ""
	}
	return templateID.(string)
}
