
#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=bluebird
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}
