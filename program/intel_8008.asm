; A clean program to test the subroutine logic.

; Main program starts here
    LAI #$1A      ; Load 26 into A. Change this value to test different cases.
    CAL check_gte_10 ; Jump to our subroutine

    ; Store the results
    LHI #$02
    LLI #$00
    LMA           ; Store original value (still in A) at $0200
    LHI #$02
    LLI #$01
    LMB           ; Store the result from B at $0201
    HLT           ; End of program

; ==================================
; Subroutine: check_gte_10
; Checks if the value in A is >= 10.
; Puts result in X (1 for true, 0 for false).
; ==================================
check_gte_10:
    CPI #$0A      ; Compare A with 10 ($0A)
    JFC is_greater_or_equal   ; Branch if Carry Clear (A >= 10)

is_less:
    LBI #$00      ; A < 10, so set B to 0 (false)
    RET           ; Return from subroutine

is_greater_or_equal:
    LBI #$01      ; A >= 10, so set B to 1 (true)
    RET           ; Return from subroutine