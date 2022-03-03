

def foo(num):
    """立方根"""
    return num ** (1 / 3)




def baz(num, mi, ma, step):
    """
    10101010001
    """
    for i in range(mi, ma, 1):
        if i * i * i  == num:
            return i
        elif i * i * i  > num:
            print(i)
            if step < 0.0001:
                return i
            else:
                return baz(num*1000, (i-1)*10, i*10, step*0.1)


def bar(a, b):
  r = 1
  while b > 1:
    if b & 1 == 1: #与运算一般可以用于取某位数，这里就是取最后一位。
      r *= a
    a *= a
    b = b >> 1 #这里等价于b//=2
  return r * a


if __name__ == '__main__':
    print(foo(8))
    print(baz(30, 0, 30, 1)*0.00001)
    print(bar(8, 2))





