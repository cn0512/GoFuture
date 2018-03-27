package ctp

// #cgo linux LDFLAGS: -fPIC -L${SRCDIR}/v6.3.5 -Wl,-rpath=${SRCDIR}/v6.3.5 -lthostmduserapi -lthosttraderapi -lstdc++
// #cgo linux CPPFLAGS: -fPIC -I${SRCDIR}/v6.3.5
// #cgo windows LDFLAGS: -fPIC -L${SRCDIR}/v6.3.5 -Wl,-rpath=${SRCDIR}/v6.3.5 ${SRCDIR}/v6.3.5/thostmduserapi.lib ${SRCDIR}/v6.3.5/thosttraderapi.lib -lthostmduserapi -lthosttraderapi
// #cgo windows CPPFLAGS: -fPIC -I${SRCDIR}/v6.3.5 -DISLIB -DWIN32 -DLIB_MD_API_EXPORT -DLIB_TRADER_API_EXPORT
import "C"
