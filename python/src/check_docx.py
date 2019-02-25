#/usr/bin/env python
# -*- coding: utf-8 -*-

from docx import Document
from docx.shared import Mm

import utils


"""
1. 检查页面边距
"""
def check_docx(document, check_output):

    sections = document.sections
    for section in sections:
        # 上页边距检查
        if utils.comparisonMmSize(37, section.top_margin):
            check_output['section'].append("上边距：检查通过")
        else:
            err_msg = "期望{}Mm, 实际{}Mm".format(37, utils.emu2Mm(section.top_margin))
            check_output['section'].append("上边距：检查失败\t{}".format(err_msg))

        # 下页边距检查
        if not utils.comparisonMmSize(27, section.bottom_margin):
            err_msg = "期望{}Mm, 实际{}Mm".format(27, utils.emu2Mm(section.bottom_margin))
            check_output['section'].append("下页边距：检查失败\t{}".format(err_msg))
        else:
            check_output['section'].append("下页边距：检查通过")

        # 左页边距检查
        if not utils.comparisonMmSize(28, section.left_margin):
            err_msg = "期望{}Mm, 实际{}Mm".format(28, utils.emu2Mm(section.left_margin))
            check_output['section'].append("左页边距：检查失败\t{}".format(err_msg))
        else:
            check_output['section'].append("左页边距：检查通过")

        # 右页边距检查
        if not utils.comparisonMmSize(26, section.right_margin):
            err_msg = "期望{}Mm, 实际{}Mm".format(26, utils.emu2Mm(section.right_margin))
            check_output['section'].append("右页边距：检查失败\t{}".format(err_msg))
        else:
            check_output['section'].append("右页边距：检查通过")

"""
2. 检查标题
2号 = 22
"""
def check_title(document, check_output):
    paragraphs = document.paragraphs
    for p in paragraphs:
        print(p.text)
        print(p.style.font.name)
        print("==========================")


"""
    main
"""
if __name__ == '__main__':
    # 读取待检测docx
    document = Document("read/demo1.docx")

    # 检查输出项
    check_output = {
        'section': [],
    }

    # 1. 检查页边距
    check_docx(document, check_output)

    # 2. 检查标题
    check_title(document, check_output)

"""
DONE 1. 公文用纸上白边为37mm，下白边为27mm，左白边为28mm,右白边为26mm。
2. 公文每面排22行、每行排28字，行距设置为30磅。
3. 公文标题使用2号方正小标宋_GBK字体，排列使用梯形或菱形。
4. 公文正文使用3号方正仿宋_GBK字体。文中结构层次数依次使用“一”、“（一）”、“1”、“（1）”标注；第一层用方正黑体_GBK字体，第二层用方正楷体_GBK字体，第三、四层用方正仿宋_GBK字体标注。
5. 发文单位署名应与印章对应，发文时间用阿拉伯数字将年、月、日标全，年份应标全称，月、日不编虚位（即1不编为01）。
6. 附件应另面编排，并在版记之前，与正文一起装订。“附件”二字及附件顺序号用3号方正黑体_GBK字体。
7. 文中数字用TimesNewRoman字体
"""
