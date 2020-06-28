
git add .

dt=$(date '+%A-%b-%d-%Y-%H-%M-%S');

git commit -m "grpc-$dt"

#git push origin

git push --set-upstream origin master

