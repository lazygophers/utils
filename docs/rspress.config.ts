import { defineConfig } from 'rspress/config';

export default defineConfig({
  root: '.',
  lang: 'zh-CN',
  title: 'LazyGophers Utils',
  description: '强大的 Go 工具库，为现代开发工作流设计',
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
  themeConfig: {
    nav: [
      { text: '快速开始', link: '/zh-CN/guide/getting-started' },
      { text: '模块概览', link: '/zh-CN/modules/overview' },
      { text: 'API 文档', link: '/zh-CN/api/overview' },
    ],
    lastUpdated: true,
  },
});
