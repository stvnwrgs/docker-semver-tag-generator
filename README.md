```input
1.0.0
1.0.1
1.1.0
2.0.0
1.1.1
2.0.1
1.2.0
3.0.0
2.1.0
2.2.0
```

cat input | docker-semver-tag-generator major 1.1.2 == ""
cat input | docker-semver-tag-generator minor 1.1.2 == "1.1"

cat input | docker-semver-tag-generator major 1.2.1 == "1"
cat input | docker-semver-tag-generator major 1.2.1 == "1.2"

cat input | docker-semver-tag-generator major 2.1.1 == ""
cat input | docker-semver-tag-generator minor 2.1.1 == "2.1"

cat input | docker-semver-tag-generator major 4.0.0 == "4"
cat input | docker-semver-tag-generator minor 4.0.0 == "4.0"
