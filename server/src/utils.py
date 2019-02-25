#!/usr/bin/env python
# -*- coding: utf-8 -*-

from docx.shared import Mm

"""
毫米数值比较
expect: int (unit:mm)
actual: int (unit:Emu 1mm=36000Emu)
"""
def comparisonMmSize(expect, actual):
    diff = actual - Mm(expect)
    
    if diff < 1000 and diff >= 0:
        # 数值相等
        return True
    else:
        # 数值不等
        return False

def emu2Mm(value):
    _EMUS_PER_MM = 36000
    return int(value / _EMUS_PER_MM)
