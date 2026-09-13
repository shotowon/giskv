# ADDI x5, x0, -42
addi x5, x0, -42

# -42 as uint32 = 4294967254
# 4294967254 < 5 → false
sltiu x7, x5, 5
