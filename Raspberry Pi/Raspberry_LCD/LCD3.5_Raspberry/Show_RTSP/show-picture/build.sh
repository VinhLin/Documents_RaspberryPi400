read -p 'Version: ' version
read -p "Build for Production? " yn

export GIT_COMMIT=$(git log -n 1 --pretty=format:"%h%x09%an%x09%ad%x09%s")
echo $GIT_COMMIT
export DATEBUILD=$(date)
echo $DATEBUILD
mkdir -p build

#packr2
#packr2 clean
# -ldflags="-s -w"
# -s: Omit the symbol table and debug information. They are not needed in a production environment in most of cases.
# -w: Omit the DWARF information.
case $yn in
[Yy]*)
  GOOS=linux GOARCH=arm GOARM=7 go build -ldflags="-s -w -X 'main.DATE=$DATEBUILD' -X 'main.VERSION=$version' -X 'main.GITCOMMIT=$GIT_COMMIT'" -o show_picture_$version . && goupx show_picture_$version && mv show_picture_$version build/
  ;;
*)
  echo "Choose No and build TEST..."
  GOOS=linux GOARCH=arm GOARM=7 go build -ldflags="-s -w -X 'main.DATE=$DATEBUILD' -X 'main.VERSION=$version' -X 'main.GITCOMMIT=$GIT_COMMIT'" -o show_picture_$version . && mv show_picture_$version build/ && scp build/show_picture_$version pi@192.168.10.100:
  ;;
esac


