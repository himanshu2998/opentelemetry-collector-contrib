#!/usr/bin/env bash

release_version=$1
module="healthcheckgrpcextention"

if [ -z "$release_version" ]
then
  echo "empty value provided for release_version"
  exit 1
fi

original_branch=$(git rev-parse --abbrev-ref HEAD)

#branch="release/$module/$release_version"
#git checkout -b "$branch"

rm -rf "$module"

cp -r "./extention/$module" "."

for d in "."; do
    echo $d
    rm -rf "$d"
done

echo "listing files inside $module/"
ls -lrt "$module/"
ls -a "$module/"

cd $module

go mod tidy
go vet ./...

cd ..

#echo "git add files for $module"
#git add .
#
#git commit -m "Commit release changes to branch $branch"
#
#tag_name="$module/$release_version"
#echo "creating and pushing tag $tag_name"
#git tag "$tag_name"
#git push origin "$tag_name"
#
#echo "going back to $original_branch"
#git checkout "$original_branch"