
public class BitVector {

	/**
	 * @param args
	 */
	public static void main(String[] args) {
		// TODO Auto-generated method stub
		findOpenNumber();
	}
	
	static long numberOfInts = 200;
	static byte[] bitfield = new byte[(int) (numberOfInts/8)];
	
	static void findOpenNumber() {
		int[] numbers = new int[100];
		int j = 1;
		for (int i:numbers) {
			if (j++ != 30) {
			numbers[i] = j;
			}
		}
		
		for (int n:numbers) {
			bitfield[n / 8] |= 1 << (n % 8);
		}
		
		print(bitfield);
		
		
	}
	
	static void print(byte[] bArray) {
		for (byte b: bArray) {
			String s1 = String.format("%8s", Integer.toBinaryString(b & 0xFF)).replace(' ', '0');
			System.out.print(s1 + "|");
		}
		System.out.println("");
		return;
	}

}
