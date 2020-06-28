
git add .

dt=$(date '+%A-%b-%d-%Y-%H-%M-%S');

git commit -m "protobuf-$dt"

#git push origin

git push --set-upstream origin master

