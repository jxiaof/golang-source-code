###
#Descripttion: 
#version: 
#Author: soovv
#Date: 2022-05-24 12:47:03
#LastEditors: hujianghong
#LastEditTime: 2022-05-24 12:56:12

# number with Prefix Generator
def gen(prefix, n=100):
    if n < 0:
        n = 10
    for i in range(1000, 1000+n):
        yield f"{prefix}{i}"



if __name__ == '__main__':
    for i in gen("felix", 10):
        print(i)
