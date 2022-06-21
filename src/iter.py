'''
Descripttion: 
version: 
Author: hujianghong
Date: 2022-04-19 20:16:08
LastEditTime: 2022-04-20 16:41:54
'''
# from 
# view / 

from hashlib import new


def baz(*args, **kwargs):
    def foo(func):
        @functools.wraps(func)
        def bar(*args, **kwargs):
            print('before', args, kwargs)
            func(*args, **kwargs)
            print('after')
        return bar
    return foo

@baz()
def run():
    print('run')



def test(li):
    """_summary_
    输入: nums = [0,1,0 ,3,12]
    输出: [1,3,12,0,0]
    Args:
        li (_type_): _description_
    """
    length = len(li) # // 2
    i,j = 0,length-1
    for c in range(length):
        if li[c] == 0:
            li[c],li[j] = li[j],li[c]
            j -= 1
        else:
            pass
            
    



if __name__ == '__main__':
    run()