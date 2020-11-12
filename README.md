# ghme

Alfred workflow for navigating to repositories. 

## Setup

Create [Personal access tokens](https://github.com/settings/tokens) with sufficient rights to query API and list repositories.

## Use

Load repositories 

```shell
`ghme load --token $GITHUB_TOKEN --file repos.json`
```

Query repositories 

```shell
`ghme query --file repos.json --value demo`
```

## Disclaimer

This is my personal project and it does not represent my employer. While I do my best to ensure that everything works, I take no responsibility for issues caused by this code.






