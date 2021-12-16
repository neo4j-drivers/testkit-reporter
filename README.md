# Testkit reporter

## Scope

The report currently includes:

 - the number of skipped tests grouped by features
 - the number of skipped tests grouped by skip reason

## Install

Make sure to install a recent [Go](https://go.dev) version.

```shell
go install github.com/neo4j-drivers/testkit-reporter/cmd/testkit-reporter@latest
```

This installs the binary to `$(go env PATH)/bin`. Make sure this is exported in your standard path.

Note: you can replace `latest` with a specific Git tag.

## Run

The reporter reads from standard input.

### 🐱
```shell
cat build.log | testkit-reporter
Skipped Tests by Feature Flags:

| Feature                             | Tests |
| ----------------------------------- | ----- |
| TMP_FULL_SUMMARY                    | 236   |
| TMP_FAST_FAILING_DISCOVERY          | 42    |
| OPT_AUTHORIZATION_EXPIRED_TREATMENT | 35    |
| TMP_CYPHER_PATH_AND_RELATIONSHIP    | 34    |
[...]

Skipped Tests by Reason:

| Reason                                                                                                        | Tests |
| ------------------------------------------------------------------------------------------------------------- | ----- |
| Needs support for Feature.TMP_FULL_SUMMARY                                                                    | 219   |
| Test does not support cluster                                                                                 | 208   |
| requires investigation                                                                                        | 167   |
| Does not raise the exception                                                                                  | 52    |
[...]
```

### ➰
```shell
curl --header "Authorization: Bearer $TEAMCITY_TOKEN" https://${TEAMCITY_URI}/downloadBuildLog.html\?buildId\=${BUILD_ID} | testkit-reporter
```

