'''
Python makes it super-easy to play with the code because there is no need to think about data types and it is super-easy to get code written and tested.add()

For this rewrite I spent just enough time with this problem that there were some obvious optimizations that do seem to give much faster runtimes that have 
nothing to do with the possible optimizations the Interviewer mentioned because he seemed to be oblivious as to how to make the code run faster.

There are some obvious configurations of data that always cause the list or array to be unsplitable.  FOr instance, a length of 1 or less cannot be split if one
assumes both lists must contain at-least 1 item. And others, as shown in the code.

The surprise, for me, was the assumption that an obvious split would be 1/2 of the sum of the list as a baseline and then begin looking for an obvious success.

The next surprise, is that it seems the Big O for the faster function seems to be worse than the original function due to the need to know the sum of the list but
Python sums lists very quickly and this does not count against the apparent runtime as much as it seems it maybe should.
'''
# The Python version:
import os, sys
import time

class SplitError(Exception):
    pass

def findSplitPoint2(arr, n):
    if (n <= 1) or ( (len(set(arr)) == 1) and (n > 1) and (arr[0] != 0) and ((n % 2) == 1) ):
        return -1
    total_sum = sum(arr)
    if (total_sum == 0):
        return 1 if (n > 1) else 0
    target_sum = total_sum / 2
    left_sum = 0
    for i in range(0,n):
        left_sum += arr[i]
        if (left_sum == target_sum):
            isEQ = left_sum == sum(arr[i+1:])
            if (isEQ):
                return i+1
    return -1

def findSplitPoint1(arr, n):
    if (len(arr) <= 1):
        return -1
    leftSum = 0
    for i in range(n):
        leftSum += arr[i]
        rightSum = 0
        for j in range(i+1, n):
            rightSum += arr[j]

        if (leftSum == rightSum): 
            return i+1
    return -1;

analysis = {}

def canSplitEqually(nums):
    startTime1 = time.time()
    splitPoint1 = findSplitPoint1(nums, len(nums))
    endTime1 = time.time()

    startTime2 = time.time()
    splitPoint2 = findSplitPoint2(nums, len(nums))
    endTime2 = time.time()

    result1 = endTime1 - startTime1
    result2 = endTime2 - startTime2
    is_1 = 'faster' if (result1 < result2) else 'slower' if (result1 > result2) else 'equal'
    is_2 = 'faster' if (result2 < result1) else 'slower' if (result2 > result1) else 'equal'
    print("[Inner-1] {} Timing result: {:.4e}".format(is_1 ,result1))
    print("[Inner-2] {} Timing result: {:.4e}".format(is_2, result2))

    key_1 = is_1+'1' if (is_1 != 'equal') else is_1
    key_2 = is_2+'2' if (is_2 != 'equal') else is_2
    if (key_1 not in analysis.keys()):
        analysis[key_1] = 0
    if (key_2 not in analysis.keys()):
        analysis[key_2] = 0
    if (key_1 != 'equal') and (key_2 != 'equal'):
        analysis[key_1] += 1
        analysis[key_2] += 1
    else:
        analysis['equal'] += 1

    if (splitPoint1 != splitPoint2):
        raise SplitError('Failure - splitPoint mismatch in {} {} -> {}.'.format(splitPoint1, splitPoint2, nums))

    cannotSplit = ( (splitPoint1== -1) or (splitPoint1== len(nums)) )

    return cannotSplit == False, result1


if (__name__ == '__main__'):
    test_cases = {}

    test_cases[1] = [[3, 1, 1, 2, 1], True]                  # {4},{4} True
    test_cases[2] = [[4, 1, 1, 2, 1], False]                 # {5},{4} or {4},{5} False
    test_cases[3] = [[8, 8], True]                           # { 8, 8 }; // {8},{8} true
    test_cases[4] = [[1], False]                             # { 1 }; // {1},{} false
    test_cases[5] = [[5, 1, 1, 1, 1, 1], True]               # { 5, 1, 1, 1, 1, 1 }; // {5},{5} True
    test_cases[6] = [[5, 1, 1, 1, 1, 1, 1], False]           # { 5, 1, 1, 1, 1, 1, 1 }; // {6},{5} or {5},{6} False
    test_cases[7] = [[1, 1, 1, 1, 4], True]                  # { 1, 1, 1, 1, 4 }; // {4},{4} True
    test_cases[8] = [[1, 1, 1, 1, 1, 1, 5], False]           # { 1, 1, 1, 1, 1, 1, 5 }; // {6},{5} or {5},{6} False
    test_cases[9] = [[1, 0], False]                          # { 1, 0 }; // {1},{} false
    test_cases[10] = [[0, 1, 0], False]
    test_cases[11] = [[1, 1, 1], False]
    test_cases[12] = [[0], False]
    test_cases[13] = [[0, 0], True]
    test_cases[14] = [[0, 0, 0], True]
    test_cases[15] = [[], False]
    test_cases[16] = [[8, 8, 8], False]
    test_cases[17] = [[0] * 100, True]
    test_cases[18] = [[0] * 1000, True]
    test_cases[19] = [[0] * 10000, True]
    test_cases[20] = [[1] * 101, False]


    count_results = 0;
    result_of_test = None
    total_run_times = 0

    test_num = 1
    count_results = 0
    i_value = 0
    for i, vector in test_cases.items():
        test_data, result = tuple(vector)
        b_value, i_value = canSplitEqually(test_data)
        result_of_test = (b_value == result)
        print('{}. {} -> {}'.format(test_num, "Correct" if (result_of_test) else "Incorrect", test_data))

        if (result_of_test == False):
            print("FAILURE after {} tests !!!".format(count_results))
            sys.exit(0)
        test_num += 1
        print()

    total_run_times += i_value
    print("Timing result: {:.4e}".format(i_value));
    print("\n")

    faster1 = 0
    slower1 = 0
    faster2 = 0
    slower2 = 0
    equals = 0
    for k,v in analysis.items():
        print('{} -> {}'.format(k,v))
        #print(k)
        if (k.find('1') > -1):
            if (k.find('faster') > -1):
                #print('faster1')
                faster1 += v
                continue
            else:
                #print('slower1')
                slower1 += v
                continue
        elif (k.find('2') > -1):
            if (k.find('faster') > -1):
                #print('faster2')
                faster2 += v
                continue
            else:
                #print('slower2')
                slower2 += v
                continue
        else:
            #print('equals')
            equals += v
            continue
        #print()
    print()
    print('faster1 = {}, slower1 = {}'.format(faster1, slower1))
    print('faster2 = {}, slower2 = {}'.format(faster2,slower2))
    print('equals = {}'.format(equals))
