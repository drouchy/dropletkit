package dropletkit

type Options struct {
  Token string
  baseUrl string
  version string
}

func DefaultOptions() Options {
  return Options{baseUrl: "https://api.digitalocean.com", version: "v2"}
}
