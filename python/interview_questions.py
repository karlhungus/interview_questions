#!env python3

# given a positive integer as a string, and an integer base, return the
# decimal value of the string.

# i.e. solution('1010', 2) -> 10

def solution1(number, base):
    dec_value = 0
    ord_0 = ord('0')
    for c in number:
        dec_value *= base
        dec_value += ord(c) - ord_0

    return dec_value

# now given the same parameters, plus a second integer base, turn the number
# back into a string representation using the second integer as base. The second
# integer is 2 <= n <= 36.

# i.e. solution('1010', 2, 16) -> 'A'


def solution2(number, base1, base2):
    dec_value = 0
    ord_0 = ord('0')
    for c in number:
        dec_value *= base1
        dec_value += ord(c) - ord_0

    lookup_table = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ'
    string = ''
    while dec_value > 0:
        mod = dec_value % base2
        dec_value = dec_value // base2
        string = '%s%s' % (lookup_table[mod], string)

    return string


assert solution1('1010', 2) == 10
assert solution2('1010', 2, 12) == 'A'
