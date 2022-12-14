import os
import random
import string
import graph_data

import numpy as np
import matplotlib.pyplot as plt

# ================ 说明 ================
# 所有的颜色 color 值可以是直观的 red、blue 等英文单词，
#   也可以是十六进制的颜色值，如 #000000，具体可查：https://www.runoob.com/html/html-colorvalues.html。

# ================ 基础设置 ================
# 是否保存成文件。True 保存，False 不保存
IS_SAVE = True
# 文件保存地址。默认：当前目录/Images/
SAVE_PATH = "./Images/"
# 文件名。不填则随机生成一个文件名
FILE_NAME = ""
# 是否在程序结束运行时弹出预览图。True 是，False 否。弹出结果图后需关闭预览图才能使程序正常退出。
IS_SHOW = True
# 标题
TITLE = "F3"

# ================ 坐标轴设置 ================
# 坐标轴颜色，分别是上、下、左、右。如果不显示，设置为 None 即可
AXIS_TOP_COLOR = None
AXIS_BOTTOM_COLOR = "black"
AXIS_LEFT_COLOR = "black"
AXIS_RIGHT_COLOR = None
# 坐标轴宽度（粗细）
AXIS_LINE_WIDTH = 1
# X、Y 轴标题。不显示设置为 "" 即可
TITLE_X = "uV"
TITLE_Y = "ms"
# 负坐标轴是否翻转到正坐标轴
IS_AXIS_TURN = True
# X 轴刻度标签，主要是用于覆盖掉上面的 SCALE_X。不需要则为 []
SCALE_LABEL_X = []
# X 轴刻度。第一个数是起点，第二个数是终点但不包括它，第三个数是间隔的大小
SCALE_X = [-100, 0, 100, 200, 300, 400, 500, 600, 700, 800, 900]
# Y 轴刻度标签，主要是用于覆盖掉上面的 SCALE_Y。不需要则为 []，数量要跟 SCALE_Y 一致。
SCALE_LABEL_Y = [4.0, 3.0, 2.0, 1.0, 0.0, -1.0, -2.0, -3.0, -4.0]
# Y 轴刻度。如果为 []，则会自动生成
SCALE_Y = [-4.0, -3.0, -2.0, -1.0, 0.0, 1.0, 2.0, 3.0, 4.0]
# X 轴起点、终点
X_START = -100
X_END = 900
# Y 轴起点、终点
Y_START = -4.0
Y_END = 4.0

# ================ 网格线设置 ================
# 是否显示网格，True 显示，False 不显示
IS_GRID = True
# 网格线线型、颜色、宽度
GRID_LINE = '-'  # '-', '--', '-.', ':', 'None', ' ', '', 'solid', 'dashed', 'dashdot', 'dotted'
GRID_COLOR = 'red'
GRID_LINE_WIDTH = "0.5"
# 网格线 X/Y 轴显示，可以选：X、Y、both
GRID_AXIS = "both"

# ================ 数据设置 ================
# 横坐标值
# X = np.array([-350, -300, -250, -200, -150, -100, -50, 0, 50, 100, 150, 200, 250, 300])
X = np.linspace(-100, 901, 501)
# 设置纵坐标值。
# 每一个大括号代表一组数据
DATA = [
    {
        "data": graph_data.data1,
        "color": "green",
        "label": "Text-1",
        "style": '-',  # 线条样式，有：'-', '--', '-.', ':', 'None', ' ', '', 'solid', 'dashed', 'dashdot', 'dotted'
        "alpha": 0.3,  # 透明度，取值是：0-1
        "width": 2,  # 线条宽度，值越大越宽
    },
    {
        "data": graph_data.data2,
        "color": "blue",
        "label": "Text-2",
        "style": ':',
        "alpha": 0.6,
        "width": 2,
    },
]
# 是否附上线条颜色说明，True 是，False 否
IS_LEGEND = True
# 线条颜色说明的摆放位置模式，共有 system 和 custom 两种模式
LEGEND_LOC_MODE = "system"
# 使用 system 模式，有这些值：
#     "best"：程序自己寻找最佳位置
#     "upper left"：左上
#     "upper center"：上
#     "upper right"：右上
#     "center left"：左
#     "center"：中
#     "center right"：右
#     "lower left"：左下
#     "lower center"：下
#     "lower right"：右下
LEGEND_LOC = "lower right"
# 使用 custom 模式，可以自定义位置
#     比如 (0.2, 0.4) 就是距离图片左边 0.2，距离图片右边 0.4 的位置
LEGEND_POS = (0.2, 0.2)
# 一行显示多少列
LEGEND_COLUMN = 1

# 程序入口
if __name__ == '__main__':
    # 创建图形对象
    fig = plt.figure()

    # 添加子图区域
    ax = fig.add_axes([0.1, 0.1, 0.8, 0.8])

    # 标题
    ax.set_title(TITLE)

    # 坐标轴 - 标题
    ax.set_xlabel(TITLE_X)
    ax.set_ylabel(TITLE_Y)
    # 坐标轴 - 颜色
    ax.spines['top'].set_color(AXIS_TOP_COLOR)
    ax.spines['bottom'].set_color(AXIS_BOTTOM_COLOR)
    ax.spines['left'].set_color(AXIS_LEFT_COLOR)
    ax.spines['right'].set_color(AXIS_RIGHT_COLOR)
    # 坐标轴 - 宽度
    axis_key = ['top', 'bottom', 'left', 'right']
    for key in axis_key:
        ax.spines[key].set_linewidth(AXIS_LINE_WIDTH)
    # 设置 X 轴刻度&标签
    ax.set_xticks(SCALE_X)
    if SCALE_LABEL_X:
        ax.set_xticklabels(SCALE_LABEL_X)
    # 设置 Y 轴刻度&标签
    ax.set_yticks(SCALE_Y)
    if SCALE_LABEL_Y:
        ax.set_yticklabels(SCALE_LABEL_Y)
    # 设置 X、Y 轴刻度范围
    ax.set_xlim(X_START, X_END)
    ax.set_ylim(Y_START, Y_END)

    # 绘制散点
    for item in DATA:
        # 负数转正数
        if IS_AXIS_TURN:
            for k, datum in enumerate(item['data']):
                item['data'][k] = -(item['data'][k])

        ax.plot(X, item['data'], c=item['color'], label=item['label'], ls=item['style'], alpha=item['alpha'],
                lw=item['width'])

    # 网格
    if IS_GRID:
        ax.grid(linestyle=GRID_LINE, color=GRID_COLOR, linewidth=GRID_LINE_WIDTH, axis=GRID_AXIS)

    # 线条颜色说明
    if IS_LEGEND:
        if LEGEND_LOC_MODE == "custom":
            ax.legend(bbox_to_anchor=LEGEND_POS)
        else:
            ax.legend(loc=LEGEND_LOC)

        # 列数
        ax.legend(ncol=LEGEND_COLUMN)

    # 文件保存
    if IS_SAVE:
        # 判断文件夹是否存在
        if not os.path.exists(SAVE_PATH):
            os.makedirs(SAVE_PATH)

        # 文件名判断
        if FILE_NAME == "":
            FILE_NAME = "graph_" + ''.join(random.sample(string.ascii_letters + string.digits, 8))

        # 保存
        plt.savefig(SAVE_PATH + FILE_NAME)

    # 预览图
    if IS_SHOW:
        plt.show()
