import { defineConfig } from '@rspress/core';

type SidebarItem = {
  text: string;
  link?: string;
  items?: SidebarItem[];
};

type GroupDef = {
  key: string;
  text: Record<string, string>;
  link: string;
  items: { text: string; link: string }[];
};

const moduleGroups: GroupDef[] = [
  {
    key: 'core',
    text: { 'zh-CN': '核心工具', en: 'Core Utilities', 'zh-TW': '核心工具' },
    link: '/modules/core/',
    items: [
      { text: 'must', link: '/modules/core/must' },
      { text: 'orm', link: '/modules/core/orm' },
    ],
  },
  {
    key: 'data',
    text: { 'zh-CN': '数据处理', en: 'Data Processing', 'zh-TW': '數據處理' },
    link: '/modules/data/',
    items: [
      { text: 'candy', link: '/modules/data/candy' },
      { text: 'json', link: '/modules/data/json' },
      { text: 'stringx', link: '/modules/data/stringx' },
      { text: 'anyx', link: '/modules/data/anyx' },
    ],
  },
  {
    key: 'cache',
    text: { 'zh-CN': '缓存策略', en: 'Cache Strategies', 'zh-TW': '緩存策略' },
    link: '/modules/cache/',
    items: [
      { text: '缓存概览', link: '/modules/cache/' },
      { text: 'LRU', link: '/modules/cache/lru' },
      { text: 'LFU', link: '/modules/cache/lfu' },
      { text: 'TinyLFU', link: '/modules/cache/tinylfu' },
      { text: 'SLRU', link: '/modules/cache/slru' },
      { text: 'MRU', link: '/modules/cache/mru' },
      { text: 'ALFU', link: '/modules/cache/alfu' },
      { text: 'ARC', link: '/modules/cache/arc' },
      { text: 'LRU-K', link: '/modules/cache/lruk' },
      { text: 'W-TinyLFU', link: '/modules/cache/wtinylfu' },
      { text: 'FBR', link: '/modules/cache/fbr' },
      { text: 'Optimal', link: '/modules/cache/optimal' },
    ],
  },
  {
    key: 'time',
    text: { 'zh-CN': '时间与调度', en: 'Time & Scheduling', 'zh-TW': '時間與調度' },
    link: '/modules/time/',
    items: [
      { text: 'xtime', link: '/modules/time/xtime' },
      { text: 'xtime996', link: '/modules/time/xtime996' },
      { text: 'xtime955', link: '/modules/time/xtime955' },
      { text: 'xtime007', link: '/modules/time/xtime007' },
    ],
  },
  {
    key: 'system',
    text: { 'zh-CN': '系统与配置', en: 'System & Configuration', 'zh-TW': '系統與配置' },
    link: '/modules/system/',
    items: [
      { text: 'config', link: '/modules/system/config' },
      { text: 'runtime', link: '/modules/system/runtime' },
      { text: 'osx', link: '/modules/system/osx' },
      { text: 'app', link: '/modules/system/app' },
      { text: 'atexit', link: '/modules/system/atexit' },
    ],
  },
  {
    key: 'network',
    text: { 'zh-CN': '网络与安全', en: 'Network & Security', 'zh-TW': '網絡與安全' },
    link: '/modules/network/',
    items: [
      { text: 'network', link: '/modules/network/network' },
      { text: 'cryptox', link: '/modules/network/cryptox' },
      { text: 'pgp', link: '/modules/network/pgp' },
      { text: 'urlx', link: '/modules/network/urlx' },
    ],
  },
  {
    key: 'concurrency',
    text: { 'zh-CN': '并发与控制流', en: 'Concurrency & Control Flow', 'zh-TW': '並發與控制流' },
    link: '/modules/concurrency/',
    items: [
      { text: 'routine', link: '/modules/concurrency/routine' },
      { text: 'wait', link: '/modules/concurrency/wait' },
      { text: 'hystrix', link: '/modules/concurrency/hystrix' },
      { text: 'singledo', link: '/modules/concurrency/singledo' },
      { text: 'event', link: '/modules/concurrency/event' },
    ],
  },
  {
    key: 'dev',
    text: { 'zh-CN': '开发与测试', en: 'Development & Testing', 'zh-TW': '開發與測試' },
    link: '/modules/dev/',
    items: [
      { text: 'fake', link: '/modules/dev/fake' },
      { text: 'randx', link: '/modules/dev/randx' },
      { text: 'defaults', link: '/modules/dev/defaults' },
      { text: 'pyroscope', link: '/modules/dev/pyroscope' },
    ],
  },
];

