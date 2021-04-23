public class Solution {
    public int SingleNumber(int[] nums) {
        if (nums.Length == 1) return nums[0];
        Array.Sort(nums);
        for (int i = 0; i < nums.Length-1; i+=2) {
            if (nums[i] != nums[i+1]) {
                return nums[i];
            }
        }

        return nums[nums.Length-1];
    }
}



/*  INITIAL ATTEMPT
    public int SingleNumber(int[] nums) {
        var map = new Dictionary<int, bool>();

        for (int i = 0; i < nums.Length; i++) {
            if (map.ContainsKey(nums[i])) {
                map[nums[i]] = true;
            } else {
                map[nums[i]] = false;
            }
        };

        return map.FirstOrDefault(map => map.Value == false).Key;
    }
*/
