#!/usr/bin/python3

#
# https://zhuanlan.zhihu.com/p/24767059
#

import os
import sys
from glob import glob


def detect(filename, cascade_file='lbpcascade_animeface.xml'):
    if not os.path.isfile(cascade_file):
        raise RuntimeError('%s: not found' % cascade_file)
    
    cascade = cv2.