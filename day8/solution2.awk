#!/usr/bin/awk -f

function getNum(target, nums) {
    split(target, targetArr, "")
    # loop through the nums looking for a matching set
    for (k = 0; k < length(nums); k++) {
        matches = 0
        # check if the length of the strings matches
        if (length(target) == length(nums[k])) {
            # check if all the characters match
            for (l in targetArr) {
                if (index(nums[k], targetArr[l])) {
                    matches++
                }
                if (matches == length(target)) {
                    return k
                }
            }
        }
    }
}

function decode(line, nums) {
    # loop through the four encoded digits and decode them
    n = 3
    for (i = 12; i <= 15; i++) {
        code[n] = getNum(line[i], nums)
        n--
    }
    # we intentionally loaded the array backwards so that we can 
    # convert to decimal easily using the indeces as exponents
    result = 0
    for (j = 0; j < 4; j++) {
        result += code[j]*(10^j)
    }
    return result;
}

function find0(line, six, nine) {
    #iterate through the line checking each string
    for (i=1; i <= 10; i++) {
        # only check 0,6,9
        if (length(line[i]) == 6) {
            # check if the string matches 0 or 6. if not, we found 9
            if (line[i] != nine && line[i] != six) {
                return line[i]
            }
        }
    }
}

function find1(line) {
    for (i = 1; i <= 10; i++) {
        if (length(line[i]) == 2) {
            return line[i]
        }
    }
}

function find2(line, five, three) {
    #iterate through the line checking each string
    for (i=1; i <= 10; i++) {
        # only check 2,3,5
        if (length(line[i]) == 5) {
            # check if the string matches 5 or 3. if not, we found 2
            if (line[i] != five && line[i] != three) {
                return line[i]
            }
        }
    }
}

function find3(line, seven) {
    # turn sev into a char array to be iterated over
    split(seven, sevArr, "")
    #iterate through the line checking each string
    for (i = 1; i <= 10; i++) {
        # only check 2,3,5
        if (length(line[i]) == 5) {
            # count how many matching characters there are
            # if there are 5, we found 5
            matches = 0
            for (j in sevArr) {
                if (index(line[i], sevArr[j])) {
                    matches++
                }
            }
            if (matches == 3) {
                return line[i]
            }
        }
    }
}

function find4(line) {
    for (i = 1; i <= 10; i++) {
        if (length(line[i]) == 4) {
            return line[i]
        }
    }
}

function find5(line, six) {
    # turn six into a char array to be iterated over
    split(six, sixArr, "")
    #iterate through the line checking each string
    for (i = 1; i <= 10; i++) {
        # only check 2,3,5
        if (length(line[i]) == 5) {
            # count how many matching characters there are
            # if there are 5, we found 5
            matches = 0
            for (j in sixArr) {
                if (index(line[i], sixArr[j])) {
                    matches++
                }
            }
            if (matches == 5) {
                return line[i]
            }
        }
    }
}

function find6(line, one) {
    # turn one into a char array to be iterated over
    split(one, oneArr, "")
    #iterate through the line checking each string
    for (i = 1; i <= 10; i++) {
        # only check 0,6,9
        if (length(line[i]) == 6) {
            # check to see if it has both of the values in oneArr
            # if not, we found 6
            for (j in oneArr) {
                if (!index(line[i], oneArr[j])) {
                    return line[i]
                }
            }
        }
    }
}

function find7(line) {
    for (i = 1; i <= 10; i++) {
        if (length(line[i]) == 3) {
            return line[i]
        }
    }
}

function find8(line) {
    for (i = 1; i <= 10; i++) {
        if (length(line[i]) == 7) {
            return line[i]
        }
    }
}

function find9(line, four) {
    # turn four into a char array to be iterated over
    split(four, fourArr, "")
    #iterate through the line checking each string
    for (i = 1; i <= 10; i++) {
        matches = 0
        # only check 0,6,9
        if (length(line[i]) == 6) {
            # check to see if it has all of the values in fourArr
            # if so, we found 0
            for (j in fourArr) {
                if (index(line[i], fourArr[j])) {
                    matches ++
                }
            }
            if (matches == 4) {
                return line[i]
            }
        }
    }
}




BEGIN {
    sum = 0;
    count = 0
}
{
    count++
    # print "@@@@@@@@@@@@@@@@@ " count " @@@@@@@@@@@@@@@@@@@@" 
    
    # build array of each line for passing to functions
    for (i = 1; i <= 15; i++) {
        line[i] = $i
    }

    # put all of our number string codes into an array
    # find the simplest numbers first
    nums[1] = find1(line); #print "1: " nums[1]
    nums[4] = find4(line); #print "4: " nums[4]
    nums[7] = find7(line); #print "7: " nums[7]
    nums[8] = find8(line); #print "8: " nums[8]
    # now find the next set
    nums[6] = find6(line, nums[1]);          #print "6: " nums[6]
    nums[9] = find9(line, nums[4]);          #print "9: " nums[9]
    nums[0] = find0(line, nums[6], nums[9]); #print "0: " nums[0]
    # now find the last set
    nums[5] = find5(line, nums[6]);          #print "5: " nums[5]
    nums[3] = find3(line, nums[7]);          #print "3: " nums[3]
    nums[2] = find2(line, nums[5], nums[3]); #print "2: " nums[2]

    # now we know what each number is, we can start building the 4 digit numbers
    sum += decode(line, nums)
}
END {
    print "sum : " sum
}
