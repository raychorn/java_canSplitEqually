import java.io.*;
import java.util.*;

class Solution {

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
   
  public static boolean canSplitEqually(int[] nums) {

    // Write your code here
    int splitPoint = findSplitPoint(nums, nums.length);

    boolean cannotSplit = (splitPoint == -1 || splitPoint == nums.length );
    return  cannotSplit == false;
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

    System.out.println("Test Results:");
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
