provider "aws" {
  access_key = "${var.aws_access_key}"
  secret_key = "${var.aws_secret_key}"
  region = "${var.aws_region_main}"
  profile = "${var.profile}"
}

{{range .providers.extra_profiles}}
provider "aws" {
  access_key = "${var.aws_access_key}"
  secret_key = "${var.aws_secret_key}"
  region = "${var.aws_region_main}"
  alias = "{{.alias}}"
  profile = "{{.profile}}"
}
{{end}}
