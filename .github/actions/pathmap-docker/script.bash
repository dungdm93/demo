#!/bin/bash
set -x

id

hostname

echo $HOME

echo $PATH

ls -ahl ~

env | sort


echo "/github/home/path/from/docker/action" >> $GITHUB_PATH
echo "/another/path/from/docker/action" >> $GITHUB_PATH

# ls -ahl /__e
