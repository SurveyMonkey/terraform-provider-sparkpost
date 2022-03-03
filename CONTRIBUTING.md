# Contributing

## Publishing a release

1. [Draft a new release](https://github.com/SurveyMonkey/terraform-provider-sparkpost/releases/new).
   - In the "Choose a tag" field, enter a new tag. Name it according to the pattern `v1.2.3`.
   - Follow semver practices. For example, if the current release is `v1.0.0` and you are publishing
     a bugfix, the new tag should be `v1.0.1`.
1. Go to the [Actions tab](https://github.com/SurveyMonkey/terraform-provider-sparkpost/actions).
   Github will begin building your release.
1. When the release workflow finishes, Github will attach signed binaries to the release notes. 
   This will also trigger a webhook that notifies the Terraform registry.
1. After a while, the [Terraform registry page](https://registry.terraform.io/providers/SurveyMonkey/sparkpost/)
   will pick up the new release, and it will appear in the list of published versions there.

If you run into problems, refer to [Terraform's publishing guide].

[Terraform's publishing guide]: https://www.terraform.io/registry/providers/publishing
