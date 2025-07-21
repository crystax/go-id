#include "funcdata.h"
#include "go_asm.h"
#include "textflag.h"

/* Copied from https://github.com/golang/go/blob/go1.23.11/src/runtime/go_tls.h#L9-L12 */
#define get_tls(r) MOVQ TLS, r
#define g(r)       0(r)(TLS*1)

/* This function retrieves the "g" pointer.
 * It's opposite to runtime·setg, which can be seen here:
 * https://github.com/golang/go/blob/go1.23.11/src/runtime/asm_amd64.s#L1146-L1152
 */
TEXT ·getg(SB), NOSPLIT, $0-8
    get_tls(CX)
    MOVQ g(CX), AX
    MOVQ AX, ret+0(FP)
    RET
