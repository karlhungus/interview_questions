public class Node implements Comparable<Node> {
    public Integer item;
    public Node left = null;
    public Node right = null;

    public Node(int item){
        super();
        this.item = item;
    }

    public boolean isBST() {
        return isBST(this, null, null);
    }

    private boolean isBST(Node node, Integer lowerBound, Integer upperBound) {
        if(node == null){
            return true;
        }
        else {
            if((node.left != null && node.left.compareTo(node) == 1) ||
                    (node.right != null && node.right.compareTo(node) == -1)){
                return false;
            }
            if((node.left != null && lowerBound != null && node.left.item < lowerBound) ||
                    (node.right != null && upperBound != null && node.right.item > upperBound)) {
                return false;
            }
            return isBST(node.left, lowerBound, item) &&
                    isBST(node.right, item, upperBound);
        }
    }

    @Override
    public int compareTo(Node o) {
        return item.compareTo(o.item);
    }
}
