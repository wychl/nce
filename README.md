# 新概念英语学习资源

本项目提供《新概念英语》（New Concept English）第一册至第四册的配套音频（MP3）和同步歌词文件（LRC），支持美音（US）和英音（UK）双版本。

## 🌐 在线体验

👉 **立即学习**：[新概念英语在线点读工具](https://welearn.shanyaohaicuo.com/nce)

免费使用，支持美音/英音切换、句子级点读、中英对照显示，电脑手机均可访问。

## 📁 目录结构

```
.
├── us/                           # 美音版本
│   ├── NCE1/                     # 第一册
│   │   ├── 001&002.Excuse Me.lrc
│   │   ├── 001&002.Excuse Me.mp3
│   │   ├── ...
│   │   └── 143&144.A Walk Through the Woods.mp3
│   ├── NCE2/                     # 第二册
│   │   ├── 01.A Private Conversation.lrc
│   │   ├── 01.A Private Conversation.mp3
│   │   ├── ...
│   │   └── 96.The Dead Return.mp3
│   ├── NCE3/                     # 第三册
│   │   ├── 01.A Puma at Large.lrc
│   │   ├── 01.A Puma at Large.mp3
│   │   └── ...
│   └── NCE4/                     # 第四册
│       ├── 01.Finding Fossil Man.lrc
│       ├── 01.Finding Fossil Man.mp3
│       └── ...
└── uk/                           # 英音版本（结构与 us 完全相同）
    ├── NCE1/
    ├── NCE2/
    ├── NCE3/
    └── NCE4/
```

## 🎵 文件命名规则

- **第一册**：每两课合并为一个文件，格式 `XXX&XXX.Title.lrc` / `.mp3`（例如 `001&002.Excuse Me.lrc`）。
- **第二、三、四册**：每课独立一个文件，格式 `XX.Title.lrc` / `.mp3`（例如 `01.A Private Conversation.lrc`）。

> 💡 文件名中的英文字母、标点与教材原题完全一致，便于匹配。

## 🎧 使用说明

你可以通过以下方式使用这些资源：

1. **直接播放**：使用任何支持 MP3 的播放器播放音频。
2. **同步歌词**：将 `.lrc` 文件与同名的 `.mp3` 放在同一目录，支持 LRC 的播放器（如手机音乐 App、电脑端 AIMP、VLC 等）会自动显示滚动字幕。
3. **在线学习工具**：本项目已集成到前端 Web 应用中（基于 Next.js + HeroUI），支持句子点读、美音/英音切换、中英对照显示。访问 [新概念英语在线学习](https://welearn.shanyaohaicuo.com/nce) 即可体验。

## 📝 LRC 文件格式

歌词文件内容示例：

```lrc
[ar:山海作品推荐]
[ti:Excuse Me]
[by:山海作品推荐]
[00:00.600]Lesson 1 | 第1课
[00:04.400]Excuse me. | 打扰一下。
[00:07.000]Yes. | 是的。
```

- 每个时间标签后使用竖线 `|` 分隔英文和中文。
- 时间格式为 `[mm:ss.xxx]`（毫秒部分可为 1~3 位）。

## 🔧 数据生成与更新

本资源目录由 `scanBooks` 脚本自动扫描并生成 `data.json` 元数据文件，供前端动态加载。如需更新书籍或单元，只需替换对应的 `.lrc` 和 `.mp3` 文件，重新运行生成脚本即可。

## 📄 许可证

资源文件仅供个人学习使用，请勿用于商业用途。音频版权归原出版方所有。

---

> 🚀 欢迎 Star 或 Fork 本项目，共同完善新概念英语学习工具！
