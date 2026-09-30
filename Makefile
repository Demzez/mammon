run:
	#
run-back:
	CONFIG_PATH=./config/config.yaml go run ./cmd/mammon/main.go
run-front:
	cd ./frontend && npm run dev