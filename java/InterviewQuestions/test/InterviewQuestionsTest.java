import org.junit.jupiter.api.Test;

import java.time.Duration;
import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

class InterviewQuestionsTest {
    private InterviewQuestions questions = new InterviewQuestions();

    @Test
    void findPairsAddingToSum() {
        //Given array : [2, 4, 3, 5, 6, -2, 4, 7, 8, 9]
        //Given sum : 7 Integer numbers, whose sum is equal to value : 7
        // (2, 5) (4, 3) (3, 4) (-2, 9)
        int []input = {2, 4, 3, 5, 6, -2, 4, 7, 8, 9};
        int sum = 7;
        List<int[]> pairs = questions.findPairsAddingToSum(input, sum);
        verifyPairExists(new int[] {2, 5}, pairs);
        verifyPairExists(new int[] {4, 3}, pairs);
        verifyPairExists(new int[] {3, 4}, pairs);
        verifyPairExists(new int[] {-2, 9}, pairs);
    }

    private void verifyPairExists(int[] expected, List<int[]> pairs) {
        assertTrue(pairs.stream().anyMatch(item ->
                (expected[0] == item[0]) && (expected[1] == item[1]) ||
                        (expected[1] == item[0]) && (expected[0] == item[1])
        ));
    }

    @Test
    public void testFib() {
        assertEquals(0,questions.fib(0));
        assertEquals(1, questions.fib(1));
        assertEquals(1, questions.fib(2));
        assertEquals(2, questions.fib(3));
        assertEquals(3, questions.fib(4));
        assertEquals(5, questions.fib(5));
        assertEquals(8, questions.fib(6));
        assertEquals(13, questions.fib(7));
        assertEquals(21, questions.fib(8));

    }

    @Test
    public void testFib2() {
        assertEquals(832040, questions.fib(30));
    }

    @Test
    public void testFibImp() {
        assertEquals(832040, questions.fibImp(30));
    }

    //@Test // This test fails as expected
    public void testFooIteratesInfinitly() {
        byte zero = 0x0;
        byte oneTwentySeven = 0x7F;
        byte two = 0x2;
        assertTimeoutPreemptively(Duration.ofMillis(100), () -> {
            questions.foo(zero,oneTwentySeven,two);
        });
    }

    @Test
    public void testFooFixedIteratesNonInfinitly() {
        byte zero = 0x0;
        byte oneTwentySeven = 0x7F;
        byte two = 0x2;
        assertTimeoutPreemptively(Duration.ofMillis(100), () -> {
            assertEquals(-254,questions.fooFixed(zero,oneTwentySeven,two));
        });
    }

}