const validatorLabels: Record<string, { text: string; items: { text: string; link: string }[] }> = {
  'zh-CN': {
    text: 'Validator',
    items: [
      { text: '概览', link: '/validator/' },
      { text: '核心函数', link: '/validator/core' },
      { text: '内置规则', link: '/validator/rules' },
      { text: '自定义验证器', link: '/validator/custom' },
      { text: '验证引擎', link: '/validator/engine' },
      { text: '错误处理', link: '/validator/errors' },
      { text: '多语言', link: '/validator/i18n' },
      { text: '性能与最佳实践', link: '/validator/performance' },
    ],
  },
  en: {
    text: 'Validator',
    items: [
      { text: 'Overview', link: '/en/validator/' },
      { text: 'Core Functions', link: '/en/validator/core' },
      { text: 'Built-in Rules', link: '/en/validator/rules' },
      { text: 'Custom Validators', link: '/en/validator/custom' },
      { text: 'Validation Engine', link: '/en/validator/engine' },
      { text: 'Error Handling', link: '/en/validator/errors' },
      { text: 'i18n', link: '/en/validator/i18n' },
      { text: 'Performance & Best Practices', link: '/en/validator/performance' },
    ],
  },
  'zh-TW': {
    text: 'Validator',
    items: [
      { text: '概覽', link: '/zh-TW/validator/' },
      { text: '核心函數', link: '/zh-TW/validator/core' },
      { text: '內建規則', link: '/zh-TW/validator/rules' },
      { text: '自訂驗證器', link: '/zh-TW/validator/custom' },
      { text: '驗證引擎', link: '/zh-TW/validator/engine' },
      { text: '錯誤處理', link: '/zh-TW/validator/errors' },
      { text: '多語言', link: '/zh-TW/validator/i18n' },
      { text: '效能與最佳實踐', link: '/zh-TW/validator/performance' },
    ],
  },
};

function withLocale(locale: string, path: string) {
  return locale === 'zh-CN' ? path : `/${locale}${path}`;
}

function buildSidebar(locale: 'zh-CN' | 'en' | 'zh-TW', labels: {
  guide: string;
  gettingStarted: string;
  modules: string;
  overview: string;
  api: string;
  apiOverview: string;
}) {
  return [
    {
      text: labels.guide,
      items: [{ text: labels.gettingStarted, link: withLocale(locale, '/guide/getting-started') }],
    },
    {
      text: labels.modules,
      items: [
        { text: labels.overview, link: withLocale(locale, '/modules/overview') },
        ...moduleGroups.map((group) => ({
          text: group.text[locale],
          link: withLocale(locale, group.link),
          items: group.items.map((item) => ({ text: item.text, link: withLocale(locale, item.link) })),
        })),
      ],
    },
    {
      text: labels.api,
      items: [{ text: labels.apiOverview, link: withLocale(locale, '/api/overview') }],
    },
  ];
}

