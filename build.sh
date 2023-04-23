echo "Build for os $1, architecture $2"
env GOOS=$1 GOARCH=$2 go build -ldflags="-X 'main.Config=$(cat ./prod.yaml)'" -o prod
env GOOS=$1 GOARCH=$2 go build -ldflags="-X 'main.Config=$(cat ./stg.yaml)'" -o stg