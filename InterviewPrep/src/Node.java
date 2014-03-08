
public class Node {
	public Node left;
	public Node right;
	
	public int getHeight() {
		int leftHeight = -1;
		int rightHeight = -1;
		if (this.left != null) {
			leftHeight = this.left.getHeight(); 
		}
		else leftHeight = 0;
		
		if (this.right != null) {
			rightHeight = this.right.getHeight();
 		}
		else rightHeight = 0;
		
		return Math.max(leftHeight, rightHeight) + 1;
	}
	
	public int getBalancedHeight() {
		int leftHeight;
		int rightHeight;
		
		if (this.left != null) {
			leftHeight = this.left.getBalancedHeight();
			if (leftHeight == -1) return -1;
		}
		else {
			leftHeight = 0;
		}
		if (this.right != null) {
			rightHeight = this.right.getBalancedHeight();
			if (rightHeight == -1) return -1;
 		}
		else {
			rightHeight = 0;
		}
		
		if (Math.abs(leftHeight - rightHeight) > 1) return -1;
		
		return Math.max(leftHeight, rightHeight) + 1;
	}
	
	public boolean isBalanced() {
		return this.getBalancedHeight() != -1;
	}

	/**
	 * @param args
	 */
	public static void main(String[] args) {
		// TODO Auto-generated method stub

	}

}
