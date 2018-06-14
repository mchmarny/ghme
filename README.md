# ghutils [![Build Status](https://travis-ci.org/mchmarny/ghutils.svg?branch=master)](https://travis-ci.org/mchmarny/ghutils) [![Go Report Card](https://goreportcard.com/badge/github.com/mchmarny/ghutils)](https://goreportcard.com/report/github.com/mchmarny/ghutils)

Collection of simple utilities for GitHub teams

## Setup 

Create [Personal access tokens](https://github.com/settings/tokens) with `admin:org` and `user` setting checked. 

## Use

List teams in organization:
`ghutils -o my-org-name`

List users in organization:
`ghutils -o my-org-name`

Print user details:
`ghutils -u someuser`

Add user to team:
`ghutils add -u someuser -t 1234567`



