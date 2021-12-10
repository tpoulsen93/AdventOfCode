#!/usr/bin/awk -f

function checkLength(value) {
    if (value == 2 || value == 3 || value == 4 || value == 7) {
        return 1
    }
    else {
        return 0
    }
}

BEGIN {
    count = 0;
}
{
    if (checkLength(length($12))) { count++ }
    if (checkLength(length($13))) { count++ }
    if (checkLength(length($14))) { count++ }
    if (checkLength(length($15))) { count++ }
}
END {
    print "count : " count
}