export default defineConfig({
  root: '.',
  lang: 'zh-CN',
  title: 'LazyGophers Utils',
  description: '面向现代 Go 工程的实用工具库文档',
  base: '/utils/',
  route: {
    extensions: ['.md', '.mdx'],
  },

  i18nSource: {
    languagesText: { 'zh-CN': '语言', 'zh-TW': '語言' },
    themeText: { 'zh-CN': '主题', 'zh-TW': '主題' },
    versionsText: { 'zh-CN': '版本', 'zh-TW': '版本' },
    menuTitle: { 'zh-CN': '菜单', 'zh-TW': '選單' },
    outlineTitle: { 'zh-CN': '目录', 'zh-TW': '目錄' },
    scrollToTopText: { 'zh-CN': '回到顶部', 'zh-TW': '回到頂部' },
    lastUpdatedText: { 'zh-CN': '最后更新于', 'zh-TW': '最後更新於' },
    prevPageText: { 'zh-CN': '上一页', 'zh-TW': '上一頁' },
    nextPageText: { 'zh-CN': '下一页', 'zh-TW': '下一頁' },
    sourceCodeText: { 'zh-CN': '源码', 'zh-TW': '原始碼' },
    searchPlaceholderText: { 'zh-CN': '搜索', 'zh-TW': '搜尋' },
    searchPanelCancelText: { 'zh-CN': '取消', 'zh-TW': '取消' },
    searchNoResultsText: { 'zh-CN': '未找到与之匹配的结果', 'zh-TW': '未找到相符結果' },
    searchSuggestedQueryText: { 'zh-CN': '试试搜索不同关键词', 'zh-TW': '試試搜尋不同關鍵詞' },
    'overview.filterNameText': { 'zh-CN': '筛选', 'zh-TW': '篩選' },
    'overview.filterPlaceholderText': { 'zh-CN': '搜索 API', 'zh-TW': '搜尋 API' },
    'overview.filterNoResultText': { 'zh-CN': '未找到匹配的 API', 'zh-TW': '未找到相符 API' },
    openInText: { 'zh-CN': '在 {{name}} 中打开', 'zh-TW': '在 {{name}} 中開啟' },
    copyMarkdownText: { 'zh-CN': '复制 Markdown', 'zh-TW': '複製 Markdown' },
    copyMarkdownLinkText: { 'zh-CN': '复制 Markdown 链接', 'zh-TW': '複製 Markdown 連結' },
    editLinkText: { 'zh-CN': '编辑此页面', 'zh-TW': '編輯此頁面' },
    codeButtonGroupCopyButtonText: { 'zh-CN': '复制代码', 'zh-TW': '複製程式碼' },
    notFoundText: { 'zh-CN': '页面未找到', 'zh-TW': '頁面未找到' },
    takeMeHomeText: { 'zh-CN': '返回首页', 'zh-TW': '返回首頁' },
  },
  locales: [
    {
      lang: 'zh-CN',
      label: '简体中文',
      title: 'LazyGophers Utils',
      description: '面向现代 Go 工程的实用工具库文档',
    },
    {
      lang: 'zh-TW',
      label: '繁體中文',
      title: 'LazyGophers Utils',
      description: '面向現代 Go 工程的實用工具庫文檔',
    },
    {
      lang: 'en',
      label: 'English',
      title: 'LazyGophers Utils',
      description: 'Practical documentation for the LazyGophers Go utility library',
    },
  ],
  themeConfig: {
    nav: [
      { text: '开始', link: '/guide/getting-started' },
      { text: '模块', link: '/modules/overview' },
      { text: 'Validator', link: '/validator/' },
      { text: 'API', link: '/api/overview' },
      { text: 'GitHub', link: 'https://github.com/lazygophers/utils' },
    ],
    sidebar: {
      '/': [
        ...buildSidebar('zh-CN', {
          guide: '指南',
          gettingStarted: '快速开始',
          modules: '模块',
          overview: '模块总览',
          api: 'API',
          apiOverview: 'API 概览',
        }),
        validatorLabels['zh-CN'],
      ],
      '/zh-CN/': [
        ...buildSidebar('zh-CN', {
          guide: '指南',
          gettingStarted: '快速开始',
          modules: '模块',
          overview: '模块总览',
          api: 'API',
          apiOverview: 'API 概览',
        }),
        validatorLabels['zh-CN'],
      ],
      '/en/': [
        ...buildSidebar('en', {
          guide: 'Guide',
          gettingStarted: 'Getting Started',
          modules: 'Modules',
          overview: 'Module Overview',
          api: 'API',
          apiOverview: 'API Overview',
        }),
        validatorLabels['en'],
      ],
      '/zh-TW/': [
        ...buildSidebar('zh-TW', {
          guide: '指南',
          gettingStarted: '快速開始',
          modules: '模組',
          overview: '模組總覽',
          api: 'API',
          apiOverview: 'API 概覽',
        }),
        validatorLabels['zh-TW'],
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
      message: 'AGPL-3.0 Licensed · Copyright © 2024-Present LazyGophers',
    },
  },
});
