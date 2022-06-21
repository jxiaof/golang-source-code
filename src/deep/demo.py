'''
Descripttion: 
version: 
Author: hujianghong
Date: 2022-03-18 19:34:33
LastEditTime: 2022-03-26 01:06:32
'''
###
# 给出一个字符串数组 words 组成的一本英语词典。返回 words 中最长的一个单词，该单词是由 words 词典中其他单词逐步添加一个字母组成。

# 若其中有多个可行的答案，则返回答案中字典序最小的单词。若无答案，则返回空字符串。

#  

# 示例 1：

# 输入：words = ["w","wo","wor","worl", "world"]
# 输出："world"
# 解释： 单词"world"可由"w", "wo", "wor", 和 "worl"逐步添加一个字母组成。
# 示例 2：

# 输入：words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
# 输出："apple"
# 解释："app

# 来源：力扣（LeetCode）
# 链接：https://leetcode-cn.com/problems/longest-word-in-dictionary
# 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
###


from os import curdir

from typing import List


class Solution:
    def longestWord(self, words: List[str]) -> str:
        words.sort()
        words_set = set(words)
        res = ""
        for word in words:
            if len(word) == 1 or word[:-1] in words_set:
                res = word if len(word) > len(res) else res
        return res

# def foo(li):
#     # 输入：words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
#     minStr = ''
#     curr = 1000
#     m = 0
#     tmp = [].append(li[0])
#     for i in li:
#         if len(i) == m + 1 and i[0: m] in tmp:
#             m = len(i)
#             minStr = i
#         if len(i) == m:
#             tmp.append(i)
#     return minStr


    # for i in range(len(li)):
    #     if (len(li[i]) - m) >= 1:
    #         # (li[i][0:m] in temMap)
    #         m += 1
    #         tmp.append(li[i][0:m])


    #     if 
        # length = len(li[i])
        # temMap[length] = 
 # !bug: todo fix

if __name__ == '__main__':
    li = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
    # foo(li)