# docker semver tag generator

docker-semver-tag-generator helps to create short semantiv version tags. By getting a list of available tags and the current tag. It determines if the new tag is the newest major or minor release. If so, it gives out the version. If not it returns "".

Tags which have prerelease or metadata in it won't be seen as a new update.

## Usage

Checkout the [/helper](helper) scripts to see how it can be used.

`cat input | docker-semver-tag-generator major 1.1.2` == "1"
`cat input | docker-semver-tag-generator minor 1.1.2` == "1.1"

### Gitlab with prefix

`./helper/gitlab-get-tags.sh | sed -e 's/v//g' | ./docker-semver-tag-generator minor 4.2.2`
