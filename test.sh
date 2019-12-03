#! /bin/bash
cat input | docker-semver-tag-generator major 1.1.2
cat input | docker-semver-tag-generator minor 1.1.2

cat input | docker-semver-tag-generator major 1.2.1
cat input | docker-semver-tag-generator major 1.2.1

cat input | docker-semver-tag-generator major 2.1.1
cat input | docker-semver-tag-generator minor 2.1.1

cat input | docker-semver-tag-generator major 4.0.0
cat input | docker-semver-tag-generator minor 4.0.0