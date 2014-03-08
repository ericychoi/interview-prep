import java.util.HashMap;

public class StringProblems {

	/**
	 * @param args
	 */
	public static void main(String[] args) {
		// TODO Auto-generated method stub
		String testStr = "abbcddddffabfhrf";
		try {
			System.out.println(testStr + ": " + uniqueSubstr(testStr));
		} catch (Exception e) {
			System.out.println("exception"); 
		}
	}
	
	static int uniqueSubstr(String str) throws Exception {
		if(str == null) throw new Exception("str null");
		if(str.length() == 1) return 1;
		HashMap<Character, Integer> charI = new HashMap<Character, Integer>();
		char[] chars = str.toCharArray();
		int start = 0;
		int end = 0;
		int maxSofar = 0;
		while (end < chars.length) {
			char curr = chars[end];
			System.out.println(start + "," + end + "," + curr + "," + maxSofar + "," + charI.toString());
			if (charI.containsKey(curr) && start <= charI.get(curr)) {
				start = charI.get(curr) + 1;
			}
			charI.put(curr, end);
			if (end-start+1 > maxSofar) {
				maxSofar = end-start+1;
			}
			end++;
		}
		
		return maxSofar;
	}
}
