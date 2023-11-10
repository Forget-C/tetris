# tetris

![img1.png](docs%2Fimg1.png)
## 设计：
将屏幕划分为网格, 根据网格坐标判断方块位置。

网格大小计算： width/blockSize, height/blockSize

![img.png](docs%2Fimg.png)

屏幕中的图形分为 当前的（current）和归档的（mesh）两部分。

当前的为可控的图形，归档的为已经落地的图形。

## 备注：
当前仅有基本功能， 还有很多bug也不好操作...
工作摸鱼时间瞎搞的， 没有参考别人的设计， 所以踩了不少坑, 不过挺好玩的 哈哈