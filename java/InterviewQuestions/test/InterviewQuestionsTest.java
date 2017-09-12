import org.junit.jupiter.api.Test;

import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

class InterviewQuestionsTest {
    private InterviewQuestions questions = new InterviewQuestions();

    @org.junit.jupiter.api.Test
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

    @Test
    public void bstSimpleTest() {
        Node node = new Node(1);
        assertTrue(node.isBST());
    }

    @Test
    public void bstSimpleFailureTest() {
        Node node = new Node(8);
        node.left = new Node(9);
        assertFalse(node.isBST());
    }

    @Test
    public void bstBoundedFailureTest() {
        Node node = new Node(8);
        node.left = new Node( 5);
        node.right = new Node(10);
        node.left.left = new Node(1);
        node.left.right = new Node(9); // > 8, means invalid
        assertFalse(node.isBST());
    }

    @Test
    public void bstWithSomeDepthTest() {
        Node node = new Node(8);
        node.left = new Node( 5);
        node.right = new Node(10);
        node.left.left = new Node(1);
        node.left.right = new Node(7);
        assertTrue(node.isBST());
    }

    /**
     *   100
     *   50  200
     * 25 75
     *   74 101
     *       ^ this should be less than 100
     */
    @Test
    public void bstWithFailureOnUpperBoundCheck(){
        Node node = new Node(100);
        node.left = new Node( 50);
        node.right = new Node(200);
        node.left.left = new Node(25);
        node.left.right = new Node(75);
        node.left.right.left = new Node(76); // greater than lowest bound
        node.left.right.right = new Node(78);

        assertFalse(node.isBST());
    }
}