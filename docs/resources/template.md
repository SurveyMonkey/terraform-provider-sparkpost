---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sparkpost_template Resource - terraform-provider-sparkpost"
subcategory: ""
description: |-
  
---

# sparkpost_template (Resource)



## Example Usage

```terraform
resource "sparkpost_template" "template_welcome" {
  template_id = "welcome-email"
  name = "Welcome Email"
  published = false
  description = "Welcome email template"
  content_html = "<html><body><p>Hi!</p></body></html>"
  content_text = "Hi!"
  content_subject = "Welcome!"
  content_from_email = "noreply@momentive.ai"
  content_headers = {
    "X-Locale" = "en"
  }
  options_transactional = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **content_from_email** (String) Email address used to compose the email's From header. The domain must be a verified sending domain.
- **content_subject** (String) The subject for the template.
- **name** (String) Editable display name. At a minimum, id or name is required upon creation. Does not have to be unique. Maximum length - 1024 bytes

### Optional

- **content_email_rfc822** (String) Pre-built message with the format as described by the message/rfc822 Content-Type.
- **content_from_name** (String) Name used to compose the email's From header.
- **content_headers** (Map of String) Object containing headers other than Subject, From, To, and Reply-To
- **content_html** (String) HTML content for the email's text/html MIME part. At a minimum, html or text is required.
- **content_reply_to** (String) Email address used to compose the email's Reply-To header.
- **content_text** (String) Text content for the email's text/plain MIME part. At a minimum, html or text is required.
- **description** (String) Description of the template. Maximum length - 1024 bytes
- **draft** (Boolean) Whether or not to read the draft or published version of the template.
- **options_click_tracking** (Boolean) Enable or disable click tracking.
- **options_open_tracking** (Boolean) Enable or disable open tracking.
- **options_transactional** (Boolean) Distinguish between transactional and non-transactional messages for unsubscribe and suppression purposes.
- **published** (Boolean) Whether or not the template will be published.
- **template_id** (String) Unique alphanumeric ID used to set the template ID. Must be unique across the SparkPost account. Maximum length - 64 bytes

### Read-Only

- **id** (String) Unique alphanumeric ID used to reference the template. Must be unique across the SparkPost account. Maximum length - 64 bytes

