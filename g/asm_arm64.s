// SPDX-License-Identifier: MIT
#include "funcdata.h"
#include "go_asm.h"
#include "textflag.h"

/* This function retrieves the "g" pointer.
 * It's opposite to runtime·setg, which can be seen here:
 * https://github.com/golang/go/blob/go1.23.11/src/runtime/asm_arm64.s#L1193-L1198
 */
TEXT ·getg(SB), NOSPLIT, $0-8
    MOVD g, R8
    MOVD R8, ret+0(FP)
    RET
