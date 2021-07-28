---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sparkpost_template Data Source - terraform-provider-sparkpost"
subcategory: ""
description: |-
  
---

# sparkpost_template (Data Source)



## Example Usage

```terraform
data "sparkpost_template" "template_welcome" {
  id = "welcome-email"
  draft = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Unique alphanumeric ID used to reference the template. Must be unique across the SparkPost account. Maximum length - 64 bytes

### Optional

- **draft** (Boolean) Whether or not to read the draft or published version of the template.

### Read-Only

- **content_email_rfc822** (String) Pre-built message with the format as described by the message/rfc822 Content-Type.
- **content_from_email** (String) Email address used to compose the email's From header. The domain must be a verified sending domain.
- **content_from_name** (String) Name used to compose the email's From header.
- **content_headers** (Map of String) Object containing headers other than Subject, From, To, and Reply-To
- **content_html** (String) HTML content for the email's text/html MIME part. At a minimum, html or text is required.
- **content_reply_to** (String) Email address used to compose the email's Reply-To header.
- **content_subject** (String) The subject for the template.
- **content_text** (String) Text content for the email's text/plain MIME part. At a minimum, html or text is required.
- **description** (String) Description of the template. Maximum length - 1024 bytes
- **name** (String) Editable display name. At a minimum, id or name is required upon creation. Does not have to be unique. Maximum length - 1024 bytes
- **options_click_tracking** (Boolean) Enable or disable click tracking.
- **options_open_tracking** (Boolean) Enable or disable open tracking.
- **options_transactional** (Boolean) Distinguish between transactional and non-transactional messages for unsubscribe and suppression purposes.
- **published** (Boolean) Whether or not the template is published.

