workflow "Deploy Release" {
  on = "push"
  resolves = [" Github Create Release"]
}

action "Filters for Master branch" {
  uses = "actions/bin/filter@master"
  args = "branch master"
}

action " Github Create Release" {
  uses = "frankjuniorr/github-create-release-action@master"
  needs = ["Filters for Master branch"]
  secrets = ["GITHUB_TOKEN"]
}
