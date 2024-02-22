
Version := 0.1

GOOS := linux
GOARCH := amd64
export GOOS
export GOARCH

build:
	go build -o linux-timezone ./main.go
	docker build --no-cache . -t linux-timezone:${Version} -f Dockerfile

install:
	-kubectl delete -f dep.yaml
	kubectl apply -f dep.yaml
	-kubectl delete -f dep-env-tz.yaml
	kubectl apply -f dep-env-tz.yaml
	-kubectl delete -f  manag-cur.yaml
	kubectl apply -f  manag-cur.yaml

clean:
	-kubectl delete -f dep.yaml
	-kubectl delete -f dep-env-tz.yaml
	-kubectl delete -f  manag-cur.yaml
