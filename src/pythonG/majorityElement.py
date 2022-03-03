
from typing import List
class Solution:
    def majorityElement1(self, nums: List[int]) -> int:
        m = len(nums) // 2
        t = tuple(nums)
        s1 = set(t)
        for i in s1:
            if nums.count(i) > m:
                return i

    def majorityElement(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return nums[0]
        nums.sort()
        return nums[len(nums) // 2]


if __name__ == "__main__":
    s = Solution()
    print(s.majorityElement([3,2,3,3,4,4,5,6,7,7,8,2,2,2,2,2,2,2,2,2,2,2,2]))