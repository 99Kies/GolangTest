# nums  = [2,7,11,15]

def row(nums, target):
    result = []
    for i in range(len(nums)-1):
        for j in range(i+1,len(nums)):
            print("i %d\tj %d\tnums[i] %d\tnums[j] %d\n"%(i,j,nums[i], nums[j]))
            if nums[i] + nums[j] == target:
                result.append(i)
                result.append(j)
                return result

    return result

class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        
        hashmap = {}
        for index, num in enumerate(nums):
            print(index, num)
            another_num = target - num
            # 存储两数之和，另外一个数的值
            if another_num in hashmap:
                return [hashmap[another_num], index]
            hashmap[num] = index
        return None

nums  = [3,2,4]
target = 6 

print(row(nums, target))

s = Solution()
s.twoSum(nums, target)