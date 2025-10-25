# build.sh
#!/bin/bash
export PATH=$PATH:/usr/local/go/bin
echo -e "\033[1;34mBuilding App...\033[0m"
go build -o app cmd/server/main.go

if [ $? -eq 0 ]; then
  echo -e "\033[1;32mBuild successful! -> ./app\033[0m"
else
  echo -e "\033[1;31m❌Build failed.\033[0m"
fi
