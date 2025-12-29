import { defineConfig } from 'rspress/config';

export default defineConfig({
  root: '.',
  lang: 'zh-CN',
  title: 'LazyGophers Utils',
  description: '强大的 Go 工具库，为现代开发工作流设计',
  base: '/utils/',
  locales: [
    {
      lang: 'zh-CN',
      label: '简体中文',
      title: 'LazyGophers Utils',
      description: '强大的 Go 工具库，为现代开发工作流设计',
    },
    {
      lang: 'zh-TW',
      label: '繁體中文',
      title: 'LazyGophers Utils',
      description: '強大的 Go 工具庫，為現代開發工作流設計',
    },
    {
      lang: 'en',
      label: 'English',
      title: 'LazyGophers Utils',
      description: 'A powerful Go utility library for modern development workflows',
    },
  ],
  plugins: [],
  themeConfig: {
    nav: [
      { text: '快速开始', link: '/zh-CN/guide/getting-started' },
      { text: '模块概览', link: '/zh-CN/modules/overview' },
      { text: 'API 文档', link: '/zh-CN/api/overview' },
      {
        text: 'GitHub',
        link: 'https://github.com/lazygophers/utils',
      },
    ],
    sidebar: {
      '/zh-CN/': [
        {
          text: '指南',
          items: [
            {
              text: '快速开始',
              link: '/zh-CN/guide/getting-started',
            },
          ],
        },
        {
          text: '模块',
          items: [
            {
              text: '模块概览',
              link: '/zh-CN/modules/overview',
            },
            {
              text: '核心工具',
              link: '/zh-CN/modules/core/',
              items: [
                {
                  text: 'must',
                  link: '/zh-CN/modules/core/must',
                },
                {
                  text: 'orm',
                  link: '/zh-CN/modules/core/orm',
                },
                {
                  text: 'validator',
                  link: '/zh-CN/modules/core/validator',
                },
              ],
            },
            {
              text: '数据处理',
              link: '/zh-CN/modules/data/',
              items: [
                {
                  text: 'candy',
                  link: '/zh-CN/modules/data/candy',
                },
                {
                  text: 'json',
                  link: '/zh-CN/modules/data/json',
                },
                {
                  text: 'stringx',
                  link: '/zh-CN/modules/data/stringx',
                },
                {
                  text: 'anyx',
                  link: '/zh-CN/modules/data/anyx',
                },
              ],
            },
            {
              text: '时间与调度',
              link: '/zh-CN/modules/time/',
              items: [
                {
                  text: 'xtime',
                  link: '/zh-CN/modules/time/xtime',
                },
                {
                  text: 'xtime996',
                  link: '/zh-CN/modules/time/xtime996',
                },
                {
                  text: 'xtime955',
                  link: '/zh-CN/modules/time/xtime955',
                },
                {
                  text: 'xtime007',
                  link: '/zh-CN/modules/time/xtime007',
                },
              ],
            },
            {
              text: '系统与配置',
              link: '/zh-CN/modules/system/',
              items: [
                {
                  text: 'config',
                  link: '/zh-CN/modules/system/config',
                },
                {
                  text: 'runtime',
                  link: '/zh-CN/modules/system/runtime',
                },
                {
                  text: 'osx',
                  link: '/zh-CN/modules/system/osx',
                },
                {
                  text: 'app',
                  link: '/zh-CN/modules/system/app',
                },
                {
                  text: 'atexit',
                  link: '/zh-CN/modules/system/atexit',
                },
              ],
            },
            {
              text: '网络与安全',
              link: '/zh-CN/modules/network/',
              items: [
                {
                  text: 'network',
                  link: '/zh-CN/modules/network/network',
                },
                {
                  text: 'cryptox',
                  link: '/zh-CN/modules/network/cryptox',
                },
                {
                  text: 'pgp',
                  link: '/zh-CN/modules/network/pgp',
                },
                {
                  text: 'urlx',
                  link: '/zh-CN/modules/network/urlx',
                },
              ],
            },
            {
              text: '并发与控制流',
              link: '/zh-CN/modules/concurrency/',
              items: [
                {
                  text: 'routine',
                  link: '/zh-CN/modules/concurrency/routine',
                },
                {
                  text: 'wait',
                  link: '/zh-CN/modules/concurrency/wait',
                },
                {
                  text: 'hystrix',
                  link: '/zh-CN/modules/concurrency/hystrix',
                },
                {
                  text: 'singledo',
                  link: '/zh-CN/modules/concurrency/singledo',
                },
                {
                  text: 'event',
                  link: '/zh-CN/modules/concurrency/event',
                },
              ],
            },
            {
              text: '开发与测试',
              link: '/zh-CN/modules/dev/',
              items: [
                {
                  text: 'fake',
                  link: '/zh-CN/modules/dev/fake',
                },
                {
                  text: 'randx',
                  link: '/zh-CN/modules/dev/randx',
                },
                {
                  text: 'defaults',
                  link: '/zh-CN/modules/dev/defaults',
                },
                {
                  text: 'pyroscope',
                  link: '/zh-CN/modules/dev/pyroscope',
                },
              ],
            },
          ],
        },
        {
          text: 'API',
          items: [
            {
              text: 'API 概览',
              link: '/zh-CN/api/overview',
            },
          ],
        },
      ],
      '/en/': [
        {
          text: 'Guide',
          items: [
            {
              text: 'Getting Started',
              link: '/en/guide/getting-started',
            },
          ],
        },
        {
          text: 'Modules',
          items: [
            {
              text: 'Module Overview',
              link: '/en/modules/overview',
            },
            {
              text: 'Core Utilities',
              link: '/en/modules/core/',
              items: [
                {
                  text: 'must',
                  link: '/en/modules/core/must',
                },
                {
                  text: 'orm',
                  link: '/en/modules/core/orm',
                },
                {
                  text: 'validator',
                  link: '/en/modules/core/validator',
                },
              ],
            },
            {
              text: 'Data Processing',
              link: '/en/modules/data/',
              items: [
                {
                  text: 'candy',
                  link: '/en/modules/data/candy',
                },
                {
                  text: 'json',
                  link: '/en/modules/data/json',
                },
                {
                  text: 'stringx',
                  link: '/en/modules/data/stringx',
                },
                {
                  text: 'anyx',
                  link: '/en/modules/data/anyx',
                },
              ],
            },
            {
              text: 'Time & Scheduling',
              link: '/en/modules/time/',
              items: [
                {
                  text: 'xtime',
                  link: '/en/modules/time/xtime',
                },
                {
                  text: 'xtime996',
                  link: '/en/modules/time/xtime996',
                },
                {
                  text: 'xtime955',
                  link: '/en/modules/time/xtime955',
                },
                {
                  text: 'xtime007',
                  link: '/en/modules/time/xtime007',
                },
              ],
            },
            {
              text: 'System & Configuration',
              link: '/en/modules/system/',
              items: [
                {
                  text: 'config',
                  link: '/en/modules/system/config',
                },
                {
                  text: 'runtime',
                  link: '/en/modules/system/runtime',
                },
                {
                  text: 'osx',
                  link: '/en/modules/system/osx',
                },
                {
                  text: 'app',
                  link: '/en/modules/system/app',
                },
                {
                  text: 'atexit',
                  link: '/en/modules/system/atexit',
                },
              ],
            },
            {
              text: 'Network & Security',
              link: '/en/modules/network/',
              items: [
                {
                  text: 'network',
                  link: '/en/modules/network/network',
                },
                {
                  text: 'cryptox',
                  link: '/en/modules/network/cryptox',
                },
                {
                  text: 'pgp',
                  link: '/en/modules/network/pgp',
                },
                {
                  text: 'urlx',
                  link: '/en/modules/network/urlx',
                },
              ],
            },
            {
              text: 'Concurrency & Control Flow',
              link: '/en/modules/concurrency/',
              items: [
                {
                  text: 'routine',
                  link: '/en/modules/concurrency/routine',
                },
                {
                  text: 'wait',
                  link: '/en/modules/concurrency/wait',
                },
                {
                  text: 'hystrix',
                  link: '/en/modules/concurrency/hystrix',
                },
                {
                  text: 'singledo',
                  link: '/en/modules/concurrency/singledo',
                },
                {
                  text: 'event',
                  link: '/en/modules/concurrency/event',
                },
              ],
            },
            {
              text: 'Development & Testing',
              link: '/en/modules/dev/',
              items: [
                {
                  text: 'fake',
                  link: '/en/modules/dev/fake',
                },
                {
                  text: 'randx',
                  link: '/en/modules/dev/randx',
                },
                {
                  text: 'defaults',
                  link: '/en/modules/dev/defaults',
                },
                {
                  text: 'pyroscope',
                  link: '/en/modules/dev/pyroscope',
                },
              ],
            },
          ],
        },
        {
          text: 'API',
          items: [
            {
              text: 'API Overview',
              link: '/en/api/overview',
            },
          ],
        },
      ],
      '/zh-TW/': [
        {
          text: '指南',
          items: [
            {
              text: '快速開始',
              link: '/zh-TW/guide/getting-started',
            },
          ],
        },
        {
          text: '模組',
          items: [
            {
              text: '模組概覽',
              link: '/zh-TW/modules/overview',
            },
            {
              text: '核心工具',
              link: '/zh-TW/modules/core/',
              items: [
                {
                  text: 'must',
                  link: '/zh-TW/modules/core/must',
                },
                {
                  text: 'orm',
                  link: '/zh-TW/modules/core/orm',
                },
                {
                  text: 'validator',
                  link: '/zh-TW/modules/core/validator',
                },
              ],
            },
            {
              text: '數據處理',
              link: '/zh-TW/modules/data/',
              items: [
                {
                  text: 'candy',
                  link: '/zh-TW/modules/data/candy',
                },
                {
                  text: 'json',
                  link: '/zh-TW/modules/data/json',
                },
                {
                  text: 'stringx',
                  link: '/zh-TW/modules/data/stringx',
                },
                {
                  text: 'anyx',
                  link: '/zh-TW/modules/data/anyx',
                },
              ],
            },
            {
              text: '時間與調度',
              link: '/zh-TW/modules/time/',
              items: [
                {
                  text: 'xtime',
                  link: '/zh-TW/modules/time/xtime',
                },
                {
                  text: 'xtime996',
                  link: '/zh-TW/modules/time/xtime996',
                },
                {
                  text: 'xtime955',
                  link: '/zh-TW/modules/time/xtime955',
                },
                {
                  text: 'xtime007',
                  link: '/zh-TW/modules/time/xtime007',
                },
              ],
            },
            {
              text: '系統與配置',
              link: '/zh-TW/modules/system/',
              items: [
                {
                  text: 'config',
                  link: '/zh-TW/modules/system/config',
                },
                {
                  text: 'runtime',
                  link: '/zh-TW/modules/system/runtime',
                },
                {
                  text: 'osx',
                  link: '/zh-TW/modules/system/osx',
                },
                {
                  text: 'app',
                  link: '/zh-TW/modules/system/app',
                },
                {
                  text: 'atexit',
                  link: '/zh-TW/modules/system/atexit',
                },
              ],
            },
            {
              text: '網絡與安全',
              link: '/zh-TW/modules/network/',
              items: [
                {
                  text: 'network',
                  link: '/zh-TW/modules/network/network',
                },
                {
                  text: 'cryptox',
                  link: '/zh-TW/modules/network/cryptox',
                },
                {
                  text: 'pgp',
                  link: '/zh-TW/modules/network/pgp',
                },
                {
                  text: 'urlx',
                  link: '/zh-TW/modules/network/urlx',
                },
              ],
            },
            {
              text: '並發與控制流',
              link: '/zh-TW/modules/concurrency/',
              items: [
                {
                  text: 'routine',
                  link: '/zh-TW/modules/concurrency/routine',
                },
                {
                  text: 'wait',
                  link: '/zh-TW/modules/concurrency/wait',
                },
                {
                  text: 'hystrix',
                  link: '/zh-TW/modules/concurrency/hystrix',
                },
                {
                  text: 'singledo',
                  link: '/zh-TW/modules/concurrency/singledo',
                },
                {
                  text: 'event',
                  link: '/zh-TW/modules/concurrency/event',
                },
              ],
            },
            {
              text: '開發與測試',
              link: '/zh-TW/modules/dev/',
              items: [
                {
                  text: 'fake',
                  link: '/zh-TW/modules/dev/fake',
                },
                {
                  text: 'randx',
                  link: '/zh-TW/modules/dev/randx',
                },
                {
                  text: 'defaults',
                  link: '/zh-TW/modules/dev/defaults',
                },
                {
                  text: 'pyroscope',
                  link: '/zh-TW/modules/dev/pyroscope',
                },
              ],
            },
          ],
        },
        {
          text: 'API',
          items: [
            {
              text: 'API 概覽',
              link: '/zh-TW/api/overview',
            },
          ],
        },
      ],
    },
    socialLinks: [
      {
        icon: 'github',
        mode: 'link',
        content: 'https://github.com/lazygophers/utils',
      },
    ],
    lastUpdated: true,
    footer: {
      message: 'MIT Licensed | Copyright © 2024-PRESENT LazyGophers',
    },
  },
});
