import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertTrue;

public class NodeTest {

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
     *   76 99
     *    ^ this should be < 75 and > 50
     */
    @Test
    public void bstWithFailureOnUpperBoundCheck(){
        Node node = new Node(100);
        node.left = new Node( 50);
        node.right = new Node(200);
        node.left.left = new Node(25);
        node.left.right = new Node(75);
        node.left.right.right = new Node(99);

        Node test = node.left.right.left = new Node(76); // greater than lowest bound
        assertFalse(node.isBST());
        test.item = 49;
        assertFalse(node.isBST());
        test.item = 51;
        assertTrue(node.isBST());
    }
}
