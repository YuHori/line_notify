setup:
	sudo yum install -y rpmdevtools yum-utils
	rpmdev-setuptree

build:
	go build line_notify.go

install:
	sudo install -m 755 line_notify /usr/bin/line_notify

package:
	rpmbuild -bb *.spec --define "dist .el7"
