
public class Sort {

	/**
	 * @param args
	 */
	public static void main(String[] args) {
		// TODO Auto-generated method stub
		int[] arr = {-1, 2, 3, 100, 332, 55, 134, 4 , 79, 10000, -22332};
		System.out.print("Before: ");
		print(arr);
		try {
			selectionSort(arr);
		} catch(Exception e) {
			// do nothing
		}
		System.out.print("After: ");
		print (arr);
	}
	
	public static void print (int[] a) {
		for (int i:a) {
			System.out.print(i);
			System.out.print(",");
		}
		System.out.println("");
	}
	
	static void selectionSort (int[] r) throws Exception {
		if (r == null) {
			throw new Exception("error");
		}
		if (r.length == 1) return;
		for (int i = 0; i < r.length; i++) {
			int minI = i;
			for (int j = i+1; j < r.length; j++) {
				if (r[minI] > r[j]) {
					minI = j;
				}
			}
			int tmp = r[minI];
			r[minI] = r[i];
			r[i] = tmp;
		}
	}
	
	

}
