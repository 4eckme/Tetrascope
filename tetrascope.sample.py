// Tetrascope sample
// R=1024px, scale=100%

import numpy as np
from PIL import Image
 
R = 1024
D = 2*R+1
 
def hex_to_rgb(value):
    lv = len(value)
    return tuple(int(value[i:i + lv // 3], 16) for i in range(0, lv, lv // 3))
 
def hex_to_str(c):
    c = c.lstrip('0x')
    if len(c) <= 6:
        return c.ljust(6, '0')
    else:
        return c.substring(0, 6)
 
data = np.zeros((D,D,3), dtype=np.uint8)
sqrR = R*R
for x in range(0, R):
    if not x % 128:
        print(x, " rows rendered")
    for y in range (x, R):
        sqr = (x*x+y*y)
        if sqr <= sqrR:
            rgbtuple = hex_to_rgb(hex_to_str(hex(sqr)))
            data[R+x,R+y] = rgbtuple
            data[R+x,R-y] = rgbtuple
            data[R-x,R+y] = rgbtuple
            data[R-x,R-y] = rgbtuple
            data[R+y,R+x] = rgbtuple
            data[R+y,R-x] = rgbtuple
            data[R-y,R+x] = rgbtuple
            data[R-y,R-x] = rgbtuple
 
img = Image.fromarray(data)
img.save("tetrascope.png")
