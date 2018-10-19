#!/bin/bash

# clone
git clone https://github.com/liguoqinjim/all_readme.git ./files2
cd files2

cp -R ../files/add_readme README.md
git add -f --ignore-errors --all
git -c user.name='liguoqinjim' -c user.email='liguoqinjim23@gmail.com' commit -m "deploy by travis"
git push -f -q https://${token}@github.com/liguoqinjim/github_emoji.git master

