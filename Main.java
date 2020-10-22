import java.io.*;
import java.util.*;

class TestCase 
{ 
    boolean result;  
    int test[] = { }; 
  
    // Constructor 
    TestCase(int a_test[], boolean the_result) 
    { 
        result = the_result; 
        test = a_test; 
    } 
  
    // Driver code 
    static void main(String args[]) 
    { 
        boolean result = true;  
        int test[] = { 3, 1, 1, 2, 1 }; 
        TestCase s = new TestCase(test, result); 
  
        // Below two statements are equivalent 
        System.out.println(s); 
        System.out.println(s.toString()); 
    } 
} 

class Main {

  /*
   * Return true if there is a place to split the array so that the sum on left
   * and right side are equal
   
   goal:  split to make equal ?  
   
   */
  static int array_slice_sum(int arr[], int start) {
    int sum_of_ints = 0;
    for (int j = start + 1; j < arr.length; j++)
      sum_of_ints += arr[j];

    return sum_of_ints;
  }

  public static int findSplitPoint2(int arr[], int n) {
    // How about using Thread Pools ? Let's break-out a function so this can be done
    // for the inner loop.
    // This time we use the function above to compute the inner loop sums,
    // let's create a baseline to see if this adds any overhead to the original.
    int leftSum = 0;
    for (int i = 0; i < n; i++) {
      leftSum += arr[i];
      int rightSum = 0;

      rightSum = array_slice_sum(arr, i);
      /*
       * for (int j = i + 1; j < n; j++) rightSum += arr[j];
       */

      if (leftSum == rightSum)
        return i + 1;
    }

    return -1;
  }

  public static int findSplitPoint1(int arr[], int n) {
    // What could we do to make this function faster?
    // Interviewer thought the last two items were being added multiple times and
    // this is true but...
    // would this be a generalized method for making this faster.
    // Adding the last two items is only faster when there are more than 2 items.
    // Almost all the methods of time required by the inner loop are not generalized
    // and depend on edge cases and will increase the code needlessly.
    // How about using Thread Pools ? Let's break-out a function so this can be done
    // for the inner loop.
    int leftSum = 0;
      for (int i = 0; i < n; i++) {
        leftSum += arr[i];
        int rightSum = 0;
        
        for (int j = i+1 ; j < n ; j++ ) 
            rightSum += arr[j] ;
            
        if (leftSum == rightSum) 
            return i+1 ; 
      }
      
      return -1;
  }
  
  public static List<Object> canSplitEqually1(int[] nums) {
    // Write your code here
    long startTime1 = System.nanoTime();
    int splitPoint = findSplitPoint1(nums, nums.length);
    long endTime1 = System.nanoTime();

    long result1 = endTime1 - startTime1;
    System.out.println("[Inner] Timing result: " + result1);

    boolean cannotSplit = (splitPoint == -1 || splitPoint == nums.length );
    
    //return  cannotSplit == false;
    return Arrays.asList(cannotSplit == false, result1);
  }

  public static List<Object> canSplitEqually2(int[] nums) {
    // Write your code here
    long startTime1 = System.nanoTime();
    int splitPoint = findSplitPoint2(nums, nums.length);
    long endTime1 = System.nanoTime();

    long result1 = endTime1 - startTime1;
    System.out.println("[Inner] Timing result: " + result1);

    boolean cannotSplit = (splitPoint == -1 || splitPoint == nums.length);

    // return cannotSplit == false;
    return Arrays.asList(cannotSplit == false, result1);
  }

  public static void runTests() {

    int test1[] = { 3, 1, 1, 2, 1 }; // {4},{4} True
    int test2[] = { 4, 1, 1, 2, 1 }; // {5},{4} or {4},{5} False
    int test3[] = { 8, 8 }; // {8},{8} true
    int test4[] = { 1 }; // {1},{} false
    int test5[] = { 5, 1, 1, 1, 1, 1 }; // {5},{5} True
    int test6[] = { 5, 1, 1, 1, 1, 1, 1 }; // {6},{5} or {5},{6} False
    int test7[] = { 1, 1, 1, 1, 4 }; // {4},{4} True
    int test8[] = { 1, 1, 1, 1, 1, 1, 5 }; // {6},{5} or {5},{6} False
    int test9[] = { 1, 0 }; // {1},{} false
    int test10[] = { 0, 1, 0 }; // {1},{} false

    boolean results[] = {
      true,
      false,
      true,
      false,
      true,
      false,
      true,
        false, false,
      false
    };

    TestCase tests[] = {
        new TestCase(test1, results[0]), new TestCase(test2, results[1]), new TestCase(test3, results[2]),
        new TestCase(test4, results[3]), new TestCase(test5, results[4]), new TestCase(test6, results[5]),
        new TestCase(test7, results[6]), new TestCase(test8, results[7]), new TestCase(test9, results[8]),
        new TestCase(test10, results[9]), };

    List<Object> the_result1;
    List<Object> the_result2;
    int count_results = 0;
    boolean result_of_test;
    long total_run_times1 = 0;
    long total_run_times2 = 0;
    for (int i = 0; i < tests.length; i++) {
      the_result1 = canSplitEqually1(tests[i].test);
      result_of_test = Boolean.valueOf(the_result1.get(0).toString()) == tests[i].result;
      System.out.println(String.valueOf(i + 1) + ". "
          + (result_of_test ? "Correct" : "Incorrect"));
      if (result_of_test == false) {
        System.out.format("FAILURE after %d tests !!!", count_results);
        System.exit(0);
      }
      total_run_times1 += Long.valueOf(the_result1.get(1).toString());
      System.out.println("Timing result: " + the_result1.get(1));
      System.out.println("\n");

      the_result2 = canSplitEqually2(tests[i].test);
      result_of_test = Boolean.valueOf(the_result2.get(0).toString()) == tests[i].result;
      System.out.println(String.valueOf(i + 1) + ". " + (result_of_test ? "Correct" : "Incorrect"));
      if (result_of_test == false) {
        System.out.format("FAILURE after %d tests !!!", count_results);
        System.exit(0);
      }
      total_run_times2 += Long.valueOf(the_result2.get(1).toString());
      System.out.println("Timing result: " + the_result1.get(1));
      System.out.println("\n");

      count_results++;
    }

    long average_run_time1 = total_run_times1 / tests.length;
    long average_run_time2 = total_run_times2 / tests.length;
    System.out.format("SUCCESS after %d of %d tests for %d !!!\n", count_results, tests.length, average_run_time1);
    System.out.format("SUCCESS after %d of %d tests for %d !!!\n", count_results, tests.length, average_run_time2);

    String faster_or_slower = "faster";
    long faster_or_slower_time = 0;
    if (average_run_time1 < average_run_time2) {
      faster_or_slower_time = average_run_time2 - average_run_time1;
      faster_or_slower = "slower";
    }
    System.out.format("Method 1 was %s than Method 2 by %d !!!\n", faster_or_slower, faster_or_slower_time);
  }

  public static void main(String[] args) {
    runTests();
  }
}
