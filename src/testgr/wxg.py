
def bar(n):
    x  = 0
    if n  < 0:
        return  -1 
    if n == 0:
        return 0
    while True:
        if x*x < n:
            x += 1
        else:
            return x
        
def foo(li):
    # lib = set(li)
    tmp = dict()
    li_b = []
    for i in li:
        if i in tmp:
            tmp[i] += 1
            li_b.append(i)
        else:
            tmp[i] = 1
    
    print(li_b)

if __name__ == "__main__":
    li = [1, 2, 3, 3]
    foo(li)
    print(bar(4))
