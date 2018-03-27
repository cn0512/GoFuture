package ctp

// #cgo linux LDFLAGS: -fPIC -L${SRCDIR}/v6.3.6 -Wl,-rpath=${SRCDIR}/v6.3.6 -lthostmduserapi -lthosttraderapi -lstdc++
// #cgo linux CPPFLAGS: -fPIC -I${SRCDIR}/v6.3.6
import "C"
