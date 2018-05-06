PREFIX="/usr/local/bin/"
NAME="hex"
all:
	CGO_ENABLED=0 go build -ldflags='-w -s' -o ${NAME}
	strip ${NAME}

install:
	install ${NAME} ${PREFIX}

package:
	gzip -k ${NAME}
