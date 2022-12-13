import os
import random
import string

import numpy as np
import matplotlib.pyplot as plt

# ================ 说明 ================
# 所有的颜色 color 值可以是直观的 red、blue 等英文单词，
#   也可以是十六进制的颜色值，如 #000000，具体可查：https://www.runoob.com/html/html-colorvalues.html。

# ================ 基础设置 ================
# 是否保存成文件。True 保存，False 不保存
IS_SAVE = False
# 文件保存地址。默认：当前目录/Images/
SAVE_PATH = "./Images/"
# 文件名。不填则随机生成一个文件名
FILE_NAME = ""
# 是否在程序结束运行时弹出预览图。True 是，False 否。弹出结果图后需关闭预览图才能使程序正常退出。
IS_SHOW = True
# 标题
TITLE = "Here's Title"

# ================ 坐标轴设置 ================
# 坐标轴颜色，分别是上、下、左、右。如果不显示，设置为 None 即可
AXIS_TOP_COLOR = None
AXIS_BOTTOM_COLOR = "black"
AXIS_LEFT_COLOR = "black"
AXIS_RIGHT_COLOR = None
# 坐标轴宽度（粗细）
AXIS_LINE_WIDTH = 1
# X、Y 轴标题。不显示设置为 "" 即可
TITLE_X = ""
TITLE_Y = "uV"
# 负坐标轴是否翻转到正坐标轴
IS_AXIS_TURN = True
# X 轴刻度标签，主要是用于覆盖掉下面的 SCALE_X。不需要则为 []
SCALE_LABEL_X = ['Left', 'Right']
# X 轴刻度。如果为 []，则会自动生成
SCALE_X = [0.3, 0.7]
# Y 轴刻度标签，主要是用于覆盖掉下面的 SCALE_Y。不需要则为 []，数量要跟 SCALE_Y 一致。
SCALE_LABEL_Y = [0, -1, -2, -3, -4]
# Y 轴刻度。如果为 []，则会自动生成
SCALE_Y = [0, 1, 2, 3, 4]

# ================ 数据设置 ================
# 横坐标值，根据上面计算出来的
X = np.array(SCALE_X)
# 设置纵坐标值。
# 每一个大括号代表一组数据
DATA = [
    {
        "data": [-1.864, -3.678],
        "color": "black",
        "label": "RI",
        "std_err": [1, 1],  # 标准差值
    },
    {
        "data": [-2.150, -4.051],
        "color": "#DCDCDC",
        "label": "RL",
        "std_err": [1, 1],  # 标准差值
    },
]
# 柱状图宽度
WIDTH = 0.1
# 柱状图间隔宽度
APART = 0.02
# 标准差线条宽度、颜色、横线长度
STD_ERR_ADDR = {
    "elinewidth": 1,
    "ecolor": "green",
    "capsize": 5
}

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
    ax.set_xticklabels(SCALE_LABEL_X)
    # 设置 Y 轴刻度&标签
    ax.set_yticks(SCALE_Y)
    ax.set_yticklabels(SCALE_LABEL_Y)

    # 绘制
    for idx, item in enumerate(DATA):
        # 计算
        left_point = (WIDTH * len(item['data']) + APART * (len(item['data']) - 1)) / 2 - WIDTH / 2
        distance = -left_point + (WIDTH + APART) * idx
        # 负数转正数
        if IS_AXIS_TURN:
            for k, datum in enumerate(item['data']):
                item['data'][k] = abs(item['data'][k])

        # 渲染
        ax.bar(X + distance, item['data'], color=item['color'], width=WIDTH, label=item['label'], yerr=item['std_err'],
               error_kw=STD_ERR_ADDR)
    # 颜色说明
    ax.legend()

    # 文件保存
    if IS_SAVE:
        # 判断文件夹是否存在
        if not os.path.exists(SAVE_PATH):
            os.makedirs(SAVE_PATH)

        # 文件名判断
        if FILE_NAME == "":
            FILE_NAME = "histogram_" + ''.join(random.sample(string.ascii_letters + string.digits, 8))

        # 保存
        plt.savefig(SAVE_PATH + FILE_NAME)

    # 预览图
    if IS_SHOW:
        plt.show()
