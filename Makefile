VERSION=1.0.0

build_linux:
	env GOOS=linux GOARCH=amd64 go build -o shellmate
	tar -zcvf shellmate_linux_amd64_v${VERSION}.tar.gz shellmate
	rm shellmate

build_windows:
	env GOOS=windows GOARCH=amd64 go build -o shellmate.exe
	tar -zcvf shellmate_windows_amd64_v${VERSION}.tar.gz shellmate.exe
	rm shellmate.exe

build_android:
	env GOOS=android GOARCH=arm64 go build -o shellmate
	tar -zcvf shellmate_android_arm64_v${VERSION}.tar.gz shellmate
	rm shellmate
