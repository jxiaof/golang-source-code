from typing import List

class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        if not nums:
            return 0
        if len(nums) == 1:
            return nums[0]

        s, m = nums[0], nums[0]
        for i in range(1, len(nums)):
            if s < 0:
                s = 0
            s += nums[i]

            if s > m:
                m = s
        return m


if __name__ == "__main__":
    s = Solution()
    print(s.maxSubArray([-2,1,-3,4,-1,2,1,-5,4]))
    print(s.maxSubArray([-2,1]))
    print(s.maxSubArray([-2,1,-3]))
        