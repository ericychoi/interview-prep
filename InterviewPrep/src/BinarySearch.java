import java.util.Arrays;
import java.util.Collections;
import java.util.List;


public class BinarySearch {

	/**
	 * @param args
	 */
	public static void main(String[] args) {
		// TODO Auto-generated method stub
		int[] list = createSortedList(32);
		
		System.out.print("array: ");
		for (int i: list) System.out.print(i + ",");
		System.out.println("");
		System.out.println("\n" + binarySearch(list, -2, 0, list.length - 1));
	}
	
	static int binarySearch(int[] ar, int elem, int s, int e) {
		if (ar == null) return -1;
		int mid = (s + e) / 2;
		System.out.println(s+","+e+","+mid);
		if (mid >= ar.length) return -1;
		else if (s == e && ar[mid] != elem) return -1;
		else if (ar[mid] == elem) return mid;
		else if (ar[mid] > elem) return binarySearch(ar, elem, s, mid - 1);
		else return binarySearch(ar, elem, mid + 1, e);
	}
	
	static int[] createRandomList(int size) {
		Integer[] list = new Integer[size];
		for (int i = 0; i < list.length; i++) list[i] = new Integer(i); 
		List<Integer> li = Arrays.asList(list);
		Collections.shuffle(li);
		int[] primList = new int[size];
		int j = 0;
		for (Integer i: li) {
			primList[j++] = i.intValue();
		}
		return primList;
	}
	
	static int[] createSortedList(int size) {
		int[] list = new int[size];
		for (int i = 0; i < list.length; i++) list[i] = i;
		return list;
	}

}
