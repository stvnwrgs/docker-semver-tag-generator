#! /bin/bash
cat versions | ./docker-semver-tag-generator major 3.0.1
cat versions | ./docker-semver-tag-generator minor 3.0.1

cat versions | ./docker-semver-tag-generator major 2.2.1
cat versions | ./docker-semver-tag-generator minor 2.2.1

cat versions | ./docker-semver-tag-generator major 2.3.0
cat versions | ./docker-semver-tag-generator minor 2.3.0

cat versions | ./docker-semver-tag-generator major 1.1.2
cat versions | ./docker-semver-tag-generator minor 1.1.2

cat versions | ./docker-semver-tag-generator major 1.2.1
cat versions | ./docker-semver-tag-generator major 1.2.1

cat versions | ./docker-semver-tag-generator major 2.1.1
cat versions | ./docker-semver-tag-generator minor 2.1.1

cat versions | ./docker-semver-tag-generator major 4.0.0
cat versions | ./docker-semver-tag-generator minor 4.0.0