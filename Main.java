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
   public static int findSplitPoint(int arr[], int n) {
      int leftSum = 0 ;
      
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
   
  public static List<Object> canSplitEqually(int[] nums) {
    // Write your code here
    long startTime1 = System.nanoTime();
    int splitPoint = findSplitPoint(nums, nums.length);
    long endTime1 = System.nanoTime();

    long result1 = endTime1 - startTime1;
    System.out.println("[Inner] Timing result: " + result1);

    boolean cannotSplit = (splitPoint == -1 || splitPoint == nums.length );
    
    //return  cannotSplit == false;
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

    boolean results[] = {
      true,
      false,
      true,
      false,
      true,
      false,
      true,
      false
    };

    TestCase tests[] = {
      new TestCase(test1, true),
      new TestCase(test2, false),
      new TestCase(test3, true),
      new TestCase(test4, false),
      new TestCase(test5, true),
      new TestCase(test6, false),
      new TestCase(test7, true),
      new TestCase(test8, false)
    };
    List<Object> the_result;

    for (int i = 0; i < tests.length; i++) {
      the_result = canSplitEqually(tests[i].test);
      System.out.println(String.valueOf(i+1) + ". " + (the_result.get(0) == tests[i].result ? "Correct" : "Incorrect"));

      System.out.println("Timing result: " + the_result.get(1));
      System.out.println("\n");
    }

    System.out.println("1. " + (canSplitEqually(test1) == true ? "Correct" : "Incorrect"));
    System.out.println("2. " + (canSplitEqually(test2) == false ? "Correct" : "Incorrect"));
    System.out.println("3. " + (canSplitEqually(test3) == true ? "Correct" : "Incorrect"));
    System.out.println("4. " + (canSplitEqually(test4) == false ? "Correct" : "Incorrect"));
    System.out.println("5. " + (canSplitEqually(test5) == true ? "Correct" : "Incorrect"));
    System.out.println("6. " + (canSplitEqually(test6) == false ? "Correct" : "Incorrect"));
    System.out.println("7. " + (canSplitEqually(test7) == true ? "Correct" : "Incorrect"));
    System.out.println("8. " + (canSplitEqually(test8) == false ? "Correct" : "Incorrect"));
  }

  public static void main(String[] args) {
    runTests();
  }
}
