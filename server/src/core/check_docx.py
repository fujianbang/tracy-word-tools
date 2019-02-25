#/usr/bin/env python
# -*- coding: utf-8 -*-

from docx import Document
from docx.shared import Mm


def check_docx(file_name):
    document = Document(file_name)

    # 1. 检查页面边距
    sections = document.sections
    for section in sections:
        if section.top_margin != Mm(37):
            print("上页边距不符合 expect:{},actually:{}".format(section.top_margin, Mm(37)))
        if section.bottom_margin != Mm(27):
            print("下页边距不符合 expect:{},actually:{}".format(section.bottom_margin, Mm(27)))
        if section.left_margin != Mm(28):
            print("左页边距不符合 expect:{},actually:{}".format(section.left_margin, Mm(28)))
        if section.right_margin != Mm(26):
            print("右页边距不符合 expect:{},actually:{}".format(section.right_margin, Mm(26)))


if __name__ == '__main__':
    check_docx("test.docx")

"""
DONE 1. 公文用纸上白边为37mm，下白边为27mm，左白边为28mm,右白边为26mm。
2. 公文每面排22行、每行排28字，行距设置为30磅。
3. 公文标题使用2号方正小标宋_GBK字体，排列使用梯形或菱形。
4. 公文正文使用3号方正仿宋_GBK字体。文中结构层次数依次使用“一”、“（一）”、“1”、“（1）”标注；第一层用方正黑体_GBK字体，第二层用方正楷体_GBK字体，第三、四层用方正仿宋_GBK字体标注。
5. 发文单位署名应与印章对应，发文时间用阿拉伯数字将年、月、日标全，年份应标全称，月、日不编虚位（即1不编为01）。
6. 附件应另面编排，并在版记之前，与正文一起装订。“附件”二字及附件顺序号用3号方正黑体_GBK字体。
7. 文中数字用TimesNewRoman字体
"""
