# ghutils [![Build Status](https://travis-ci.org/mchmarny/ghutils.svg?branch=master)](https://travis-ci.org/mchmarny/ghutils) [![Go Report Card](https://goreportcard.com/badge/github.com/mchmarny/ghutils)](https://goreportcard.com/report/github.com/mchmarny/ghutils)

Collection of simple utilities for GitHub teams

## Setup

Create [Personal access tokens](https://github.com/settings/tokens) with `admin:org` and `user` setting checked.

## Use

List teams in organization:
```shell
`ghutils -o my-org-name`
```

List users in organization:
```shell
`ghutils -o my-org-name`
```

Print user details:
```shell
`ghutils -u someuser`
```

Add user to team:
```shell
`ghutils add -u someuser -t 1234567`
```



## Better Use

If you are adding users to specific teams frequently, get the ID of that
team using the `ghutils -o my-org-name` command and create an alias in
your `~/.bash_profile` where `1234567` is the ID of that team

```shell
add2Team() {
    ghutils add -u $1 -t 1234567
}
```

Then whenever you need to add user to that team simply run `add2Team username`







