#!/usr/bin/awk -f

BEGIN {
    depth = 0;
    horizontal = 0;
    aim = 0;
}
{
    if ($1 == "down") { aim += $2 }
    if ($1 == "up") { aim -= $2 }
    if ($1 == "forward") {
        horizontal += $2;
        depth += $2 * aim;
    }
}
END {
    print "depth : " depth
    print "horizontal : " horizontal 
    print "horizontal * depth = " horizontal * depth
}
