#!/bin/bash

# remove tags from github
git tag -d 0.0.1-dev
git push origin :0.0.1-dev

# push new tags
git tag -a v0.1.0 -m "First release"
git push origin v0.1.0
