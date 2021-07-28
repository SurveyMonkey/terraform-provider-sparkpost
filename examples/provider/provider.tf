provider "sparkpost" {
  # Recommend using SPARKPOST_API_KEY environment variable instead
  api_key = "APIKEY"

  # Sparkpost API url can be overriden to use the EU region
  base_url = "https://api.sparkpost.com"
}
