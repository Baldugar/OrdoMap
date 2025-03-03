set -e
go fmt
go build
./ordo-map -settings settings.development.json