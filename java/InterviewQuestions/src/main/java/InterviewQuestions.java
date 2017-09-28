import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class InterviewQuestions {

    //fn = fn-1 + fn-2
    //f(0) = 0
    //f(1) = 1
    public int fib(int n) {
        if (n == 0 || n == 1) {
            return n;
        } else {
            return fib(n - 1) + fib(n - 2);
        }
    }

    public int fibImp(int n) {
        Map<Integer, Integer> storage = new HashMap<Integer, Integer>();
        storage.put(0, 0);
        storage.put(1, 1);
        return fibImpIntern(n, storage);
    }

    private int fibImpIntern(int n, Map<Integer, Integer> stored) {
        if (stored.containsKey(n)) {
            return stored.get(n);
        } else {
            Integer second = fibImpIntern(n - 2, stored);
            Integer first = fibImpIntern(n - 1, stored);
            stored.put(n - 1, first);
            return first + second;
        }
    }

    /**
     * This wasn't my solution in the interview -- i did something recursive!??
     */
    public List<int[]> findPairsAddingToSum(int[] input, int sum) {
        //Find all pairs of numbers in an array that add up to sum
        //e.g. input = [8, -1, 2, 6, 5], sum = 7 output = [[8, -1], [2, 5]]
        List<int[]> output = new ArrayList<>();
        for(int i=0;i<input.length - 1; i++) {
            int current = input[i];
            for(int j=i+1;j<input.length;j++){
                int test = input[j];
                if(current + test == sum){
                    output.add(new int[] {current, test});
                }
            }
        }
        return output;
    }

    //TODO: (other questions i've gotten and answered but i haven't done them again)
    /**
     * Question
     * Magic Squares -- implement the odd magic square algo
     */

    /**
     * Given a tree print out each row in normal than reverse order
     * eg input tree
     *      1
     *    2   3
     *  4  5 7  6
     *  8  9
     *  output
     *  1
     *  3,2
     *  4,5,7,6
     *  9,8
     */


    /**
     *  Find what's wrong with this snippit, what is a fix
     */
    public int foo(byte x, byte y, byte z) {
        int result = 0;
        for (byte i = x; i < y; i += z) {
            result += i;
        }
        return result;
    }

    public int fooFixed(byte x, byte y, byte z) {
        int result = 0;
        int count = 0;
        for (byte i = x; i < y; i += z) {
            count += 1;

            if(count >= 256){
                return result;
            }
            result += i;
        }
        return result;
    }
}
