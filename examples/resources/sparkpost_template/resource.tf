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
