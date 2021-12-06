#!/usr/bin/awk -f


BEGIN {
    depth=0;
    horizontal=0;
}
{
    if ($1 == "forward") {horizontal+=$2};
    if ($1 == "down") {depth+=$2};
    if ($1 == "up") {depth-=$2};
}
END {
    print "depth : " depth
    print "horizontal : " horizontal 
    print "horizontal * depth = " horizontal*depth
}