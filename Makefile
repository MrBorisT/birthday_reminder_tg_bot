CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
PACKAGE=github.com/MrBorisT/birthday_reminder_tg_bot/cmd/app

build: bindir
	GOOS=linux go build -o ${BINDIR}/app ${PACKAGE}

build-run-win:
	go build -o app.exe cmd/app/main.go && ./app.exe

bindir:
	mkdir -p ${BINDIR}