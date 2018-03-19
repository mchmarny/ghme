# github-teams-utils [![Build Status](https://travis-ci.org/mchmarny/github-teams-utils.svg?branch=master)](https://travis-ci.org/mchmarny/github-teams-utils)

Collection of simple utilities for GitHub teams

## Setup 

Create [Personal access tokens](https://github.com/settings/tokens) with `admin:org` and `user` setting checked. 

## Use

List teams in organization:
`github-teams-utils -o my-org-name`

Print user details:
`github-teams-utils -u someuser`

Add user to team:
`github-teams-utils add -u someuser -t 1234567`



