public class Node implements Comparable<Node> {
    public Integer item;
    public Node left = null;
    public Node right = null;

    public Node(int item){
        super();
        this.item = item;
    }

    public boolean isBST() {
        return isBST(null, null);
    }

    private boolean isBST(Integer lowerBound, Integer upperBound) {
        if(left == null && right == null){
            return true;
        }
        else {
            if((left != null && compareTo(left) == -1) ||
                    (right != null && compareTo(right) == 1)){
                return false;
            }
            if((left != null && lowerBound != null && left.item < lowerBound) ||
                    (right != null && upperBound != null && right.item > upperBound)) {
                return false;
            }
            return (left != null && left.isBST(lowerBound, item)) &&
                    (right != null && right.isBST(item, upperBound));
        }
    }

    @Override
    public int compareTo(Node o) {
        return item.compareTo(o.item);
    }
}
