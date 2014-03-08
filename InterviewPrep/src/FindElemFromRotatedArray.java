
public class FindElemFromRotatedArray {

	/**
	 * @param args
	 */
	public static void main(String[] args) {
		int[] arr = {15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 14};
		System.out.println(findElem(arr, 5));
	}
	
	// return index when the offset occurs
	// offset is the start of the sorted array
	static int findOffset(int[] arr, int s, int e) {
		if (e - s == 1) {
			return e;
		}
		int mid = (int)((s + e) / 2);
		if (arr[mid] < arr[s]) {
			return findOffset(arr, s, mid);
		}
		else return findOffset(arr, mid, e);
	}
	
	static int binarySearch(int[] arr, int elem, int s, int e) {
		if (arr[s] == elem) return s;
		if (arr[e] == elem) return e;
		int mid = s > e ? (((e + arr.length) - s) / 2 + s) % arr.length : (e - s) / 2 + s;
		//System.out.println(s + "," + e + "," + mid);
		if (arr[mid] > elem) {
			return binarySearch(arr, elem, s, mid);
		}
		else return binarySearch(arr, elem, mid, e);
	}
	
	static int findElem(int[] arr, int elem) {
		// find offset
		int offset = -1;
		offset = findOffset(arr, 0, arr.length - 1);
		
		// binary search with offset
		return binarySearch(arr, elem, offset, (offset + arr.length - 1) % arr.length );
	}
}
