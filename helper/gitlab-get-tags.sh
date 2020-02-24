#!/bin/bash
##################### Usage #####################
# Gets tags from a gitlab registry. Uses default
# CI Environment variables.
#
# Defaults to semver tags. Can be overwritten by
# setting TAGS_MATCHING_REGEX.
#
# CI_JOB_TOKEN
# CI_SERVER_HOST
# CI_PROJECT_PATH_SLUG (group%2Fproject)
# CI_TAG_PAGES
#

# default to semver
TAGS_MATCHING_REGEX="${TAGS_MATCHING_REGEX:-^(v)*([0-9]+)\.([0-9]+)\.([0-9]+)(?:-([0-9A-Za-z-]+(?:\.[0-9A-Za-z-]+)*))?(?:\+[0-9A-Za-z-]+)?$}"
PER_PAGE="${PER_PAGE:-100}"
SCRAPE_PAGES=${SCRAPE_PAGES:-100}


if [ -z ${CI_JOB_TOKEN+x} ]; then echo "CI_JOB_TOKEN is unset"; exit 1; fi
if [ -z ${CI_SERVER_HOST+x} ]; then echo "CI_SERVER_HOST is unset"; exit 1; fi
if [ -z ${CI_PROJECT_PATH_SLUG+x} ]; then echo "CI_PROJECT_PATH_SLUG is unset"; exit 1; fi

# Get the first associated registry
REGISTRY_ID=$(curl -s -XGET --header "Authorization: Bearer ${CI_JOB_TOKEN}" "https://${CI_SERVER_HOST}/api/v4/projects/${CI_PROJECT_PATH_SLUG}/registry/repositories/" |jq '.[0].id')

ALL_TAGS=""
for i in {1..100}
do
  # Collect tag names for registry
  TAGS_RAW=$(curl -s -XGET --header "Authorization: Bearer ${CI_JOB_TOKEN}" "https://${CI_SERVER_HOST}/api/v4/projects/${CI_PROJECT_PATH_SLUG}/registry/repositories/${REGISTRY_ID}/tags?per_page=$PER_PAGE&page=$i")
  # exit if error

  rc=$?; if [[ $rc != 0 ]]; then echo "The request was unsucessful."; exit $rc ; fi

  # break if last page is reached
  if [[ $TAGS_RAW == "[]" ]]; then
    break
  fi

  TAGS_CLEAN=$(echo "$TAGS_RAW" | jq -r '.[].name' | grep -E ${TAGS_MATCHING_REGEX}) # 

  ALL_TAGS=$(echo "$TAGS_CLEAN"; echo "$ALL_TAGS";)
done

# print tags remove empty lines
echo "$ALL_TAGS" | sed '/^$/d'
