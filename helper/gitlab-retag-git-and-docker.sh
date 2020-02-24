#! /bin/bash

CURRENT_VERSION=$1
REGISTRY_IMAGE=${REGISTRY_IMAGE}
REPO_URL=${REPO_URL}
export GIT_DIR=${GIT_DIR}

if [ -z ${REGISTRY_IMAGE+x} ]; then echo "REGISTRY_IMAGE is unset"; exit 1; fi
if [ -z ${REPO_URL+x} ]; then echo "REPO_URL is unset"; exit 1; fi

function retag_git {
  cd $GIT_DIR
  git pull --tags origin $REPO_URL HEAD:master
  if git rev-parse v$2 >/dev/null 2>&1
  then
      echo "Found old v$2 tag. Will delete it."
      git tag --delete v$2
      git push --delete origin v$2
  fi

  echo "Creating new tag for v$2 which points to v$1"
  git checkout v$1
  git tag v$2
  git checkout master
  git push --tags origin $REPO_URL HEAD:master
  cd -
}

function retag_docker {
  cd $GIT_DIR
  echo "Retagging $2 to point to $1"
  docker tag $REGISTRY_IMAGE:v$1 $REGISTRY_IMAGE:v$2
  docker push $REGISTRY_IMAGE:v$2
  cd -
}

MAJOR=$(./helper/gitlab-get-tags.sh | sed -e 's/v//g' | ./docker-semver-tag-generator major  $CURRENT_VERSION)

MINOR=$(./helper/gitlab-get-tags.sh | sed -e 's/v//g' | ./docker-semver-tag-generator minor $CURRENT_VERSION)

if [ "$MAJOR" != "" ]; then
    retag_docker $CURRENT_VERSION $MAJOR
fi

if [ "$MINOR" != "" ]; then
    retag_docker $CURRENT_VERSION $MINOR
fi