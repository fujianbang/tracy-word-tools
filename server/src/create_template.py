#/usr/bin/env python
# -*- coding: utf-8 -*-

from docx import Document
from docx.shared import Mm
from docx.enum.section import WD_SECTION

import logging


def create_template():
    document = Document()

    # 1. 设置页面边距
    sections = document.sections
    print("sections length: {}".format(len(sections)))

    for s in sections:
        print(s.start_type)

    document.add_section(WD_SECTION.ODD_PAGE)

    for s in sections:
        print(s.start_type)

    # save the docx file
    document.save("test.docx")

if __name__ == '__main__':
    create_template()

"""
DONE 1. 公文用纸上白边为37mm，下白边为27mm，左白边为28mm,右白边为26mm。
2. 公文每面排22行、每行排28字，行距设置为30磅。
3. 公文标题使用2号方正小标宋_GBK字体，排列使用梯形或菱形。
4. 公文正文使用3号方正仿宋_GBK字体。文中结构层次数依次使用“一”、“（一）”、“1”、“（1）”标注；第一层用方正黑体_GBK字体，第二层用方正楷体_GBK字体，第三、四层用方正仿宋_GBK字体标注。
5. 发文单位署名应与印章对应，发文时间用阿拉伯数字将年、月、日标全，年份应标全称，月、日不编虚位（即1不编为01）。
6. 附件应另面编排，并在版记之前，与正文一起装订。“附件”二字及附件顺序号用3号方正黑体_GBK字体。
7. 文中数字用TimesNewRoman字体
"""
