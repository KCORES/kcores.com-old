KCORES.com 主站
---------------

# Desc

该项目为 kcores.com 主站和主站静态内容生成器. 

# Usage

## 添加新文章

在 ```./content-builder/database/``` 中存放要生成的静态内容的 json 配置文件, 格式为:

```json
{
    "toipc": "家用 NAS 指南",
    "topicIcon": "./assets/content-image/topic-icon/nas-for-home.png",
    "topicDesc": "你的第一台 NAS 何必是 NAS, 老板! 来 100T 硬盘!",
    "longTopicDesc": "你的第一台 NAS 何必是 NAS, 老板! 来 100T 硬盘!",
    "entryList": [
        {
            "title": "家用 NAS 指南 2 - 搭建家庭 NAS 服务器有什么好方案？",
            "desc": "树莓派4B能跑满千兆吗? 树莓派4B作为web服务器最大能达到多少QPS?",
            "link": "https://zhuanlan.zhihu.com/p/84879836",
            "author": "@Karminski-牙医",
            "date": "2019-09-30",
            "cover": "./assets/content-image/entry-cover/v2-3aa84c70ac31e91cdb2866419fe2f284_r.jpg"
        }
    ]
}
```

其中字段分别为: 

- topic 为话题标题, 长度限制 8 中文字符 (超长会折行影响排版， 其他的字段也有长度限制). 
- topicIcon 为话题图标, 大小为 100px * 100px.
- topicDesc 为话题描述, 长度限制为 24 中文字符.
- longTopicDesc 为话题详情页的话题描述, 长度限制为 100 中文字符
- entryList 为 文章列表
    - title 为文章标题, 长度限制 45 中文字符
    - desc 为文章概述
    - link 为文章地址
    - author 为文章作者
    - date 为文章创建日期, 格式为 YYYY-mm-DD
    - cover 为文章封面图, 宽度为 236px, 高度可以自适应, 故不限高度. 

编辑完毕 json 配置文件后, 还需要将相应的图片放入 ```./assets/content-image/toipc-icon``` (存放话题图标), ```./assets/content-image/entry-cover/``` (存放文章封面图) 中.

## 生成内容

编辑完毕内容后, 进入 ```../content-builder/bin/``` 中, 运行 ```content-builder.exe```. 会自动生成内容.  

如果你的系统不是 windows, 或修改了 content-builder, 那么可以进入 ```./content-builder/src/``` 运行 go build 生成你需要的 content-builder 可执行文件.


## 发布

目前还未配置 CI/CD系统, 因此上线请联系 @Karminski-牙医.


